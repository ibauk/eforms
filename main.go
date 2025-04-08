package main

import (
	// "log"

	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/mail"
	"net/smtp"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbpath = flag.String("db", "./otps.db", "Database to use")
var port = flag.String("port", "1014", "Port to service requests")

const defaultTokenSize = 4

var MyDB *gorm.DB

type EMAILCFG struct {
	SenderName     string
	SenderEmail    string
	RecipientName  string
	RecipientEmail string
	SMTPServer     string
	AuthUser       string
	AuthPassword   string
}

type EVENTCFG struct {
	RallyKey      string   `json:"RallyKey"`
	RallyDesc     string   `json:"RallyDesc"`
	MaxTeeshirts  int      `json:"MaxTeeshirts"`
	TeeshirtSizes []string `json:"TeeshirtSizes"`
	MaxPatches    int      `json:"MaxPatches"`
}

const BBR = `{
	"RallyKey": "BBR25",
	"RallyDesc":"2025 Brit Butt Rally",
	"MaxTeeshirts": 2,
	"TeeshirtSizes":["S","M","L","XL","XXL"],
	"MaxPatches": 2
}`

func intval(x string) int {

	re := regexp.MustCompile(`(\d+)`)
	sm := re.FindSubmatch([]byte(x))
	if len(sm) < 2 {
		return 0
	}
	n, _ := strconv.Atoi(string(sm[1]))
	if strings.Contains(x, "-") {
		n = 0 - n
	}
	return n

}

func json_response(w http.ResponseWriter, ok bool, msg string) {

	fmt.Fprint(w, `{"ok":`)
	if ok {
		fmt.Fprint(w, `true`)
	} else {
		fmt.Fprint(w, `false`)
	}
	fmt.Fprintf(w, `,"msg":"%v"}`, msg)
}

func json_requests(w http.ResponseWriter, r *http.Request) {

	len := intval(r.FormValue("len"))
	if len < 1 {
		len = defaultTokenSize
	}
	email := r.FormValue("email")
	token := r.FormValue("token")

	if email == "" {
		json_response(w, false, "no email supplied")
		return
	}

	if token == "" {
		token, err := OTPGenerate(MyDB, email, len) //Parameters: database, email, otp length
		if err != nil {
			json_response(w, false, "error generating token")
			return
		}
		json_response(w, true, token)
		msg := `<p>Your token is <strong><em>` + token + `</em></strong></p>`
		sendmail(email, "Hello sailor", msg)
		return
	}
	ok := OTPValid(MyDB, email, token)
	json_response(w, ok, token)
}

var htmlheader = `
<!DOCTYPE html>
<html lang="en">
<head>
<title>eforms</title>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<link rel="stylesheet" href="eforms.css">
<script src="eforms.js"></script>
</head>
<body>
`

func sendmail(email string, subj string, msg string) { // msg is used for subject and body so keep it short

	from := mail.Address{Name: emailcfg.SenderName, Address: emailcfg.SenderEmail}
	to := mail.Address{Name: email, Address: email}
	body := msg

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj
	headers["Date"] = time.Now().Format("Mon, 02 Jan 2006 15:04:05 -0700")

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	message += mime + "\r\n" + body

	// Connect to the SMTP Server
	servername := emailcfg.SMTPServer

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", emailcfg.AuthUser, emailcfg.AuthPassword, host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		fmt.Printf("Can't send email - %v\n", err)
		return
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		fmt.Printf("Can't send email - %v\n", err)
		return
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		fmt.Printf("Can't send email - %v\n", err)
		return
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		fmt.Printf("Can't send email - %v\n", err)
		return
	}

	if err = c.Rcpt(to.Address); err != nil {
		fmt.Printf("Can't send email - %v\n", err)
		return
	}

	// Data
	w, err := c.Data()
	if err != nil {
		fmt.Printf("Can't send email - %v\n", err)
		return
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		fmt.Printf("Can't send email - %v\n", err)
		return
	}

	err = w.Close()
	if err != nil {
		fmt.Printf("Can't send email - %v\n", err)
		return
	}

	c.Quit()

}

func start_signup(w http.ResponseWriter, r *http.Request) {

	var cfg EVENTCFG
	err := json.Unmarshal([]byte(BBR), &cfg)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, htmlheader)
	fmt.Fprint(w, `<div>`)
	fmt.Fprint(w, `<fieldset><label for="email">Please enter your email address</label> `)
	fmt.Fprint(w, `<input type="email" id="email" name="email" style="width:20em;"> `)
	fmt.Fprint(w, `<input type="button" value="verify" onclick="trigger_email_validation(this)">`)
	fmt.Fprint(w, ` </fieldset>`)
	fmt.Fprint(w, `</div>`)
	fmt.Fprintf(w, `<div>%v</div>`, cfg)
	fmt.Fprint(w, `</body></html>`)

}
func main() {
	var err error
	flag.Parse()
	MyDB, err = gorm.Open(sqlite.Open(*dbpath), &gorm.Config{}) // Connect any database for Postgresql, Sqlite
	if err != nil {
		panic("failed to connect to database")
	}

	fileserver := http.FileServer(http.Dir("."))
	http.Handle("/", fileserver)

	http.HandleFunc("/x", json_requests)
	http.HandleFunc("/s", start_signup)
	log.Fatal(http.ListenAndServe(":"+*port, nil))

}

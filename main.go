package main

import (
	// "log"

	"crypto/tls"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"net/mail"
	"net/smtp"
	"net/url"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbpath = flag.String("db", "./eforms.db", "Forms database to use")
var otcpath = flag.String("otc", "./otps.db", "OTC Database to use")
var port = flag.String("port", "1014", "Port to service requests")

const defaultTokenSize = 4

var MyOTC *gorm.DB
var MyDB *sql.DB

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

type EVENTMAP map[string]EVENTCFG

const BBR = `{
  "bbr25": {
    "RallyKey": "bbr25",
    "RallyDesc": "2025 Brit Butt Rally",
    "MaxTeeshirts": 2,
    "TeeshirtSizes": [
      "S",
      "M",
      "L",
      "XL",
      "XXL"
    ],
    "MaxPatches": 2
  },
  "rblr25": {
    "RallyKey": "rblr25",
    "RallyDesc": "2025 RBLR1000",
    "MaxTeeshirts": 2,
    "TeeshirtSizes": [
      "S",
      "M",
      "L",
      "XL",
      "XXL"
    ],
    "MaxPatches": 2
  }
}`

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

func fetchEvent(key string) EVENTCFG {
	var cfg EVENTMAP
	err := json.Unmarshal([]byte(BBR), &cfg)
	checkerr(err)
	for k, v := range cfg {
		if k == key {
			return v
		}
	}
	return EVENTCFG{}
}
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

func json_response(w http.ResponseWriter, ok bool, msg string, entrant int) {

	fmt.Fprint(w, `{"ok":`)
	if ok {
		fmt.Fprint(w, `true`)
	} else {
		fmt.Fprint(w, `false`)
	}
	fmt.Fprintf(w, `,"msg":"%v","entrant":%v}`, msg, entrant)
}

func json_requests(w http.ResponseWriter, r *http.Request) {

	len := intval(r.FormValue("len"))
	if len < 1 {
		len = defaultTokenSize
	}
	entrant := intval(r.FormValue("entrant"))
	email := r.FormValue("email")
	token := r.FormValue("token")
	rally := r.FormValue("rally")
	if rally == "" {
		rally = "bbr25"
	}

	if email == "" {
		json_response(w, false, "no email supplied", 0)
		return
	}

	if token == "" {
		token, err := OTPGenerate(MyOTC, email, len) //Parameters: database, email, otp length
		if err != nil {
			json_response(w, false, "error generating token", 0)
			return
		}
		json_response(w, true, "", 0)

		fmt.Println(r.Proto + " ... " + r.Host + " === " + r.URL.Host)

		cfg := fetchEvent(rally)
		msg := fmt.Sprintf(`<h1>%s</h1><p>Please verify your email by entering the code <strong><em>%s</em></strong>`, cfg.RallyDesc, token)
		msg += fmt.Sprintf(` or by <a href="http://%s/s?email=%s&token=%s&rally=%s">clicking here</a>.</p>`, r.Host, url.QueryEscape(email), url.QueryEscape(token), url.QueryEscape(rally))
		sendmail(email, "Your code is "+token, msg)
		return
	}
	ok := OTPValid(MyOTC, email, token)
	if ok {
		token, entrant = lookup_ridername_from_email(email)
		if token == "" {
			token = "ok"
		}
	}
	json_response(w, ok, token, entrant)
}

func json_lookup_iba(w http.ResponseWriter, r *http.Request) {

	f := r.FormValue("f")
	l := r.FormValue("l")
	e := r.FormValue("e")
	if f == "" || l == "" || e == "" {
		fmt.Fprint(w, `{"ok": false,"msg": "Both names not supplied"}`)
		return
	}
	iba, email := lookupIBAWeb(f, l)
	fmt.Fprint(w, `{"ok": `)
	if email == e && iba != "" {
		fmt.Fprint(w, `true`)
	} else {
		fmt.Fprint(w, `false`)
	}

}

// This is called when a code is used to verify an email address
func lookup_ridername_from_email(email string) (string, int) {

	res := ""
	n := 0
	sqlx := "SELECT ifnull(First,''),ifnull(Last,''),PersonID FROM persons WHERE Email=?"
	stmt, err := MyDB.Prepare(sqlx)
	checkerr(err)
	defer stmt.Close()
	rows, err := stmt.Query(email)
	checkerr(err)
	defer rows.Close()
	if !rows.Next() {
		return res, n
	}
	var f, l string
	err = rows.Scan(&f, &l, &n)
	checkerr(err)
	res = f + " " + l
	return res, n

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

func json_save_datafield(w http.ResponseWriter, r *http.Request) {

	const riderprefix = "Rider"
	const pillionprefix = "Pillion"

	entrant := intval(r.FormValue("e"))
	rider := intval(r.FormValue("r"))
	pillion := intval(r.FormValue("p"))

	f := r.FormValue("f")
	v := r.FormValue("v")

	if entrant < 1 {
		json_response(w, false, "no entrant supplied", entrant)
		return
	}
	sqlx := ""
	if f == "" {
		json_response(w, false, "field is blank", entrant)
		return
	}
	if slices.Contains(RiderFields, f) {
		ff := f[len(riderprefix):]
		if rider < 1 {
			// Make new record
			json_response(w, false, "new rider not implemented yet", entrant)
			return
		} else {
			sqlx = "UPDATE persons SET " + ff + "=? WHERE PersonID=" + strconv.Itoa(rider)
		}
	} else if slices.Contains(PillionFields, f) {
		ff := f[len(pillionprefix):]
		if pillion < 1 {
			// Make new record
			json_response(w, false, "new pillion not implemented yet", entrant)
			return
		} else {
			sqlx = "UPDATE persons SET " + ff + "=? WHERE PersonID=" + strconv.Itoa(pillion)
		}
	} else {
		sqlx = "UPDATE entrants SET " + f + "=? WHERE EntrantNumber=" + strconv.Itoa(entrant)
	}
	fmt.Println(sqlx)
	stmt, err := MyDB.Prepare(sqlx)
	checkerr(err)
	defer stmt.Close()
	_, err = stmt.Exec(v)
	checkerr(err)
	json_response(w, true, "ok", entrant)

}

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

func send_token_form(w http.ResponseWriter, r *http.Request, hide bool) {

	tl := intval(r.FormValue("tokenlen"))
	if tl < 1 {
		tl = defaultTokenSize
	}

	tkn := r.FormValue("token")
	fmt.Fprint(w, `<fieldset class="tokenzone`)
	if hide {
		fmt.Fprint(w, ` hide`)
	}
	fmt.Fprint(w, `">`)
	fmt.Fprintf(w, `<input type="hidden" id="tokenlen" value="%v">`, tl)
	fmt.Fprint(w, `<label for="vtchar1">Please enter the code</label> `)

	fmt.Fprint(w, `<span class="field">`)
	for i := 1; i <= tl; i++ {
		c := ""
		if len(tkn) >= i {
			c = tkn[i-1 : i]
		}
		fmt.Fprintf(w, `<input type="text" id="vtchar%v" class="verify-token" oninput="tokenInput(this)" value="%v"> `, i, c)
	}
	fmt.Fprint(w, `</span>`)
	fmt.Fprint(w, `<input type="button" id="checktoken" value="Verify" onclick="verify_email_validation(this)"> `)
	fmt.Fprint(w, ` &nbsp;&nbsp; <span id="checkresult"> </span>`)
	fmt.Fprint(w, `</fieldset>`)
	if !hide && tkn != "" {
		fmt.Fprint(w, `<script>verify_email_validation(document.getElementById('checktoken'))</script>`)
	}

}

func show_entry_form(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")
	if email == "" {
		start_signup(w, r)
		return
	}
	rally := r.FormValue("rally")

	personid := intval(r.FormValue("er"))
	if personid < 1 {
		personid = start_new_person_record(email)
	}

	er := fetch_entrant(rally, personid)
	if er.EntrantNumber < 1 {
		er.EntrantNumber = start_new_entrant_record(rally, personid)
	}

	ev := fetch_event_record(rally)
	er.Event = ev
	tprd := template.Must(template.New("tprd").Parse(tp_RiderDetails))
	tppn := template.Must(template.New("tppn").Parse(tp_PillionDetails))
	tpbk := template.Must(template.New("tpbk").Parse(tp_BikeDetails))
	tpnk := template.Must(template.New("tpnk").Parse(tp_NokDetails))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, htmlheader)
	fmt.Fprint(w, `<article class="signupform">`)
	cfg := fetchEvent(rally)
	fmt.Fprintf(w, `<h1>%v entry form</h1>`, cfg.RallyDesc)

	fmt.Fprint(w, `<form method="post" >`)
	fmt.Fprintf(w, `<input type="hidden" id="EntrantNumber" name="EntrantNumber" value="%v">`, er.EntrantNumber)
	fmt.Fprintf(w, `<input type="hidden" id="RiderID" name="RiderID" value="%v">`, er.RiderID)
	fmt.Fprintf(w, `<input type="hidden" id="PillionID" name="PillionID" value="%v">`, er.PillionID)
	err := tprd.Execute(w, er)
	checkerr(err)
	err = tpbk.Execute(w, er)
	checkerr(err)
	err = tpnk.Execute(w, er)
	checkerr(err)
	err = tppn.Execute(w, er)
	checkerr(err)
	fmt.Fprint(w, `</form>`)
	fmt.Fprint(w, `</article>`)
}

func start_new_entrant_record(rally string, riderid int) int {

	res := 1
	sqlx := "SELECT ifnull(max(EntrantNumber),0) FROM entrants WHERE EventCode=?"

	stmt, err := MyDB.Prepare(sqlx)
	checkerr(err)
	defer stmt.Close()
	rows, err := stmt.Query(rally)
	checkerr(err)
	defer rows.Close()
	if rows.Next() {
		var mx int
		err = rows.Scan(&mx)
		checkerr(err)
		res = mx + 1
	}
	rows.Close()
	stmt.Close()

	dt := time.Now()
	dtx := dt.Format(time.DateOnly)

	sqlx = "INSERT INTO entrants (EventCode,RiderID,EntrantNumber,DateCreated,DateUpdated) VALUES(?,?,?,?,?)"
	stmt, err = MyDB.Prepare(sqlx)
	checkerr(err)
	defer stmt.Close()
	_, err = stmt.Exec(rally, riderid, res, dtx, dtx)
	checkerr(err)

	return res

}

func start_new_person_record(email string) int {

	sqlx := "SELECT ifnull(max(PersonID),0) FROM persons"
	person := intval(getStringFromDB(sqlx, "0")) + 1

	sqlx = "INSERT INTO persons (PersonID,Email) VALUES(?,?)"
	stmt, err := MyDB.Prepare(sqlx)
	checkerr(err)
	defer stmt.Close()
	_, err = stmt.Exec(person, email)
	checkerr(err)
	return person

}

// This is the start of the whole procedure
func start_signup(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")
	rally := r.FormValue("rally")
	token := r.FormValue("token")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, htmlheader)
	fmt.Fprint(w, `<article class="signupform">`)
	cfg := fetchEvent(rally)
	fmt.Fprintf(w, `<h1>%v entry form</h1>`, cfg.RallyDesc)
	fmt.Fprint(w, `<fieldset><label for="email">Please enter your email address</label> `)
	fmt.Fprintf(w, `<input type="hidden" id="rally" name="rally" value="%v" onchange="retry_email(this)">`, rally)
	fmt.Fprintf(w, `<input type="hidden" id="token" name="token" value="%v">`, token)
	fmt.Fprintf(w, `<input type="email" id="email" name="email" value="%v" onchange="retry_email(this)"> `, email)
	fmt.Fprint(w, `<input type="button" id="tevbtn" disabled value="verify" onclick="trigger_email_validation(this)">`)
	fmt.Fprint(w, ` </fieldset>`)

	fmt.Fprint(w, `<fieldset class="tokenzone hide">Please check your email for a one time code then enter the code below</fieldset>`)

	send_token_form(w, r, token == "")

	fmt.Fprint(w, `<div id="confirmID" class="hide">`)
	fmt.Fprint(w, `<fieldset><label for="x"> - is this you?</label> `)
	fmt.Fprint(w, `<input type="button" id="x" value="Yes, let's do it" onclick="show_form_start()"> `)
	fmt.Fprint(w, `<input type="button" value="No, use a different email" onclick="show_signup_start()"> `)
	fmt.Fprint(w, `</div>`)

	fmt.Fprint(w, `</article>`)

	fmt.Fprint(w, `</body></html>`)

}

func main() {
	var err error
	flag.Parse()
	MyOTC, err = gorm.Open(sqlite.Open(*otcpath), &gorm.Config{}) // Connect any database for Postgresql, Sqlite
	if err != nil {
		panic("failed to connect to OTC database")
	}

	MyDB, err = sql.Open("sqlite3", *dbpath)
	if err != nil {
		panic("Failed to connect to forms database")
	}

	debug_fetcher()

	fileserver := http.FileServer(http.Dir("."))
	http.Handle("/", fileserver)

	http.HandleFunc("/l", json_lookup_iba)
	http.HandleFunc("/s", start_signup)
	http.HandleFunc("/f", show_entry_form)
	http.HandleFunc("/x", json_requests)
	http.HandleFunc("/z", json_save_datafield)
	log.Fatal(http.ListenAndServe(":"+*port, nil))

}

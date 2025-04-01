package main

import (
	// "log"

	"flag"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbpath = flag.String("db", "./otps.db", "Database to use")
var port = flag.String("port", "1014", "Port to service requests")

const defaultTokenSize = 4

var MyDB *gorm.DB

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
		return
	}
	ok := OTPValid(MyDB, email, token)
	json_response(w, ok, token)
}

func main() {
	var err error
	flag.Parse()
	MyDB, err = gorm.Open(sqlite.Open(*dbpath), &gorm.Config{}) // Connect any database for Postgresql, Sqlite
	if err != nil {
		panic("failed to connect to database")
	}

	http.HandleFunc("/x", json_requests)
	log.Fatal(http.ListenAndServe(":"+*port, nil))

}

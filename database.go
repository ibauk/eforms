package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// URL of live Rides database lookup script
const LiveLookupUrl = "https://rdb.ironbutt.co.uk/lookup.php"

type PERSON_RECORD struct {
	First            string
	Last             string
	IBA              string
	RBL              string
	Addr1            string
	Addr2            string
	Town             string
	County           string
	Postcode         string
	Country          string
	Phone            string
	Email            string
	PersonID         int
	AlternativeEmail string
}

var RiderFields = []string{"RiderFirst", "RiderLast", "RiderIBA", "RiderRBL", "RiderAddr1", "RiderAddr2", "RiderTown", "RiderCounty", "RiderPostcode", "RiderCountry", "RiderPhone", "RiderEmail", "RiderAlternativeEmail"}
var PillionFields = []string{"PillionFirst", "PillionLast", "PillionIBA", "PillionRBL", "PillionAddr1", "PillionAddr2", "PillionTown", "PillionCounty", "PillionPostcode", "PillionCountry", "PillionPhone", "PillionEmail", "PillionAlternativeEmail"}

const PersonFieldsSQL = `ifnull(First,''),ifnull(Last,''),ifnull(IBA,''),ifnull(RBL,'N'),
						ifnull(Addr1,''),ifnull(Addr2,''),ifnull(Town,''),ifnull(County,''),
						ifnull(Postcode,''),ifnull(Country,''),
						ifnull(Phone,''),ifnull(Email,''),ifnull(PersonID,0),
						ifnull(AlternativeEmail,'')
`

type ENTRANT_RECORD struct {
	EventCode       string
	EntrantNumber   int
	RecordStatus    int
	DateCreated     string
	DateUpdated     string
	RiderID         int
	PillionID       int
	Rider           PERSON_RECORD
	Pillion         PERSON_RECORD
	RiderFirst      string
	RiderLast       string
	RiderIBA        string
	RiderRBL        string
	RiderAddr1      string
	RiderAddr2      string
	RiderTown       string
	RiderCounty     string
	RiderPostcode   string
	RiderCountry    string
	RiderPhone      string
	RiderEmail      string
	RiderNoviceYN   string
	HasPillionYN    string
	PillionFirst    string
	PillionLast     string
	PillionIBA      string
	PillionRBL      string
	PillionAddr1    string
	PillionAddr2    string
	PillionTown     string
	PillionCounty   string
	PillionPostcode string
	PillionCountry  string
	PillionPhone    string
	PillionEmail    string
	PillionNoviceYN string
	Bike            string
	BikeReg         string
	OdoCountsMK     string
	NokName         string
	NokPhone        string
	NokRelation     string
	Tshirts         string
	Patches         int
	RouteClass      string
	FreeCampingYN   string
	Sponsorship     string
	PaymentMethod   string

	// Fields below include from other sources for reporting purposes

	Event EVENT_RECORD
}

type EVENT_OPTIONS struct {
	OfferCampingYN    bool
	AskDistance2Venue bool
	Ask4Sponsorship   bool
	OfferScoringEmail bool
	Ask4RBL           bool
	AskNoviceYN       bool
	OfferRouteClasses bool
}
type EVENT_RECORD struct {
	EventCode          string
	FullTitle          string
	GenericName        string
	RiderFee           string
	PillionFee         string
	TshirtFee          string
	TshirtSizes        []string
	PatchFee           string
	MaxTshirts         int
	MaxPatches         int
	RouteClasses       []string
	PaymentMethods     []string
	SponsorshipOptions []string
	Options            EVENT_OPTIONS
}

const EntrantFieldsSQL = `EventCode,EntrantNumber,RecordStatus,ifnull(DateCreated,''),ifnull(DateUpdated,''),
						ifnull(RiderID,0),ifnull(PillionID,0),ifnull(RiderNoviceYN,'N'),
						ifnull(HasPillionYN,''),ifnull(PillionNoviceYN,'N'),
						ifnull(Bike,''),ifnull(BikeReg,''),ifnull(OdoCountsMK,'M'),
						ifnull(NokName,''),ifnull(NokPhone,''),ifnull(NokRelation,''),
						ifnull(Tshirts,''),ifnull(Patches,0),ifnull(RouteClass,''),ifnull(FreeCampingYN,'N'),
						ifnull(Sponsorship,''),ifnull(PaymentMethod,'')
						`

func debug_fetcher() {

	sqlx := "SELECT " + EntrantFieldsSQL + " FROM entrants WHERE EntrantNumber=1"

	er := fetch_entrant_record(sqlx)
	fmt.Printf("Entrant: %v\n", er)
	evt := fetch_event_record("rblr25")
	fmt.Printf("Event: %v\n", evt)
}

func fetch_entrant(rally string, riderid int) ENTRANT_RECORD {

	sqlx := "SELECT " + EntrantFieldsSQL + " FROM entrants WHERE EventCode='" + rally + "' AND RiderID='" + strconv.Itoa(riderid) + "'"

	return fetch_entrant_record(sqlx)
}

func fetch_entrant_record(sqlx string) ENTRANT_RECORD {

	var er ENTRANT_RECORD

	rows, err := MyDB.Query(sqlx)
	checkerr(err)
	defer rows.Close()
	if rows.Next() {
		scan_entrant_record(rows, &er)
		if er.RiderID > 0 {
			er.Rider = fetch_person_record(er.RiderID)
		}
		if er.PillionID > 0 {
			er.Pillion = fetch_person_record(er.PillionID)
		}
	}
	return er

}

func fetch_event_record(EventCode string) EVENT_RECORD {

	var ev EVENT_RECORD
	var cfg []byte

	ev.EventCode = EventCode

	sqlx := "SELECT Config FROM events WHERE EventCode='" + EventCode + "'"
	rows, err := MyDB.Query(sqlx)
	checkerr(err)
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&cfg)
		checkerr(err)
		err = json.Unmarshal(cfg, &ev)
		checkerr(err)
	}
	return ev
}

func fetch_person_record(PersonID int) PERSON_RECORD {

	var p PERSON_RECORD

	p.PersonID = PersonID

	sqlx := "SELECT " + PersonFieldsSQL + " FROM persons WHERE PersonID=" + strconv.Itoa(PersonID)
	rows, err := MyDB.Query(sqlx)
	checkerr(err)
	defer rows.Close()
	if !rows.Next() {
		return p
	}
	err = rows.Scan(&p.First, &p.Last, &p.IBA, &p.RBL, &p.Addr1, &p.Addr2, &p.Town, &p.County, &p.Postcode, &p.Country, &p.Phone, &p.Email, &p.PersonID, &p.AlternativeEmail)
	checkerr(err)
	return p
}

func getStringFromDB(sqlx string, defval string) string {

	res := defval
	rows, err := MyDB.Query(sqlx)
	checkerr(err)
	defer rows.Close()
	if !rows.Next() {
		return res
	}
	err = rows.Scan(&res)
	checkerr(err)
	return res
}

func lookupIBAWeb(first, last string) (string, string) {

	type LookupResponse struct {
		Iba   string
		Sname string
		Email string
	}
	var lresp LookupResponse
	var client http.Client

	url := LiveLookupUrl + "?f=" + url.QueryEscape(first) + "&l=" + url.QueryEscape(last)

	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("*** can't access online members database\n*** %v\n", err)
		return "", ""
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("*** can't access online members database\n*** %v\n", err)
			return "", ""
		}
		//bodyString := string(bodyBytes)
		json.Unmarshal(bodyBytes, &lresp)
		//fmt.Printf("%v\n", bodyString)
	}
	return lresp.Iba, lresp.Email
}

func scan_entrant_record(rows *sql.Rows, er *ENTRANT_RECORD) {

	err := rows.Scan(&er.EventCode, &er.EntrantNumber, &er.RecordStatus, &er.DateCreated, &er.DateUpdated, &er.RiderID, &er.PillionID, &er.RiderNoviceYN, &er.HasPillionYN, &er.PillionNoviceYN, &er.Bike, &er.BikeReg, &er.OdoCountsMK, &er.NokName, &er.NokPhone, &er.NokRelation, &er.Tshirts, &er.Patches, &er.RouteClass, &er.FreeCampingYN, &er.Sponsorship, &er.PaymentMethod)
	checkerr(err)
}

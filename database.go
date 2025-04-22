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
	ScoringEmail    string

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

const RBLR_Definition = `{
  "FullTitle": "2025 RBLR1000",
  "GenericName": "RBLR1000",
  "RiderFee": "90",
  "PillionFee": "10",
  "TshirtFee": "15",
  "TshirtSizes": [
    "S",
    "M",
    "L",
    "XL",
    "XXL"
  ],
  "MaxTshirts": 2,
  "MaxPatches": 2,
  "PatchFee": "5",
  "RouteClasses": [
    "A-North Clockwise",
    "B-North Anticlockwise",
    "C-South Clockwise",
    "D-South Anticlockwise",
    "E-500 Clockwise",
    "F-500 Anticlockwise"
  ],
  "VenueName": "Squires",
  "DistanceUnit": "miles",
  "SponsorshipOptions": [
    "Include £25 now",
    "Include £50 now",
    "Include £75 now",
    "Include £100 now",
    "I'll bring a cheque to Squires"
  ],
  "Options": {
    "OfferCampingYN": true,
    "AskDistance2Venue": true,
    "Ask4Sponsorship": true,
    "OfferScoringEmail": false,
    "Ask4RBL": true,
    "AskNoviceYN": true,
	"OfferRouteClasses": true
  }
}`

const BBR_Definition = `{
  "FullTitle": "2025 Park & Ride Coddiwomple",
  "GenericName": "Coddiwomple",
  "RiderFee": "90",
  "PillionFee": "10",
  "TshirtFee": "15",
  "TshirtSizes": [
    "S",
    "M",
    "L",
    "XL",
    "XXL"
  ],
  "MaxTshirts": 2,
  "MaxPatches": 0,
  "PatchFee": "5",
  "VenueName": "Squires",
  "DistanceUnit": "miles",
  "Options": {
    "OfferCampingYN": false,
    "AskDistance2Venue": false,
    "Ask4Sponsorship": false,
    "OfferScoringEmail": true,
    "Ask4RBL": false,
    "AskNoviceYN": true,
	"OfferRouteClasses": false
  }
}
`

/*
*
const EntrantFields = `EventCode,EntrantNumber,RecordStatus,ifnull(DateCreated,”),ifnull(DateUpdated,”),

	RiderID,PillionID
	ifnull(RiderFirst,''),ifnull(RiderLast,''),ifnull(RiderIBA,''),ifnull(RiderRBL,0),
	ifnull(RiderAddr1,''),ifnull(RiderAddr2,''),ifnull(RiderTown,''),ifnull(RiderCounty,''),ifnull(RiderPostcode,''),ifnull(RiderCountry,''),
	ifnull(RiderPhone,''),ifnull(RiderEmail,''),ifnull(RiderNoviceYN,'N'),
	ifnull(HasPillionYN,''),
	ifnull(PillionFirst,''),ifnull(PillionLast,''),ifnull(PillionIBA,''),ifnull(PillionRBL,0),
	ifnull(PillionAddr1,''),ifnull(PillionAddr2,''),ifnull(PillionTown,''),ifnull(PillionCounty,''),
	ifnull(PillionPostcode,''),ifnull(PillionCountry,''),
	ifnull(PillionPhone,''),ifnull(PillionEmail,''),ifnull(PillionNoviceYN,'N'),
	ifnull(Bike,''),ifnull(BikeReg,''),ifnull(OdoCountsMK,'M'),
	ifnull(NokName,''),ifnull(NokPhone,''),ifnull(NokRelation,''),
	ifnull(Tshirts,''),ifnull(Patches,0),ifnull(RouteClass,''),ifnull(FreeCampingYN,'N'),
	ifnull(Sponsorship,''),ifnull(PaymentMethod,''),ifnull(ScoringEmail,'')
	`

*
*/
const EntrantFields = `EventCode,EntrantNumber,RecordStatus,ifnull(DateCreated,''),ifnull(DateUpdated,''),
						RiderID,PillionID
						ifnull(HasPillionYN,''),
						ifnull(Bike,''),ifnull(BikeReg,''),ifnull(OdoCountsMK,'M'),
						ifnull(NokName,''),ifnull(NokPhone,''),ifnull(NokRelation,''),
						ifnull(Tshirts,''),ifnull(Patches,0),ifnull(RouteClass,''),ifnull(FreeCampingYN,'N'),
						ifnull(Sponsorship,''),ifnull(PaymentMethod,''),ifnull(ScoringEmail,'')
						`

func debug_fetcher() {

	sqlx := "SELECT " + EntrantFields + " FROM entrants WHERE EntrantNumber=1"

	er := fetch_entrant_record(sqlx)
	fmt.Printf("Entrant: %v\n", er)
	evt := fetch_event_record("rblr25")
	fmt.Printf("Event: %v\n", evt)
}

func fetch_entrant(rally string, riderid int) ENTRANT_RECORD {

	sqlx := "SELECT " + EntrantFields + " FROM entrants WHERE EventCode='" + rally + "' AND RiderID='" + strconv.Itoa(riderid) + "'"

	return fetch_entrant_record(sqlx)
}

func fetch_entrant_record(sqlx string) ENTRANT_RECORD {

	var er ENTRANT_RECORD

	rows, err := MyDB.Query(sqlx)
	checkerr(err)
	defer rows.Close()
	if rows.Next() {
		scan_entrant_record(rows, &er)
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

	err := rows.Scan(&er.EventCode, &er.EntrantNumber, &er.RecordStatus, &er.DateCreated, &er.DateUpdated, &er.RiderFirst, &er.RiderLast, &er.RiderIBA, &er.RiderRBL, &er.RiderAddr1, &er.RiderAddr2, &er.RiderTown, &er.RiderCounty, &er.RiderPostcode, &er.RiderCountry, &er.RiderPhone, &er.RiderEmail, &er.RiderNoviceYN, &er.HasPillionYN, &er.PillionFirst, &er.PillionLast, &er.PillionIBA, &er.PillionRBL, &er.PillionAddr1, &er.PillionAddr2, &er.PillionTown, &er.PillionCounty, &er.PillionPostcode, &er.PillionCountry, &er.PillionPhone, &er.PillionEmail, &er.PillionNoviceYN, &er.Bike, &er.BikeReg, &er.OdoCountsMK, &er.NokName, &er.NokPhone, &er.NokRelation, &er.Tshirts, &er.Patches, &er.RouteClass, &er.FreeCampingYN, &er.Sponsorship, &er.PaymentMethod, &er.ScoringEmail)
	checkerr(err)
}

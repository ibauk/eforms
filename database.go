package main

import (
	"database/sql"
	"fmt"
)

type ENTRANT_RECORD struct {
	EventCode       string
	EntrantNumber   int
	RecordStatus    int
	DateCreated     string
	DateUpdated     string
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

type EVENT_RECORD struct {
	EventCode          string
	EventTitle         string
	NoviceEvent        string
	RiderFee           string
	PillionFee         string
	TshirtFee          string
	PatchFee           string
	MaxTshirts         int
	MaxPatches         int
	RouteClasses       string
	OfferCampingYN     string
	PaymentMethods     string
	SponsorshipOptions string
}

const EntrantFields = `EventCode,EntrantNumber,RecordStatus,DateCreated,ifnull(DateUpdated,''),RiderFirst,RiderLast,ifnull(RiderIBA,''),ifnull(RiderRBL,0),
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
						ifnull(Sponsorship,''),ifnull(PaymentMethod,'')
						`

func debug_fetcher() {

	sqlx := "SELECT " + EntrantFields + " FROM entrants WHERE EntrantNumber=1"

	er := fetch_entrant_record(sqlx)
	fmt.Printf("%v\n", er)
}

func fetch_entrant(rally string, email string) ENTRANT_RECORD {

	sqlx := "SELECT " + EntrantFields + " FROM entrants WHERE EventCode='" + rally + "' AND RiderEmail='" + email + "'"

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

func scan_entrant_record(rows *sql.Rows, er *ENTRANT_RECORD) {

	err := rows.Scan(&er.EventCode, &er.EntrantNumber, &er.RecordStatus, &er.DateCreated, &er.DateUpdated, &er.RiderFirst, &er.RiderLast, &er.RiderIBA, &er.RiderRBL, &er.RiderAddr1, &er.RiderAddr2, &er.RiderTown, &er.RiderCounty, &er.RiderPostcode, &er.RiderCountry, &er.RiderPhone, &er.RiderEmail, &er.RiderNoviceYN, &er.HasPillionYN, &er.PillionFirst, &er.PillionLast, &er.PillionIBA, &er.PillionRBL, &er.PillionAddr1, &er.PillionAddr2, &er.PillionTown, &er.PillionCounty, &er.PillionPostcode, &er.PillionCountry, &er.PillionPhone, &er.PillionEmail, &er.PillionNoviceYN, &er.Bike, &er.BikeReg, &er.OdoCountsMK, &er.NokName, &er.NokPhone, &er.NokRelation, &er.Tshirts, &er.Patches, &er.RouteClass, &er.FreeCampingYN, &er.Sponsorship, &er.PaymentMethod)
	checkerr(err)
}

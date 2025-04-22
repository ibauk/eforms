BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "entrants" (
	"EventCode"	TEXT NOT NULL,
	"EntrantNumber"	INTEGER NOT NULL DEFAULT 0,
	"RecordStatus"	INTEGER NOT NULL DEFAULT 0,
	"DateCreated"	TEXT NOT NULL,
	"DateUpdated"	TEXT,
	"RiderFirst"	TEXT,
	"RiderLast"	TEXT,
	"RiderIBA"	TEXT,
	"RiderRBL"	TEXT DEFAULT 'N',
	"RiderAddr1"	TEXT,
	"RiderAddr2"	TEXT,
	"RiderTown"	TEXT,
	"RiderCounty"	TEXT,
	"RiderPostcode"	TEXT,
	"RiderCountry"	TEXT,
	"RiderPhone"	TEXT,
	"RiderEmail"	TEXT,
	"RiderNoviceYN"	TEXT NOT NULL DEFAULT 'N',
	"HasPillionYN"	TEXT NOT NULL DEFAULT 'N',
	"PillionFirst"	TEXT,
	"PillionLast"	TEXT,
	"PillionIBA"	TEXT,
	"PillionRBL"	TEXT DEFAULT 'N',
	"PillionAddr1"	TEXT,
	"PillionAddr2"	TEXT,
	"PillionTown"	TEXT,
	"PillionCounty"	TEXT,
	"PillionPostcode"	TEXT,
	"PillionCountry"	TEXT,
	"PillionPhone"	TEXT,
	"PillionEmail"	TEXT,
	"PillionNoviceYN"	TEXT NOT NULL DEFAULT 'N',
	"Bike"	TEXT,
	"BikeReg"	TEXT,
	"OdoCountsMK"	TEXT NOT NULL DEFAULT 'M',
	"NokName"	TEXT,
	"NokPhone"	TEXT,
	"NokRelation"	TEXT,
	"Tshirts"	TEXT,
	"Patches"	INTEGER NOT NULL DEFAULT 0,
	"RouteClass"	TEXT NOT NULL DEFAULT '',
	"FreeCampingYN"	TEXT NOT NULL DEFAULT 'N',
	"Sponsorship"	TEXT,
	"PaymentMethod"	TEXT NOT NULL DEFAULT 'PAYPAL',
	"ScoringEmail"	TEXT,
	"RiderID"	INTEGER,
	"PillionID"	INTEGER
);
CREATE TABLE IF NOT EXISTS "events" (
	"EventCode"	TEXT NOT NULL,
	"Config"	TEXT NOT NULL DEFAULT '{}'
);
CREATE TABLE IF NOT EXISTS "persons" (
	"First"	TEXT,
	"Last"	TEXT,
	"IBA"	TEXT,
	"RBL"	TEXT,
	"Addr1"	TEXT,
	"Addr2"	TEXT,
	"Town"	TEXT,
	"County"	TEXT,
	"Postcode"	TEXT,
	"Country"	TEXT,
	"Phone"	TEXT,
	"Email"	TEXT NOT NULL,
	"PersonID"	INTEGER NOT NULL,
	"AlternativeEmail"	TEXT
);
CREATE UNIQUE INDEX IF NOT EXISTS "email" ON "persons" (
	"Email"	ASC
);
CREATE UNIQUE INDEX IF NOT EXISTS "personid" ON "persons" (
	"PersonID"	ASC
);
COMMIT;

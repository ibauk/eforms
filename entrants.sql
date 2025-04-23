BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "entrants" (
	"EventCode"	TEXT NOT NULL,
	"EntrantNumber"	INTEGER NOT NULL DEFAULT 0,
	"RecordStatus"	INTEGER NOT NULL DEFAULT 0,
	"DateCreated"	TEXT NOT NULL,
	"DateUpdated"	TEXT,
	"RiderID"	INTEGER,
	"PillionID"	INTEGER,
	"RiderNoviceYN"	TEXT NOT NULL DEFAULT 'N',
	"HasPillionYN"	TEXT NOT NULL DEFAULT 'N',
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
	"PaymentMethod"	TEXT NOT NULL DEFAULT 'PAYPAL'
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
INSERT INTO "events" VALUES ('bbr25','{  "FullTitle": "2025 Park & Ride Coddiwomple",
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
}');
INSERT INTO "events" VALUES ('rblr25','{  "FullTitle": "2025 RBLR1000",
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
    "I''ll bring a cheque to Squires"
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
}');
INSERT INTO "persons" VALUES ('Bruce','Wayne',NULL,NULL,'14 Woodfield','Kingsley','Bordon','','GU35 9NB','UK','123','stammers.bob@gmail.com',23,'bobby@baby.com');
INSERT INTO "persons" VALUES ('Gillian','Anderson',NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,'webmaster@ironbutt.co.uk',32,NULL);
CREATE UNIQUE INDEX IF NOT EXISTS "email" ON "persons" (
	"Email"	ASC
);
CREATE UNIQUE INDEX IF NOT EXISTS "personid" ON "persons" (
	"PersonID"	ASC
);
COMMIT;

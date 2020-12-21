# EFORMS DATABASE TABLES

## ef_events

Field       | Type | Default | Description
---         | ---  | ---     | ---
event_type  | int  | 0       | 0 = Rally
event_code  | txt  | n/a     | The unique short code for this event. BBR21, Inv21, ...
event_name  | txt  | n/a     | Full event title - 'Brit Butt Rally 2021'
rider_fee   | num | 0.00 | Entry fee for rider
pillion_fee | num | 0.00 | Entry fee for pillion

## eve_entrants

Field           | Type  | Default | Description
---             | ---   | ---     | ---
event_code      | txt   | n/a     | Unique short code of the event entered
entrant_number  | int   | a/i     | Serial number of this entrant within this event_code
record_status   | int   | 0       | 0 = normal live record; 1 = suspended/cancelled
date_created    | date  |         | Date this record is first created
date_updated    | date  |         | Date last updated or null
rider_first     | txt   |         | First name of rider
rider_last      | txt   |         | Last name of rider
rider_iba       | txt   |         | IBA number of rider or null
address1        | txt   |         | First line of postal address
address2        | txt   |         | Second line of postal address
address_city    | txt   |         | Postal city
address_state   | txt   |         | Postal state or county
address_postalcode | txt |        | Postcode, Zip code, etc
country         | txt   | UK      | Country of origin
rider_phone     | txt   |         | Phone number of rider for use during the event
rider_email     | txt   |         | Rider's email address
has_pillion     | int   | 0       | 0 = rider only; 1 = rider+pillion
pillion_first   | txt   |         | Pillion's first name
pillion_last    | txt   |         | Pillion's last name
pillion_iba     | txt   |         | Pillion's IBA number, if known
bike            | txt   |         | Bike make and model
bikereg         | txt   |         | Registration of bike
nok_name        | txt   |         | Name of next of kin
nok_phone       | txt   |         | Phone for next of kin
nok_relation    | txt   |         | Relationship of NOK to rider
claim_method    | int   | 0       | 0 = EBC; 1 = non-EBC
payment_method  | int   | 0       | 0 = paypal; 1 = FOC

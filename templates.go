package main

const tp_RiderDetails = `
<div class="RiderDetails">

<fieldset class="group"><legend>Rider details</legend>

<fieldset class="reqd">
<label for="RiderFirst">Name</label>
<input type="text" class="firstname" id="RiderFirst" name="RiderFirst" placeholder="First" value="{{.Rider.First}}" oninput="oi(this)" onchange="oc(this)">
<input type="text" class="lastname" id="RiderLast" name="RiderLast" placeholder="Last" value="{{.Rider.Last}}" oninput="oi(this)" onchange="oc(this)">
</fieldset>

<fieldset class="address">
<label for="RiderAddr1">Postal address</label>
<input type="text" class="addr1" id="RiderAddr1" name="RiderAddr1" value="{{.Rider.Addr1}}" oninput="oi(this)" onchange="oc(this)">
<input type="text" class="addr2" id="RiderAddr2" name="RiderAddr2" value="{{.Rider.Addr2}}" oninput="oi(this)" onchange="oc(this)">
<input type="text" class="town" id="RiderTown" name="RiderTown" value="{{.Rider.Town}}" oninput="oi(this)" onchange="oc(this)">
<input type="text" class="county" id="RiderCounty" name="RiderCounty" value="{{.Rider.County}}" oninput="oi(this)" onchange="oc(this)">
<input type="text" class="postcode" id="RiderPostcode" name="RiderPostcode" value="{{.Rider.Postcode}}" oninput="oi(this)" onchange="oc(this)">
<input type="text" class="country" id="RiderCountry" name="RiderCountry" value="{{.Rider.Country}}" oninput="oi(this)" onchange="oc(this)">
</fieldset>

<fieldset>
<label for="RiderPhone">Mobile phone</label>
<input type="text" class="phone" id="RiderPhone" name="RiderPhone" value="{{.Rider.Phone}}" oninput="oi(this)" onchange="oc(this)">
</fieldset>

<fieldset>
<label for="RiderEmail">Email</label>
<input type="email" id="RiderEmail" name="RiderEmail" readonly value="{{.Rider.Email}}" oninput="oi(this)" onchange="oc(this)">
</fieldset>
{{if .Event.Options.OfferScoringEmail}}
<fieldset>
<label for="ScoringEmail">Alternative email</label>
<input type="email" placeholder="May be used for scoring purposes" id="ScoringEmail" name="RiderAlternativeEmail" value="{{.Rider.AlternativeEmail}}" oninput="oi(this)" onchange="oc(this)">
</fieldset>
{{end}}



<fieldset class="hide">
<label for="RiderIBA">Rider's IBA number</label>
<input type="text" class="ibanumber" id="RiderIBA" name="RiderIBA" value="{{.Rider.IBA}}" oninput="oi(this)" onchange="oc(this)">
</fieldset>

{{if .Event.Options.Ask4RBL}}
<fieldset>
<label for="RiderRBL">Royal British Legion</label>
<select id="RiderRBL" name="RiderRBL" oninput="oi(this)" onchange="oc(this)">
<option value="" {{if eq .Rider.RBL ""}}selected{{end}}>Please choose an option</option>
<option value="N" {{if eq .Rider.RBL "N"}}selected{{end}}>No Legion association</option>
<option value="L" {{if eq .Rider.RBL "L"}}selected{{end}}>I'm an ordinary Legion member</option>
<option value="R" {{if eq .Rider.RBL "R"}}selected{{end}}>I am a Legion Rider (RBLR) member</option>
</select>
</fieldset>
{{end}}

{{if .Event.Options.AskNoviceYN}}
<fieldset>
<label for="RiderNoviceYN">Is this your first {{.Event.GenericName}}?</label>
<select id="RiderNoviceYN" name="RiderNoviceYN" oninput="oi(this)" onchange="oc(this)">
<option value="N" {{if eq .RiderNoviceYN "N"}}selected{{end}}>No, not my first time</option>
<option value="Y" {{if eq .RiderNoviceYN "Y"}}selected{{end}}>Yes, I'm a novice</option>
</select>
</fieldset>
{{end}}

</fieldset>

</div>
`

const tp_PillionDetails = `
<div class="OptionalPillion">
<fieldset>
<label for="HasPillionYN">Are you riding with a pillion?</label>
<select id="HasPillionYN" name="HasPillionYN" onchange="flipPillionDetails(this)" oninput="oi(this)" onchange="oc(this)">
<option value="N" {{if eq .HasPillionYN "N"}}selected{{end}}>No pillion, solo rider</option>
<option value="Y" {{if eq .HasPillionYN "Y"}}selected{{end}}>Yes, we're two up</option>
</select>
</fieldset>
<div id="divPillionDetails" class="PillionDetails {{if eq .HasPillionYN "N"}}hide{{end}}">


<fieldset class="group"><legend>Pillion details</legend>

<fieldset>
<label for="PillionEmail">Email</label>
<input type="email" id="PillionEmail" name="PillionEmail" 
{{if eq "" .Pillion.Email}}
 data-valid="0"
{{else}}
 data-valid="1"
{{end}} value="{{.Pillion.Email}}" oninput="oi(this)" onchange="oc(this)">
</fieldset>


<fieldset class="reqd">
<label for="PillionFirst">Name</label>
<input type="text" class="firstname" id="PillionFirst" name="PillionFirst" placeholder="First" value="{{.Pillion.First}}" oninput="oi(this)" onchange="oc(this)">
<input type="text" class="lastname" id="PillionLast" name="PillionLast" placeholder="Last" value="{{.Pillion.Last}}" oninput="oi(this)" onchange="oc(this)">
</fieldset>

<fieldset class="address">
<label for="PillionAddr1">Postal address</label>
<input type="text" class="addr1" id="PillionAddr1" name="PillionAddr1" placeholder="Street address" value="{{.Pillion.Addr1}}" oninput="oi(this)" onchange="oc(this)">
<input type="text" class="addr2" id="PillionAddr2" name="PillionAddr2" placeholder="Address line 2" value="{{.Pillion.Addr2}}" oninput="oi(this)" onchange="oc(this)">
<input type="text" class="town" id="PillionTown" name="PillionTown" placeholder="Town/city" value="{{.Pillion.Town}}" oninput="oi(this)" onchange="oc(this)">
<input type="text" class="county" id="PillionCounty" name="PillionCounty" placeholder="County/region" value="{{.Pillion.County}}" oninput="oi(this)" onchange="oc(this)">
<input type="text" class="postcode" id="PillionPostcode" name="PillionPostcode" placeholder="Postcode" value="{{.Pillion.Postcode}}" oninput="oi(this)" onchange="oc(this)">
<input type="text" class="country" id="PillionCountry" name="PillionCountry" placeholder="Country" value="{{.Pillion.Country}}" oninput="oi(this)" onchange="oc(this)">
</fieldset>

<fieldset>
<label for="PillionPhone">Mobile phone</label>
<input type="text" class="phone" id="PillionPhone" name="PillionPhone" value="{{.Pillion.Phone}}" oninput="oi(this)" onchange="oc(this)">
</fieldset>

<fieldset class="hide">
<label for="PillionIBA">Pillion's IBA number</label>
<input type="text" class="ibanumber" id="PillionIBA" name="PillionIBA" value="{{.Pillion.IBA}}" oninput="oi(this)" onchange="oc(this)">
</fieldset>

{{if .Event.Options.Ask4RBL}}
<fieldset>
<label for="PillionRBL">Royal British Legion</label>
<select id="PillionRBL" name="PillionRBL" oninput="oi(this)" onchange="oc(this)">
<option value="" {{if eq .Pillion.RBL ""}}selected{{end}}>Please choose an option</option>
<option value="N" {{if eq .Pillion.RBL "N"}}selected{{end}}>No Legion association</option>
<option value="L" {{if eq .Pillion.RBL "L"}}selected{{end}}>I'm an ordinary Legion member</option>
<option value="R" {{if eq .Pillion.RBL "R"}}selected{{end}}>I am a Legion Pillion (RBLR) member</option>
</select>
</fieldset>
{{end}}

{{if .Event.Options.AskNoviceYN}}
<fieldset>
<label for="PillionNoviceYN">Is this your first {{.Event.GenericName}}?</label>
<select id="PillionNoviceYN" name="PillionNoviceYN" oninput="oi(this)" onchange="oc(this)">
<option value="N" {{if eq .PillionNoviceYN "N"}}selected{{end}}>No, not my first time</option>
<option value="Y" {{if eq .PillionNoviceYN "Y"}}selected{{end}}>Yes, I'm a novice</option>
</select>
</fieldset>
{{end}}





</fieldset>
</div>
</div>
`

const tp_BikeDetails = `
<div class="BikeDetails">
<fieldset>
<label for="Bike">Bike</label>
<input type="text" class="Bike" id="Bike" name="Bike" placeholder="Make & model" value="{{.Bike}}" oninput="oi(this)" onchange="oc(this)">
<input type="text" class="BikeReg" id="BikeReg" name="BikeReg" placeholder="Reg." value="{{.BikeReg}}" oninput="oi(this)" onchange="oc(this)">

<span>
<label for="OdoCountsMK">Odometer counts</label>
<select id="OdoCountsMK" name="OdoCountsMK" oninput="oi(this)" onchange="oc(this)">
<option value="M" {{if eq .OdoCountsMK "M"}}selected{{end}}>miles</option>
<option value="K" {{if eq .OdoCountsMK "K"}}selected{{end}}>Kilometres</option>
</select>
</span>
</fieldset>
</div>
`

const tp_NokDetails = `
<div class="NokDetails">
<fieldset class="address">
<label for="NokName">Emergency contact details</label>
<input type="text" placeholder="Name" class="NokName" id="NokName" name="NokName" value="{{.NokName}}" oninput="oi(this)" onchange="oc(this)">
<input type="text" placeholder="Phone" class="NokPhone" id="NokPhone" name="NokPhone" value="{{.NokPhone}}" oninput="oi(this)" onchange="oc(this)">
<input type="text" placeholder="Relationship to rider" class="NokRelation" id="NokRelation" name="NokRelation" value="{{.NokRelation}}" oninput="oi(this)" onchange="oc(this)">
</fieldset>
</div>
`

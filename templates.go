package main

const tp_RiderDetails = `
<div class="RiderDetails">

<fieldset class="group"><legend>Rider details</legend>

<fieldset class="reqd">
<label for="RiderFirst">Name</label>
<input type="text" class="firstname" id="RiderFirst" name="RiderFirst" placeholder="First" value="{{.RiderFirst}}">
<input type="text" class="lastname" id="RiderLast" name="RiderLast" placeholder="Last" value="{{.RiderLast}}">
</fieldset>

<fieldset class="address">
<label for="RiderAddr1">Postal address</label>
<input type="text" class="addr1" id="RiderAddr1" name="RiderAddr1" value="{{.RiderAddr1}}">
<input type="text" class="addr2" id="RiderAddr2" name="RiderAddr2" value="{{.RiderAddr2}}">
<input type="text" class="town" id="RiderTown" name="RiderTown" value="{{.RiderTown}}">
<input type="text" class="county" id="RiderCounty" name="RiderCounty" value="{{.RiderCounty}}">
<input type="text" class="postcode" id="RiderPostcode" name="RiderPostcode" value="{{.RiderPostcode}}">
<input type="text" class="country" id="RiderCountry" name="RiderCountry" value="{{.RiderCountry}}">
</fieldset>

<fieldset>
<label for="RiderPhone">Mobile phone</label>
<input type="text" class="phone" id="RiderPhone" name="RiderPhone" value="{{.RiderPhone}}">
</fieldset>

<fieldset>
<label for="RiderEmail">Email</label>
<input type="email" id="RiderEmail" name="RiderEmail" readonly value="{{.RiderEmail}}">
</fieldset>
{{if .Event.Options.OfferScoringEmail}}
<fieldset>
<label for="ScoringEmail">Alternative email</label>
<input type="email" placeholder="May be used for scoring purposes" id="ScoringEmail" name="ScoringEmail" value="{{.ScoringEmail}}">
</fieldset>
{{end}}



<fieldset class="hide">
<label for="RiderIBA">Rider's IBA number</label>
<input type="text" class="ibanumber" id="RiderIBA" name="RiderIBA" value="{{.RiderIBA}}">
</fieldset>

{{if .Event.Options.Ask4RBL}}
<fieldset>
<label for="RiderRBL">Royal British Legion</label>
<select id="RiderRBL" name="RiderRBL">
<option value="" {{if eq .RiderRBL ""}}selected{{end}}>Please choose an option</option>
<option value="N" {{if eq .RiderRBL "N"}}selected{{end}}>No Legion association</option>
<option value="L" {{if eq .RiderRBL "L"}}selected{{end}}>I'm an ordinary Legion member</option>
<option value="R" {{if eq .RiderRBL "R"}}selected{{end}}>I am a Legion Rider (RBLR) member</option>
</select>
</fieldset>
{{end}}

{{if .Event.Options.AskNoviceYN}}
<fieldset>
<label for="RiderNoviceYN">Is this your first {{.Event.GenericName}}?</label>
<select id="RiderNoviceYN" name="RiderNoviceYN">
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
<select id="HasPillionYN" name="HasPillionYN" onchange="flipPillionDetails(this)">
<option value="N" {{if eq .HasPillionYN "N"}}selected{{end}}>No pillion, solo rider</option>
<option value="Y" {{if eq .HasPillionYN "Y"}}selected{{end}}>Yes, we're two up</option>
</select>
</fieldset>
<div id="divPillionDetails" class="PillionDetails {{if eq .HasPillionYN "N"}}hide{{end}}">


<fieldset class="group"><legend>Pillion details</legend>

<fieldset class="reqd">
<label for="PillionFirst">Name</label>
<input type="text" class="firstname" id="PillionFirst" name="PillionFirst" placeholder="First" value="{{.PillionFirst}}">
<input type="text" class="lastname" id="PillionLast" name="PillionLast" placeholder="Last" value="{{.PillionLast}}">
</fieldset>

<fieldset class="address">
<label for="PillionAddr1">Postal address</label>
<input type="text" class="addr1" id="PillionAddr1" name="PillionAddr1" placeholder="Street address" value="{{.PillionAddr1}}">
<input type="text" class="addr2" id="PillionAddr2" name="PillionAddr2" placeholder="Address line 2" value="{{.PillionAddr2}}">
<input type="text" class="town" id="PillionTown" name="PillionTown" placeholder="Town/city" value="{{.PillionTown}}">
<input type="text" class="county" id="PillionCounty" name="PillionCounty" placeholder="County/region" value="{{.PillionCounty}}">
<input type="text" class="postcode" id="PillionPostcode" name="PillionPostcode" placeholder="Postcode" value="{{.PillionPostcode}}">
<input type="text" class="country" id="PillionCountry" name="PillionCountry" placeholder="Country" value="{{.PillionCountry}}">
</fieldset>

<fieldset>
<label for="PillionPhone">Mobile phone</label>
<input type="text" class="phone" id="PillionPhone" name="PillionPhone" value="{{.PillionPhone}}">
</fieldset>

<fieldset>
<label for="PillionEmail">Email</label>
<input type="email" id="PillionEmail" name="PillionEmail" value="{{.PillionEmail}}">
</fieldset>

<fieldset class="hide">
<label for="PillionIBA">Pillion's IBA number</label>
<input type="text" class="ibanumber" id="PillionIBA" name="PillionIBA" value="{{.PillionIBA}}">
</fieldset>

{{if .Event.Options.Ask4RBL}}
<fieldset>
<label for="PillionRBL">Royal British Legion</label>
<select id="PillionRBL" name="PillionRBL">
<option value="" {{if eq .PillionRBL ""}}selected{{end}}>Please choose an option</option>
<option value="N" {{if eq .PillionRBL "N"}}selected{{end}}>No Legion association</option>
<option value="L" {{if eq .PillionRBL "L"}}selected{{end}}>I'm an ordinary Legion member</option>
<option value="R" {{if eq .PillionRBL "R"}}selected{{end}}>I am a Legion Pillion (RBLR) member</option>
</select>
</fieldset>
{{end}}

{{if .Event.Options.AskNoviceYN}}
<fieldset>
<label for="PillionNoviceYN">Is this your first {{.Event.GenericName}}?</label>
<select id="PillionNoviceYN" name="PillionNoviceYN">
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
<input type="text" class="Bike" id="Bike" name="Bike" placeholder="Make & model" value="{{.Bike}}">
<input type="text" class="BikeReg" id="BikeReg" name="BikeReg" placeholder="Reg." value="{{.BikeReg}}">

<span>
<label for="OdoCountsMK">Odometer counts</label>
<select id="OdoCountsMK" name="OdoCountsMK">
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
<input type="text" placeholder="Name" class="NokName" id="NokName" name="NokName" value="{{.NokName}}">
<input type="text" placeholder="Phone" class="NokPhone" id="NokPhone" name="NokPhone" value="{{.NokPhone}}">
<input type="text" placeholder="Relationship to rider" class="NokRelation" id="NokRelation" name="NokRelation" value="{{.NokRelation}}">
</fieldset>
</div>
`

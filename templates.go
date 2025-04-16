package main

const tp_RiderDetails = `
<div class="RiderDetails">

<fieldset class="reqd">
<label for="RiderFirst">Rider's Name</label>
<input type="text" class="firstname" id="RiderFirst" name="RiderFirst" placeholder="First" value="{{.RiderFirst}}">
<input type="text" class="lastname" id="RiderLast" name="RiderLast" placeholder="Last" value="{{.RiderLast}}">
</fieldset>

<fieldset class="hide">
<label for="RiderIBA">Rider's IBA number</label>
<input type="text" class="ibanumber" id="RiderIBA" name="RiderIBA" value="{{.RiderIBA}}">
</fieldset>

<fieldset>
<label for="RiderRBL">Royal British Legion</label>
<select id="RiderRBL" name="RiderRBL">
<option value="" {{if eq .RiderRBL ""}}selected{{end}}>Please choose an option</option>
<option value="N" {{if eq .RiderRBL "N"}}selected{{end}}>No Legion association</option>
<option value="L" {{if eq .RiderRBL "L"}}selected{{end}}>I'm an ordinary Legion member</option>
<option value="R" {{if eq .RiderRBL "R"}}selected{{end}}>I am a Legion Rider (RBLR) member</option>
</select>
</fieldset>

<fieldset>
<label for="RiderNoviceYN">Is this your first {{.Event.NoviceEvent}}?</label>
<select id="RiderNoviceYN" name="RiderNoviceYN">
<option value="N" {{if eq .RiderNoviceYN "N"}}selected{{end}}>No, not my first time</option>
<option value="Y" {{if eq .RiderNoviceYN "Y"}}selected{{end}}>Yes, I'm a novice</option>
</select>
</fieldset>

<fieldset class="address">
<label for="RiderAddr1">Rider's address</label>
<input type="text" class="addr1" id="RiderAddr1" name="RiderAddr1" value="{{.RiderAddr1}}">
<input type="text" class="addr2" id="RiderAddr2" name="RiderAddr2" value="{{.RiderAddr2}}">
<input type="text" class="town" id="RiderTown" name="RiderTown" value="{{.RiderTown}}">
<input type="text" class="county" id="RiderCounty" name="RiderCounty" value="{{.RiderCounty}}">
<input type="text" class="postcode" id="RiderPostcode" name="RiderPostcode" value="{{.RiderPostcode}}">
<input type="text" class="country" id="RiderCountry" name="RiderCountry" value="{{.RiderCountry}}">
</fieldset>

<fieldset>
<label for="RiderPhone">Rider's mobile phone</label>
<input type="text" class="phone" id="RiderPhone" name="RiderPhone" value="{{.RiderPhone}}">
</fieldset>

<fieldset>
<label for="RiderEmail">Rider's email</label>
<input type="email" id="RiderEmail" name="RiderEmail" readonly value="{{.RiderEmail}}">
</fieldset>
</div>
`

const tp_PillionDetails = `
<div class="PillionDetails">

<fieldset class="reqd">
<label for="PillionFirst">Pillion's Name</label>
<input type="text" class="firstname" id="PillionFirst" name="PillionFirst" placeholder="First" value="{{.PillionFirst}}">
<input type="text" class="lastname" id="PillionLast" name="PillionLast" placeholder="Last" value="{{.PillionLast}}">
</fieldset>

<fieldset class="hide">
<label for="PillionIBA">Pillion's IBA number</label>
<input type="text" class="ibanumber" id="PillionIBA" name="PillionIBA" value="{{.PillionIBA}}">
</fieldset>

<fieldset>
<label for="PillionRBL">Royal British Legion</label>
<select id="PillionRBL" name="PillionRBL">
<option value="" {{if eq .PillionRBL ""}}selected{{end}}>Please choose an option</option>
<option value="N" {{if eq .PillionRBL "N"}}selected{{end}}>No Legion association</option>
<option value="L" {{if eq .PillionRBL "L"}}selected{{end}}>I'm an ordinary Legion member</option>
<option value="R" {{if eq .PillionRBL "R"}}selected{{end}}>I am a Legion Pillion (RBLR) member</option>
</select>
</fieldset>

<fieldset>
<label for="PillionNoviceYN">Is this your first {{.Event.NoviceEvent}}?</label>
<select id="PillionNoviceYN" name="PillionNoviceYN">
<option value="N" {{if eq .PillionNoviceYN "N"}}selected{{end}}>No, not my first time</option>
<option value="Y" {{if eq .PillionNoviceYN "Y"}}selected{{end}}>Yes, I'm a novice</option>
</select>
</fieldset>

<fieldset class="address">
<label for="PillionAddr1">Pillion's address</label>
<input type="text" class="addr1" id="PillionAddr1" name="PillionAddr1" value="{{.PillionAddr1}}">
<input type="text" class="addr2" id="PillionAddr2" name="PillionAddr2" value="{{.PillionAddr2}}">
<input type="text" class="town" id="PillionTown" name="PillionTown" value="{{.PillionTown}}">
<input type="text" class="county" id="PillionCounty" name="PillionCounty" value="{{.PillionCounty}}">
<input type="text" class="postcode" id="PillionPostcode" name="PillionPostcode" value="{{.PillionPostcode}}">
<input type="text" class="country" id="PillionCountry" name="PillionCountry" value="{{.PillionCountry}}">
</fieldset>

<fieldset>
<label for="PillionPhone">Pillion's mobile phone</label>
<input type="text" class="phone" id="PillionPhone" name="PillionPhone" value="{{.PillionPhone}}">
</fieldset>

<fieldset>
<label for="PillionEmail">Pillion's email</label>
<input type="email" id="PillionEmail" name="PillionEmail" readonly value="{{.PillionEmail}}">
</fieldset>
</div>
`

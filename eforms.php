<?php

$KONSTANTS['DBFILENAME'] = "eforms.db";

// Open the database	
try {
	$DB = new SQLite3($KONSTANTS['DBFILENAME']);
	$DB->exec("PRAGMA busy_timeout = 5000");  // Set lock timeout to 15 seconds
} catch(Exception $ex) {
	echo("OMG ".$ex->getMessage().' file=[ '.$KONSTANTS['DBFILENAME'].' ]');
}

function startPageHTML() {

    $nl = "\n";

    echo('<!DOCTYPE html>'.$nl);
    echo('<html lang="en">'.$nl);
    echo('<head>'.$nl);
    echo('<title>EForms</title>'.$nl);
    echo('<meta charset="UTF-8" />'.$nl);
    echo('<meta name="viewport" content="width=device-width, initial-scale=1" />'.$nl);
    echo('</head><body>'.$nl);

}

function finishHTML() {

    echo('</body></html>');
}

function showNewEntrant($event_code) {

    global $DB;

    startPageHTML();  
    echo('<div>');
    echo('<span class="vlabel"><label for="rider_first">Rider name: first / last</label> ');
    echo('<input type="text" id="rider_first" name="rider_first"> ');
    echo('<input type="text" id="rider_last" name="rider_last">');
    echo('</span>');
    echo('</div>');
}
?>

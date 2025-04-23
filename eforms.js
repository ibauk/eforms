"use strict";

const validbtn = "checktoken";
const vtc = "vtchar";
const vtl = vtc.length;

function clear_token() {
  let check = document.querySelector("#checktoken");
  if (check) check.disabled = false;
  let token = document.querySelector("#token");
  if (token) token.value = "";
  for (let i = 1; ; i++) {
    let x = vtc + i;
    let y = document.getElementById(x);
    if (!y) break;
    y.value = "";
    if (i == 1) y.focus();
  }
}

function flipPillionDetails(obj) {
  console.log("flipPillionDetails " + obj.value);
  let pd = document.getElementById("divPillionDetails");
  if (!pd) return;
  if (obj.value == "Y") {
    pd.classList.remove("hide");
    console.log('Disabling inputs')
    let pemail = document.querySelector("#PillionEmail")
    if (!pemail) return
    console.log("or enabling them")
    let ok = (pemail.value !="" && !pemail.classList.contains("oi"))
    console.log("ok is "+ok)
    
    var pdl = Array.prototype.slice.call(pd.querySelectorAll("*"))

    pdl.forEach(elem => {
      if (elem.tagName=="INPUT" || elem.tagName=="SELECT") {
        if (elem.name != "PillionEmail") elem.disabled=!ok
      }
    });
  } else {
    pd.classList.add("hide");
  }
}

function showConfirmID(ridername, entrant) {
  console.log(
    'showConfirmID called with "' + ridername + '" and [' + entrant + "]"
  );
  let cid = document.getElementById("confirmID");
  if (!cid) return;
  let lbl = document.querySelector("#confirmID label");
  if (!lbl) return;
  let rn = ridername.trim();
  if (rn == "") {
    lbl.innerHTML = "You're a new entrant?";
  } else {
    lbl.innerHTML = "<strong>" + ridername + "</strong> - is this you?";
  }
  cid.setAttribute("data-entrant", entrant);
  cid.setAttribute("data-ridername", ridername);
  cid.classList.remove("hide");
  let inp = document.querySelector("#confirmID input");
  if (inp) inp.focus();
}

function show_form_start() {
  console.log("show_form_start called");
  let cid = document.getElementById("confirmID");
  if (!cid) return;
  let email = document.getElementById("email");
  if (!email) return;
  let rally = document.getElementById("rally");
  let ridername = cid.getAttribute("data-ridername");
  let entrant = cid.getAttribute("data-entrant");

  console.log("sfs calling now");
  let url = "/f?email=" + encodeURIComponent(email.value);
  url += "&rally=" + encodeURIComponent(rally.value);
  url += "&rn=" + encodeURIComponent(ridername);
  url += "&er=" + encodeURIComponent(entrant);
  window.location.href = url;
}

function show_signup_start() {
  window.location.href = "/s";
}
function retry_email(obj) {
  let tevbtn = document.getElementById("tevbtn");
  if (tevbtn) tevbtn.disabled = false;
  const tz = document.getElementsByClassName("tokenzone");
  for (let i = 0; i < tz.length; i++) tz[i].classList.add("hide");
}
function tokenInput(inp) {
  let vtmax = document.getElementById("tokenlen").value;

  let c = "";
  let id = inp.getAttribute("id").substring(vtl);
  for (; id <= vtmax; id++) {
    let x = document.getElementById(vtc + id);
    if (!x) break;
    x.value += c;
    c = "";
    if (x.value.length > 1) {
      c = x.value.substring(1);
      x.value = x.value.substring(0, 1);
    } else break;
  }
  if (id < vtmax) {
    let x = document.getElementById(vtc + (id + 1));
    if (x) x.focus();
  } else {
    let x = document.getElementById(validbtn);
    if (x) x.focus();
  }
}

function trigger_email_validation(obj) {
  const checkFailed = "&#9746;";
  const checkOK = "&#9745;";
  console.log("trigger_email_validtion");
  let email = document.querySelector("#email").value;
  if (email == "") return;
  let rally = document.querySelector("#rally").value;
  if (rally == "") return;

  console.log("tev still here");
  if (obj) obj.disabled = true;

  let url = "/x?email=" + encodeURIComponent(email);
  url += "&rally=" + encodeURIComponent(rally);
  let token = document.querySelector("#token");
  if (token && token.value != "")
    url += "&token=" + encodeURIComponent(token.value);
  let res = document.getElementById("checkresult");
  fetch(url)
    .then((response) => {
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
      return response.json();
    })
    .then((data) => {
      if (!data.ok) {
        console.error(`Validation failed for ${data.msg}`);
        if (res) res.innerHTML = checkFailed;
        clear_token();
      } else {
        console.log(data);
        const tz = document.getElementsByClassName("tokenzone");
        for (let i = 0; i < tz.length; i++) tz[i].classList.remove("hide");

        if (res && data.msg != "") {
          res.innerHTML = checkOK;
          if (data.msg != "ok") {
            showConfirmID(data.msg, data.entrant);
          } else {
            showConfirmID("", 0);
          }
        }
        if (data.msg == "") {
          let x = document.getElementById(vtc + "1");
          if (x) x.focus();
        }
      }
    })
    .catch((error) => {
      console.error("Fetch error:", error);
    });
}

function verify_email_validation(obj) {
  console.log("verify_email_validation called");
  if (obj) obj.disabled = true;
  let tkn = document.getElementById("token");
  if (!tkn) return;
  tkn.value = "";
  for (let id = 1; ; id++) {
    let x = document.getElementById(vtc + id);
    if (!x) break;
    tkn.value += x.value;
  }

  console.log("vev calling now");
  trigger_email_validation();
}

// Background transmission routines

const myStackItem = "eformsStack";

function oc(obj) {
  saveData(obj);
}

function oi(obj) {
  obj.classList.remove("oc");
  obj.classList.add("oi");
  // Checkbox handler
  obj.setAttribute("data-chg", "1");
  // autosave handler
  if (obj.timer) {
    clearTimeout(obj.timer);
  }
  obj.timer = setTimeout(saveData, 3000, obj);
}

function saveData(obj) {
  if (obj.getAttribute("data-static") == "") obj.setAttribute("data-chg", "");
  console.log("saveData: " + obj.name);
  if (obj.timer) {
    clearTimeout(obj.timer);
  }

  let ent = document.getElementById("EntrantNumber").value;
  let rid = document.getElementById("RiderID").value;
  let pil = document.getElementById("PillionID").value;
  let val = obj.value;
  switch (obj.name) {
    case "RiderPostcode":
    case "PillionPostcode":
    case "BikeReg":
    case "RiderCountry":
    case "PillionCountry":
      val = val.toUpperCase();
      break;
  }

  let url = encodeURI("/z?e=" + ent + "&r=" + rid + "&p=" + pil + "&f=" + obj.name + "&v=" + val);
  stackTransaction(url, obj.id);
  sendTransactions()
  
}

function sendTransactions() {
  let stackx = sessionStorage.getItem(myStackItem);
  if (stackx == null) return;

  let stack = JSON.parse(stackx);

  //console.log(stack);

  let errlog = document.getElementById("errlog");

  while (stack.length > 0) {
    let itm = stack[0];
    stack.splice(0, 1);
    sessionStorage.setItem(myStackItem, JSON.stringify(stack));
    console.log("Sending: " + itm.url);

    fetch(itm.url)
      .then((response) => {
        if (!response.ok) {
          // Handle HTTP errors
          stackTransaction(itm.url, itm.objid);
          //if (errlog){errlog.innerHTML=`HTTP error! Status: ${response.status}`}

          throw new Error(`HTTP error! Status: ${response.status}`);
        }
        return response.json();
      })
      .then((data) => {
        if (data.err) {
          // Handle JSON error field
          console.error(`Error: ${data.msg}`);
        } else {
          // Process the data if no error
          //if (errlog){errlog.innerHTML="Hello sailor: "+JSON.stringify(data)}
          console.log("Data:", data);
          document.getElementById(itm.objid).classList.replace("oi", "oc");
        }
      })
      .catch((error) => {
        // Handle network or other errors
        //if (errlog) {errlog.innerHTML="ERROR CAUGHT"}
        stackTransaction(itm.url, itm.objid);
        console.error("Fetch error:", error);
        return;
      });
  }
}

function stackTransaction(url, objid) {
  console.log(url);
  let newTrans = {};
  newTrans.url = url;
  newTrans.objid = objid;
  newTrans.sent = false;

  const stackx = sessionStorage.getItem(myStackItem);
  let stack = [];
  if (stackx != null) {
    stack = JSON.parse(stackx);
  }
  stack.push(newTrans);
  sessionStorage.setItem(myStackItem, JSON.stringify(stack));
}

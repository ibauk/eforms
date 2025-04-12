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
  let email = document.querySelector("#email").value;
  if (email == "") return;
  let rally = document.querySelector("#rally").value;
  if (rally == "") return;

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

        if (res && data.msg != "") res.innerHTML = checkOK;
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
  if (obj) obj.disabled = true;
  let tkn = document.getElementById("token");
  if (!tkn) return;
  tkn.value = "";
  for (let id = 1; ; id++) {
    let x = document.getElementById(vtc + id);
    if (!x) break;
    tkn.value += x.value;
  }

  trigger_email_validation();
}

"use strict";

function tokenInput(inp) {
  const vtc = "vtchar";
  const vtl = vtc.length;
  let vtmax = document.getElementById("tokenlen").value;

  let id = inp.getAttribute("id").substr(vtl);
  console.log("Checking " + id);
  if (inp.value.length > 1) {
    let c = inp.value.substr(1);
    inp.value = inp.value.substr(0, 1);
    id++;
  }
  id++;
  if (id <= vtmax) {
    console.log("Focusing on " + vtc + id);
    document.getElementById(vtc + id).focus();
  }
}
function trigger_email_validation(obj) {
  let email = document.querySelector("#email").value;
  if (email == "") return;
  let url = "/x?email=" + encodeURIComponent(email);
  fetch(url)
    .then((response) => {
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
      return response.json();
    })
    .then((data) => {
      if (!data.ok) {
        console.error(`Error! ${data.msg}`);
      } else {
        console.log(data);
      }
    })
    .catch((error) => {
      console.error("Fetch error:", error);
    });
}

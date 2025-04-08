"use strict";

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

/*
 * client.js
 */

fetch("/api/a $adjective $noun")
  .then((response) => response.json())
  .then((data) => {
    // console.log(data);
    document.getElementById("example").innerHTML = `
          <div>
              <h3> ${data.phrase} </h3>
          </div>`;
  });

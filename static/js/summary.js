var iframe = document.getElementById("markdown")

window.onclick = function (e) {
    if (e.target.localName == 'a') {
        e.preventDefault();
        var md_addr = e.target.getAttribute("href");
        iframe.src = md_addr;
    }
}

var searchbar = document.querySelector("#searchbar");
searchbar.addEventListener("keyup", function(event) {
  if (event.keyCode === 13) {
   event.preventDefault();
   var md_addr = "/search?search_string=" + searchbar.value;
   iframe.src = md_addr;
   searchbar.value = "";
  }
});


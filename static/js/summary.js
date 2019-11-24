var iframe = document.getElementById("markdown")

function selectChapter(chapterELement){
  // unselect others
  const elements = [...document.querySelectorAll("a")];
  elements.forEach(element=>{
    element.parentElement.classList.remove("selected")
  })
  
  // select new one
  chapterELement.parentElement.classList.add("selected")
}

window.onclick = function (e) {
    if (e.target.localName == 'a') {
        e.preventDefault();
        var md_addr = e.target.getAttribute("href");
        iframe.src = md_addr;
        // select new one
        this.selectChapter(e.target)
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



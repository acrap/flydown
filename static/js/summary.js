var iframe = document.getElementById("markdown")

function unselectAllChapters(){
  const elements = [...document.querySelectorAll("a")];
  elements.forEach(element=>{
    element.parentElement.classList.remove("selected")
  })
  
}

function selectChapter(chapterELement){
  // unselect others
  unselectAllChapters()
  
  // select new one
  chapterELement.parentElement.classList.add("selected")
}

function selectChapterByURL( url ){
  const chapters = [...document.querySelectorAll("div.summary a")];
  for (index in chapters){
      href = chapters[index].href;
      if(href.localeCompare(url)==0){
        selectChapter(chapters[index]);
        return;
      }
    }
  // nothing found
  unselectAllChapters()
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



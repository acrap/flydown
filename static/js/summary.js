function unselectAllChapters(){
  const elements = [...document.querySelectorAll("a")];
  elements.forEach(element=>{
    element.classList.remove("selected")
  })
  
}

function selectChapter(chapterElement){
  // unselect others
  unselectAllChapters()
  
  // select new one
  chapterElement.classList.add("selected")
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

function doMdRequest(md_addr) {
  var xhr = new XMLHttpRequest();
  var allParams = getAllUrlParams(md_addr);
  var search_string = decodeURI(allParams.search_string);
  var entry_number = allParams.n;

  xhr.open('GET', md_addr, false);
  xhr.send();
  if (xhr.status != 200) {
    alert(xhr.status + ': ' + xhr.statusText);
  } else {
    var md_viewer = document.querySelector("#markdown-viewer");
    md_viewer.innerHTML = ""
    md_viewer.innerHTML = xhr.responseText;
    Prism.highlightAll();
    // store the last page for 1 day, but not search pages
    if (!md_addr.includes("search?search_string")){
      if(md_addr.includes("search_string")){
        COOKIE.set("last-page",md_addr.substring(0,md_addr.indexOf('?')));
      }else{
        COOKIE.set("last-page",md_addr, 1);
      }
      
    }
  }
  
  if (entry_number == undefined) {
    if (search_string != undefined) {
      var nodes = get_nodes_containing_text("*", search_string);
      nodes.forEach(element => {
        highlightText(element, search_string)
      })
    }
  } else {
    // we got search string and entry number, so we should highlight and scroll
    if (search_string != undefined) {
      scrollAndHighlight(search_string, entry_number)
    }
  }
  
}

// catch clicks on links to handle them properly 
window.onclick = function (e) {
    if (e.target.localName == 'a') {
        e.preventDefault();
        var link_addr = e.target.getAttribute("href");
        if (e.target.host == window.location.host){
          doMdRequest(link_addr)
          // select new one
          selectChapter(e.target)
        }
        else{
          window.open(link_addr,'_blank');
        }
    }
}

var search_bar = document.querySelector("#search-bar");

search_bar.addEventListener("keyup", function(event) {
  if (event.keyCode === 13) {
   event.preventDefault();
   var md_addr = "/search?search_string=" + search_bar.value;
   doMdRequest(md_addr)
   search_bar.value = "";
  }
});





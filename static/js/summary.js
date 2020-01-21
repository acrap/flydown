let buildInStylesPath = "/static/css/"

function unselectAllChapters() {
  const elements = [...document.querySelectorAll("a")];
  elements.forEach(element => {
    element.classList.remove("selected")
  })

}

function selectChapter(chapterElement) {
  // unselect others
  unselectAllChapters()

  // select new one
  chapterElement.classList.add("selected")
}

function selectChapterByURL(url) {
  const chapters = [...document.querySelectorAll("div.summary a")];
  for (index in chapters) {
    href = chapters[index].href;
    if (href.localeCompare(url) == 0) {
      selectChapter(chapters[index]);
      return;
    }
  }
  // nothing found
  unselectAllChapters()
}

function calcChapterUrl(md_url){
  md_url = md_url.replace("/","%2F")
  return "/"+md_url.replace(".md","")
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

    window.history.pushState("",md_addr,calcChapterUrl(md_addr));
  }

  if (entry_number == undefined) {
    if (search_string != undefined) {
      var nodes = get_nodes_containing_text("*", search_string);
      nodes.forEach(element => {
        highlightText(element, search_string)
      })
    }
    // scroll to the top
    var book_body = document.querySelector(".book-body");
    book_body.scrollTop = 0;

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
    if (e.target.host == window.location.host) {
      // select new one
      selectChapter(e.target);

      doMdRequest(link_addr)
    }
    else {
      window.open(link_addr, '_blank');
    }
  }
}

var search_bar = document.querySelector("#search-bar");

search_bar.addEventListener("keyup", function (event) {
  if (event.keyCode === 13) {
    event.preventDefault();
    var md_addr = "/search?search_string=" + search_bar.value;
    doMdRequest(md_addr)
    search_bar.value = "";
  }
});

function switchMode(event) {
  event = event || window.event;
  event.preventDefault();
  theme = COOKIE.get("theme");
  mainStyle = document.getElementById("theme-css")
  prismStyle = document.getElementById("prism-theme-css")
  if ((theme == null) ||
    (theme.localeCompare("light") == 0)) {
    COOKIE.set("theme", "dark");
    mainStyle.href = buildInStylesPath + "dark.css"
    prismStyle.href = buildInStylesPath + "prism-dark.css"
  } else {
    COOKIE.set("theme", "light");
    mainStyle.href = buildInStylesPath + "light.css"
    prismStyle.href = buildInStylesPath + "prism-light.css"
  }
  Prism.highlightAll();
}

var hideButton = document.getElementById("hide-button")
book_body = document.querySelector(".book-body");
summary = document.querySelector(".summary");

function toggleSummary(event){
  event = event || window.event;
  event.preventDefault();
  if (summary.hidden) {
    summary.hidden = false;
    book_body.style.left = "17vw";
    book_body.style.width = "83vw";
  } else {
    summary.hidden = true;
    book_body.style.left = "0vw";
    book_body.style.width = "100vw";
  }
}

window.addEventListener('popstate', function(event) {
  this.location.reload()
}, false);

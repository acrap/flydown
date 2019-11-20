var iframe = document.getElementById("markdown")

window.onclick = function (e) {
    if (e.target.localName == 'a') {
        e.preventDefault()
        var md_addr = e.target.getAttribute("href")
        iframe.src = md_addr
    }
}

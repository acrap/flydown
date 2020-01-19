package flydown

import (
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"
)

const defaultColorTheme = "light"

type preferences struct {
	lastPage string
	theme    string
}

// MdGenerator is a main markdown generator structure
var MdGenerator Generator

func getUserPreferences(r *http.Request) (c preferences) {
	userCookie, err := r.Cookie("last-page")
	if err != nil {
		c.lastPage = MdGenerator.readmeName
	} else {
		c.lastPage = userCookie.Value
	}

	userCookie, err = r.Cookie("theme")
	if err != nil {
		c.theme = defaultColorTheme
	} else {
		c.theme = userCookie.Value
	}
	return
}

// RenderMdHandleFunc handler to host markdown as html page on the fly
func RenderMdHandleFunc(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.URL.Path, ".md") {
		http.Error(w, "Unsupported file format", 404)
	}
	md, err := MdGenerator.NewMarkdownFromFile(r.URL.Path[1:])
	if err != nil {
		http.Error(w, "Cannot parse md file", 404)
		return
	}
	w.Write([]byte(md.htmlRepresentation))

}

// MainHandleFunc handler to host the main page with summary on the left side and markdown on the right
func MainHandleFunc(w http.ResponseWriter, r *http.Request) {
	var mdPage *Markdown
	var err error
	if strings.Contains(r.URL.Path, ".md") {
		RenderMdHandleFunc(w, r)
		return
	}

	md, err := MdGenerator.NewMarkdownFromFile(MdGenerator.summaryName)
	if err != nil {
		http.Error(w, "Cannot parse md file", 404)
		return
	}
	templateBytes, err := ioutil.ReadFile("./templates/index.html")
	if err != nil {
		http.Error(w, "No such template to render", 500)
		return
	}

	preferences := getUserPreferences(r)
	var pageAddr string
	if len(r.URL.Path) > 1 {
		pageAddr = strings.ReplaceAll(r.URL.Path[1:], "%2F", "/") + ".md"
	} else {
		pageAddr = preferences.lastPage
	}

	mdPage, err = MdGenerator.NewMarkdownFromFile(pageAddr)
	if err != nil {
		http.Error(w, "Cannot load the page", 500)
		return
	}

	data := struct {
		Theme    string
		Summary  string
		Favicon  string
		BookName string
		Markdown string
	}{
		Theme:    preferences.theme,
		Summary:  md.htmlRepresentation,
		Favicon:  MdGenerator.favicon,
		BookName: MdGenerator.bookName,
		Markdown: mdPage.htmlRepresentation,
	}
	template, err := template.New("index").Parse(string(templateBytes))
	err = template.Execute(w, data)
	if err != nil {
		http.Error(w, "Cannot get summary", 500)
	}
}

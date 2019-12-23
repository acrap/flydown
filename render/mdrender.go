package flydown

import (
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"
)

// MdGenerator is a main markdown generator structure
var MdGenerator Generator

// RenderMdHandleFunc handler to host markdown as html page on the fly
func RenderMdHandleFunc(w http.ResponseWriter, r *http.Request) {
	var templateBytes []byte
	if !strings.Contains(r.URL.Path, ".md") {
		http.Error(w, "Unsupported file format", 404)
	}
	r.URL.Path = strings.Replace(r.URL.Path, "md/", "", 1)
	md, err := MdGenerator.NewMarkdownFromFile(r.URL.Path[1:])
	if err != nil {
		http.Error(w, "Cannot parse md file", 404)
		return
	}

	templateBytes, err = ioutil.ReadFile("./templates/md.html")
	if err != nil {
		http.Error(w, "No such template to render", 500)
		return
	}
	tmpl, err := template.New("markdown").Parse(string(templateBytes))
	data := struct {
		Markdown string
	}{
		Markdown: md.htmlRepresentation,
	}
	tmpl.Execute(w, data)
}

// MainHandleFunc handler to host the main page with summary on the left side and markdown on the right
func MainHandleFunc(w http.ResponseWriter, r *http.Request) {
	md, err := MdGenerator.NewMarkdownFromFile("summary.md")
	if err != nil {
		http.Error(w, "Cannot parse md file", 404)
		return
	}
	templateBytes, err := ioutil.ReadFile("./templates/index.html")
	if err != nil {
		http.Error(w, "No such template to render", 500)
		return
	}

	data := struct {
		Summary string
		Favicon string
	}{
		Summary: md.htmlRepresentation,
		Favicon: MdGenerator.favicon,
	}
	tmpl, err := template.New("index").Parse(string(templateBytes))
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Cannot get summary", 500)
	}
}

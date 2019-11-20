package flydown

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/gomarkdown/markdown"
	mdhtml "github.com/gomarkdown/markdown/html"
)

var folderToHost string

// SetFolderToHost used to set the directory with markdown which should be hosted
func SetFolderToHost(folder string) error {
	if _, err := os.Stat(folder + string(os.PathSeparator) + "summary.md"); os.IsNotExist(err) {
		return err
	}
	folderToHost = folder
	return nil
}

// RenderMdHandleFunc handler to host markdown as html page on the fly
func RenderMdHandleFunc(w http.ResponseWriter, r *http.Request) {
	var templateBytes []byte
	opts := mdhtml.RendererOptions{
		Flags:          mdhtml.CommonFlags,
		RenderNodeHook: nil,
	}
	renderer := mdhtml.NewRenderer(opts)
	if !strings.Contains(r.URL.Path, ".md") {
		http.Error(w, "Unsupported file format", 404)
	}
	r.URL.Path = strings.Replace(r.URL.Path, "md/", "", 1)
	mdfilepath := folderToHost + string(os.PathSeparator) + r.URL.Path[1:]

	md, err := ioutil.ReadFile(mdfilepath)
	if err != nil {
		http.Error(w, "No such file", 404)
		return
	}
	mdHTMLBytes := markdown.ToHTML([]byte(md), nil, renderer)
	templateBytes, err = ioutil.ReadFile("./templates/md.html")
	if err != nil {
		http.Error(w, "No such template to render", 500)
		return
	}
	tmpl, err := template.New("markdown").Parse(string(templateBytes))
	data := struct {
		Markdown string
	}{
		Markdown: string(mdHTMLBytes),
	}
	tmpl.Execute(w, data)
}

// MainHandleFunc handler to host the main page with summary on the left side and markdown on the right
func MainHandleFunc(w http.ResponseWriter, r *http.Request) {
	opts := mdhtml.RendererOptions{
		Flags:          mdhtml.CommonFlags,
		RenderNodeHook: nil,
	}
	renderer := mdhtml.NewRenderer(opts)
	pathToSummary := folderToHost + string(os.PathSeparator) + "summary.md"
	md, err := ioutil.ReadFile(pathToSummary)

	mdBytes := markdown.ToHTML(md, nil, renderer)

	if err != nil {
		http.Error(w, "No such file", 404)
		return
	}
	templateBytes, err := ioutil.ReadFile("./templates/index.html")
	if err != nil {
		http.Error(w, "No such template to render", 500)
		return
	}
	data := struct {
		Summary string
	}{
		Summary: string(mdBytes),
	}
	tmpl, err := template.New("index").Parse(string(templateBytes))
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Cannot get summary", 500)
	}
}

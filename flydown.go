package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/NYTimes/gziphandler"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	mdhtml "github.com/gomarkdown/markdown/html"
)

// return (ast.GoToNext, true) to tell html renderer to skip rendering this node
// (because you've rendered it)
func renderHookDropCodeBlock(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	// skip all nodes that are not CodeBlock nodes
	if _, ok := node.(*ast.CodeBlock); !ok {
		return ast.GoToNext, false
	}
	// custom rendering logic for ast.CodeBlock. By doing nothing it won't be
	// present in the output
	return ast.GoToNext, true
}

var folderToHost string

func mainHandleFunc(w http.ResponseWriter, r *http.Request) {
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

func renderMdHandleFunc(w http.ResponseWriter, r *http.Request) {
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

func main() {
	var currentDir string
	var err error
	currentDir, err = os.Getwd()
	if err != nil {
		log.Fatal("please pass the folder to share")
	}

	flag.StringVar(&folderToHost, "share_folder", currentDir, "Path to the directory you want to share")
	flag.Parse()
	http.Handle("/static/", gziphandler.GzipHandler(
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./static")))))
	http.HandleFunc("/md/", renderMdHandleFunc)
	http.HandleFunc("/", mainHandleFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/NYTimes/gziphandler"
	"github.com/acrap/flydown/flydown"
	"github.com/gomarkdown/markdown/ast"
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

func main() {
	var folderToHost string
	var currentDir string
	var err error
	currentDir, err = os.Getwd()
	if err != nil {
		log.Fatal("please pass the folder to share")
	}

	flag.StringVar(&folderToHost, "share_folder", currentDir, "Path to the directory you want to share")
	flag.Parse()
	if err = flydown.SetFolderToHost(folderToHost); err != nil {
		log.Fatal(err)
	}
	http.Handle("/static/", gziphandler.GzipHandler(
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./static")))))
	http.HandleFunc("/md/", flydown.RenderMdHandleFunc)
	http.HandleFunc("/", flydown.MainHandleFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/NYTimes/gziphandler"
	flydown "github.com/acrap/flydown/render"
)

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
	http.HandleFunc("/search", flydown.SearchHandleFunc)
	http.HandleFunc("/", flydown.MainHandleFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

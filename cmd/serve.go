/*
Copyright © 2020 Andrey Strunin acrapmonster@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/NYTimes/gziphandler"
	flydown "github.com/acrap/flydown/render"
	"github.com/spf13/cobra"
)

const defPort = 8080
const defBookName = "My book"

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve markdown folder",
	Run: func(cmd *cobra.Command, args []string) {
		var folderToHost string
		var err error

		if err != nil {
			log.Fatal("please pass the folder to share")
		}

		folderToHost, _ = cmd.Flags().GetString("shareFolder")
		port, _ := cmd.Flags().GetInt("port")
		portStr := ":" + strconv.Itoa(port)
		bookName, _ := cmd.Flags().GetString("bookName")
		ipStr, _ := cmd.Flags().GetString("ip")
		flag.Parse()

		if err = flydown.MdGenerator.Init(folderToHost, bookName); err != nil {
			log.Fatal(err)
		}
		http.Handle("/static/", gziphandler.GzipHandler(
			http.StripPrefix("/static/",
				http.FileServer(http.Dir("./static")))))
		http.Handle("/public/", gziphandler.GzipHandler(
			http.StripPrefix("/public/",
				http.FileServer(http.Dir(folderToHost+"/public")))))

		http.HandleFunc("/search", flydown.SearchHandleFunc)
		http.HandleFunc("/", flydown.MainHandleFunc)
		fmt.Fprintf(os.Stdout, "Served on http://%s%s\n", ipStr, portStr)
		log.Fatal(http.ListenAndServe(ipStr+portStr, nil))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	currentDir, _ := os.Getwd()
	serveCmd.Flags().StringP("shareFolder", "s", currentDir+string(os.PathSeparator)+"doc", "Pass the path to directory with markdown content")
	serveCmd.Flags().StringP("ip", "i", "127.0.0.1", "Pass the IP addr")
	serveCmd.Flags().IntP("port", "p", defPort, "Pass the port")
	serveCmd.Flags().StringP("bookName", "n", defBookName, "Pass the name of your book")
}

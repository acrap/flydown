package flydown

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gomarkdown/markdown"
	mdhtml "github.com/gomarkdown/markdown/html"
)

const contextLines int = 4

func searchStringInFile(filename string, searchStr string) (context []string, lines []int, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	contextRing := make([]string, contextLines)
	cRingCounter := 0

	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)

	line := 1
	// https://golang.org/pkg/bufio/#Scanner.Scan
	for scanner.Scan() {
		currentLineStr := scanner.Text()
		isResultFound := strings.Contains(currentLineStr, searchStr)
		if isResultFound {
			isResultInsideLink := false
			linkEndIndex := 0
			linkStartIndex := strings.Index(currentLineStr, string("]("))
			if linkStartIndex > 0 {
				linkEndIndex = strings.Index(currentLineStr, string(")"))
				if strings.Index(currentLineStr, searchStr) > linkStartIndex && strings.Index(currentLineStr, searchStr) < linkEndIndex {
					isResultInsideLink = true
					isResultFound = false
					fmt.Println(currentLineStr)
				}
			}
			if !isResultInsideLink {
				currentLineStr = strings.Replace(currentLineStr, searchStr, "**"+searchStr+"**", -1)
			}

		}
		if cRingCounter < contextLines {
			contextRing = append(contextRing, currentLineStr)
			cRingCounter++
		} else {
			contextRing = append(contextRing[1:], currentLineStr)
		}
		if isResultFound {
			lines = append(lines, line)
			curContext := ""
			for _, l := range contextRing {
				if len(l) > 0 {
					curContext += l
					curContext += "\n"
				}
			}
			context = append(context, curContext)
			// reset ring buffer
			cRingCounter = 0
			contextRing = nil

		}
		line++
	}

	if err := scanner.Err(); err != nil {
		// Handle the error
		log.Printf("error in search handler %v\n", err)
	}
	return context, lines, err
}

// SearchHandleFunc handler for searching request
func SearchHandleFunc(w http.ResponseWriter, r *http.Request) {
	var searchStr string
	type result struct {
		filename string
		lines    []int
		context  []string
	}
	opts := mdhtml.RendererOptions{
		Flags:          mdhtml.CommonFlags,
		RenderNodeHook: nil,
	}
	renderer := mdhtml.NewRenderer(opts)

	var results []result
	resultMd := ""
	r.ParseForm()

	searchStr = r.URL.Query()["search_string"][0]
	err := filepath.Walk(folderToHost, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".md") {

			context, lines, err := searchStringInFile(path, searchStr)
			if err == nil && lines != nil {
				curResult := result{filename: path, lines: nil}
				curResult.lines = lines
				curResult.context = context
				results = append(results, curResult)
			}

		}
		return nil
	})
	for _, r := range results {
		for i, c := range r.context {
			fileAndLineLink := fmt.Sprintf("[%s:%d](%s)\n\n", r.filename, r.lines[i], r.filename)
			md := fileAndLineLink + c + "\n" + "\n"
			mdHTMLBytes := markdown.ToHTML([]byte(md), nil, renderer)
			resultMd += string(mdHTMLBytes)
		}
	}

	if err != nil {
		http.Error(w, "Error in search handler", 404)
	}

	templateBytes, err := ioutil.ReadFile("./templates/search.html")
	if err != nil {
		http.Error(w, "No such template to render", 500)
		return
	}
	data := struct {
		SearchResults string
	}{
		SearchResults: resultMd,
	}
	tmpl, err := template.New("search").Parse(string(templateBytes))
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Cannot get summary", 500)
	}
}

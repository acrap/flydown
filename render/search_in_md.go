package flydown

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"text/template"
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
		currentLineStrLow := strings.ToLower(currentLineStr)
		isResultFound := strings.Contains(currentLineStrLow, searchStr)
		if isResultFound {
			linkEndIndex := 0
			linkStartIndex := strings.Index(currentLineStr, string("]("))
			indexFoundResult := strings.Index(currentLineStrLow, searchStr)
			if linkStartIndex > 0 {
				linkEndIndex = strings.Index(currentLineStr, string(")"))
				if (indexFoundResult > linkStartIndex) && (indexFoundResult < linkEndIndex) {
					isResultFound = false
				}
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

type result struct {
	filename string
	lines    []int
	context  []string
}

func searchInMdFiles(searchStr string) chan result {
	var wg sync.WaitGroup
	resultsChan := make(chan result, 2)
	filepath.Walk(mdGenerator.rootMdFolder, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".md") {
			wg.Add(1)
			go func(path string, searchStr string, resultsChan chan<- result, wg *sync.WaitGroup) {
				context, lines, err := searchStringInFile(path, searchStr)
				if err == nil && lines != nil {
					curResult := result{filename: path, lines: nil}
					curResult.lines = lines
					curResult.context = context
					resultsChan <- curResult
				}
				wg.Done()
			}(path, searchStr, resultsChan, &wg)
		}
		return nil
	})

	go func() {
		wg.Wait()
		close(resultsChan)
	}()
	return resultsChan
}

// SearchHandleFunc handler for searching request
func SearchHandleFunc(w http.ResponseWriter, r *http.Request) {
	var err error

	resultMd := ""
	r.ParseForm()
	searchStr := r.URL.Query().Get("search_string")
	if searchStr == "" {
		http.Error(w, "Empty search request", 400)
		return
	}
	searchStr = strings.ToLower(searchStr)
	resultsChan := searchInMdFiles(searchStr)
	entryNum := 0

	waitForTemplateChannel := make(chan []byte)
	go func(channel chan []byte) {
		templateBytes, _ := ioutil.ReadFile("./templates/search.html")
		channel <- templateBytes
	}(waitForTemplateChannel)

	for {
		res, ok := <-resultsChan
		if ok == false {
			break
		}
		for i, c := range res.context {
			additionalParams := fmt.Sprintf("?%s=%s&%s=%s&", "search_string", searchStr, "n", strconv.Itoa(entryNum))
			fixedLink := strings.ReplaceAll(res.filename, mdGenerator.rootMdFolder, "md")
			fileAndLineLink := fmt.Sprintf("[%s:%d](%s)\n\n", res.filename, res.lines[i], fixedLink+additionalParams)
			md := fileAndLineLink + c + "\n" + "\n"
			resultMd += ConvertMdStrToHTML(md)
			entryNum++
		}
	}

	if resultMd == "" {
		resultMd += ConvertMdStrToHTML("# Results not found")
	}
	data := struct {
		SearchResults string
	}{
		SearchResults: resultMd,
	}
	templateBytes := <-waitForTemplateChannel
	if templateBytes == nil {
		http.Error(w, "No such template to render", 500)
		return
	}

	tmpl, err := template.New("search").Parse(string(templateBytes))
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Cannot get summary", 500)
	}
}

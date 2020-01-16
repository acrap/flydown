package flydown

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gomarkdown/markdown"
	mdhtml "github.com/gomarkdown/markdown/html"
)

// Markdown is a base structure
type Markdown struct {
	htmlRepresentation     string
	markdownRepresentation string
}

// Generator is a struct that keeps common settings for markdown generator
type Generator struct {
	rootMdFolder string
	favicon      string
	bookName     string
	summaryName  string
	readmeName   string
}

// NewMarkdown create the new Markdown structure
func NewMarkdown(mdByteArray []byte) (result Markdown) {
	opts := mdhtml.RendererOptions{
		Flags:          mdhtml.CommonFlags,
		RenderNodeHook: nil,
	}

	renderer := mdhtml.NewRenderer(opts)
	htmlBytes := markdown.ToHTML(mdByteArray, nil, renderer)
	result.htmlRepresentation = string(htmlBytes)
	result.markdownRepresentation = string(mdByteArray)
	return
}

func searchForFilename(nameList []string, folder string) (result string, err error) {
	for _, fileName := range nameList {
		if _, err := os.Stat(folder + string(os.PathSeparator) + fileName); !os.IsNotExist(err) {
			return fileName, nil
		}
	}
	return "", os.ErrNotExist
}

// FindSummary finds and stores summary filename
func (generator *Generator) FindSummary() error {
	summaryName, err := searchForFilename([]string{"summary.md", "SUMMARY.md", "Summary.md"}, generator.rootMdFolder)
	if err != nil {
		return err
	}
	generator.summaryName = summaryName
	return nil
}

// FindReadme finds and stores readme filename
func (generator *Generator) FindReadme() error {
	readmeName, err := searchForFilename([]string{"readme.md", "README.md", "Readme.md"}, generator.rootMdFolder)
	if err != nil {
		return err
	}
	generator.readmeName = readmeName
	return nil
}

// Init initialize new generator
func (generator *Generator) Init(rootMdFolder string, bookName string) error {
	generator.rootMdFolder = rootMdFolder
	if err := generator.FindSummary(); err != nil {
		return fmt.Errorf("Summary file is not found")
	}
	if err := generator.FindReadme(); err != nil {
		return fmt.Errorf("Readme file is not found")
	}
	generator.favicon = generator.getFavicon()
	generator.bookName = bookName
	return nil
}

// GetFavicon returns custom favicon if it does exist, otherwise returns a default one
func (generator *Generator) getFavicon() string {
	fullPath := string(os.PathSeparator) + "public" + string(os.PathSeparator) + "favicon.png"
	_, err := os.Stat(generator.rootMdFolder + fullPath)
	if os.IsNotExist(err) {
		return "/static/images/favicon.png"
	}
	return fullPath
}

// NewMarkdownFromFile create Markdown struct from file
func (generator *Generator) NewMarkdownFromFile(filename string) (result *Markdown, err error) {
	fullPath := generator.rootMdFolder + string(os.PathSeparator) + filename
	fileBytes, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}
	md := NewMarkdown(fileBytes)
	return &md, nil
}

// ConvertMdStrToHTML simply convert markdown string to html string
func ConvertMdStrToHTML(mdStr string) string {
	var md Markdown
	md = NewMarkdown([]byte(mdStr))
	return md.htmlRepresentation
}

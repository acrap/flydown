package flydown

import (
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

# flydown
![flydown_screenshot](https://i.ibb.co/Fx4j29v/Screenshot-from-2019-12-20-15-38-03.png)

Simple, cross-platform and lightweight markdown server written on Go. Renders HTML from markdown on the fly. 

In addition to the standard syntax flydown supports some extensions, such as: tables, fenced code blocks, definition lists etc. The full list you can find [in this chapter in gomarkdown(backend) readme](https://github.com/gomarkdown/markdown#extensions).

## Current status

The current status is Pre-Beta. The main goal of the project is to serve internal documentation for developer teams (Knowledge Base). 
Flydown also can be used to serve online books, like a [Gitbook](https://github.com/GitbookIO/gitbook). Unfortunately, self-hosted Gitbook version is no longer active development. Another problem, that Gitbook is not ideal for  

## Customization
In current status, flydown doesn't have sheer customization possibilities. 
But in future releases, it's going to support custom CSS styles and other stuff, like collaborative editing from the web, discussions etc.

## Search capabilities 
flydown currently supports only simple case insensitive search requests. 

## Currently supported OS
* Linux

> Windows support is untested yet

## Usage 

The repository contains an example folder with markdown. To serve it, go to the source directory and run:
```
go run ./flydown.go --share_folder=example_md
```
You can view the result: [http://localhost:8080](http://localhost:8080)

### Markdown folder structure

Markdown folder should have a structure similar to Gitbook. 

It should have at least:
* *summary.md* - summary, that will be visualized on the left side of the web page
* *readme.md* - the main page

## summary.md example

```markdown
* [Getting Started](/md/readme.md)
* [Dependencies](md/deps.md)
```


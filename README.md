# flydown
![flydown_screenshot](https://i.ibb.co/Fx4j29v/Screenshot-from-2019-12-20-15-38-03.png)

Simple, cross-platform and lightweight markdown server written on Go. Renders html from markdown on the fly. 

## Current status

Current status is Pre-Beta. The main goal of the project is serve internal documentation for developer teams (Knowledge Base). 
Flydown also can be used to serve online books, like a Gitbook.

## Customization
In current status flydown doesn't have a sheere customization possibilities. 
But in future releases it's going to support custom CSS styles and other stuff, like collaborative editing from the web, discussions etc.

## Search capabilities 
Flydown currently supports only simple case insensetive search requests. 

## Currently supported OS
* Linux

> Windows support is untested yet


## Usage 

The repository contains example folder with markdown. To serve it, go to the source directory and run:
```
go run ./flydown.go --share_folder=example_md
```
You can view the result on: [http://localhost:8080](http://localhost:8080)

### Markdown folder structure

Markdown folder should have a structure similar with Gitbook. 

It should have at least:
* *summary.md* (summary, that will be visualized on the left side of the web page) 
* *readme.md* (the main page) inside.


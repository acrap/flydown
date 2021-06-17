![badge](https://action-badges.now.sh/acrap/flydown?action=Go)

# flydown
![logo](https://i.ibb.co/sq28KyP/flydown-logo-small.png)

Simple, cross-platform and lightweight markdown server written on Go. Renders HTML from markdown on the fly. 

In addition to the standard syntax flydown supports some extensions, such as: tables, fenced code blocks, definition lists etc. The full list you can find [in this chapter in gomarkdown(backend) readme](https://github.com/gomarkdown/markdown#extensions).

[![flydown.gif](https://s5.gifyu.com/images/flydown4f30dd47ccd462e7.gif)](https://gifyu.com/image/mJP3)

The video is scaled in purpose to make the content more visible. To check the original scale see the image by [the link](https://s5.gifyu.com/images/flydown_orig_scale.png).

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

## Installation

Download the latest binary release for your architecture here: [https://github.com/acrap/flydown/releases](https://github.com/acrap/flydown/releases)

```bash
wget https://github.com/acrap/flydown/releases/download/0.1.2/flydown_0.1.2_amd64.tar.xz
```

Extract the archive:

```bash
tar -xf flydown_0.1.2_amd64.tar.xz
```

Go to the directory and run install.sh script as superuser (it copies some files to `/usr/share/flydown` and binary to `/usr/bin`)

```bash
cd flydown_0.1.2_amd64
sudo ./install.sh
```

That's it. Now you can run flydown to serve it's own documentation as simple as:

```bash
flydown serve
```

## Usage 

The repository contains an example folder with markdown. To serve it, go to the source directory and run:
```bash
go run ./flydown.go serve --shareFolder=doc 
```
You can view the result using the address that will appear in the message on terminal: 

```
Served on http://127.0.0.1:8080
```
So [http://127.0.0.1:8080](http://127.0.0.1:8080) is address used by default.

Here is the full list of flags for serve command:

```
Flags:
  -n, --bookName string   Pass the name of your book (default "My book")
  -f, --folder string     Directory with markdown content to host (flydown documentation by default) 
  -h, --help              help for serve
  -i, --ip string         Pass the IP addr (default "127.0.0.1")
  -p, --port int          Pass the port (default 8080)
  -v, --verbose           Enable verbose
```

### Markdown folder structure

Markdown folder should have a structure similar to Gitbook. 

It should have at least:
* *summary.md* - summary, that will be visualized on the left side of the web page
* *readme.md* - the main page

## summary.md example

```markdown

* [Getting Started](readme.md)
* [Dependencies](deps.md)

```


DESTDIR ?= /usr/share/flydown/
LINKER_FLAGS = -ldflags="-s -w"

.PHONY: get_deps

get_deps:
	go get . 

compile : 
	go build $(LINKER_FLAGS) flydown.go

# you need to install upx to your host and then goupx with:
# go get github.com/pwaller/goupx
compress: 
	goupx flydown

compile_arm5:
	GOARCH=arm GOARM=5 go build $(LINKER_FLAGS) flydown.go 

compile_arm7:
	GOARCH=arm GOARM=7 go build $(LINKER_FLAGS) flydown.go 

compile_arm8:
	GOARCH=arm64 go build $(LINKER_FLAGS) flydown.go 

install:
	install -d $(DESTDIR)
	cp -f flydown /usr/bin
	cp -r doc $(DESTDIR)
	cp -r static $(DESTDIR)
	cp -r templates $(DESTDIR)

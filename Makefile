DESTDIR ?= /usr/share/flydown/
REL_DIR = ./release
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

release:
ifndef VERSION
	@echo "specify the version"
	exit 1
endif
ifndef ARCH
	@echo "specify the arch"
	exit 1
endif
	install -d $(REL_DIR)/flydown_$(VERSION)_$(ARCH)
	cp -f flydown $(REL_DIR)/flydown_$(VERSION)_$(ARCH)
	cp -f install.sh $(REL_DIR)/flydown_$(VERSION)_$(ARCH)
	cp -af doc $(REL_DIR)/flydown_$(VERSION)_$(ARCH)
	cp -af static $(REL_DIR)/flydown_$(VERSION)_$(ARCH)
	cp -af templates $(REL_DIR)/flydown_$(VERSION)_$(ARCH)
	cd $(REL_DIR);tar cf - flydown_$(VERSION)_$(ARCH)/ | xz -z - > flydown_$(VERSION)_$(ARCH).tar.xz
	rm -rf $(REL_DIR)/flydown_$(VERSION)_$(ARCH)

clear_release:
	rm -rf $(REL_DIR)


install:
	install -d $(DESTDIR)
	cp -f flydown /usr/bin
	cp -r doc $(DESTDIR)
	cp -r static $(DESTDIR)
	cp -r templates $(DESTDIR)

GO ?= go

# command to build and run on the local OS.
GO_BUILD = go build

# command to compiling the distributable. Specify GOOS and GOARCH for the target OS.
GO_DIST = GOOS=linux GOARCH=amd64 go build

.PHONY: dist

all: clean prepare dist

clean:
	rm -rf dist

prepare:
	mkdir -p dist

deps:
	$GO) get -t ./...

dist: urlencode urldecode

urlencode:
	$(GO_DIST) -o dist/urlencode cmd/urlencode/main.go

urldecode:
	$(GO_DIST) -o dist/urldecode cmd/urldecode/main.go

install:
	$(GO) get ./cmd/...

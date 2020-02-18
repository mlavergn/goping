###############################################
#
# Makefile
#
###############################################

.DEFAULT_GOAL := build

.PHONY: test

VERSION := 0.1.0

ver:
	@sed -i '' 's/^const Version = "[0-9]\{1,3\}.[0-9]\{1,3\}.[0-9]\{1,3\}"/const Version = "${VERSION}"/' src/ping/ping.go

lint:
	$(shell go env GOPATH)/bin/golint ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

build:
	go build -v ./...

clean:
	go clean ...

demo: build
	go build -o demo cmd/demo.go

test: build
	go test -v -count=1 ./...

github:
	open "https://github.com/mlavergn/goping"

release:
	zip -r goping.zip LICENSE README.md Makefile cmd src
	hub release create -m "${VERSION} - Go Ping" -a goping.zip -t master "v${VERSION}"
	open "https://github.com/mlavergn/goping/releases"

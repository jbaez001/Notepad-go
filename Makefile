VERSION ?= $(shell git describe --tags --always --dirty)
DATE_COMPILED ?= $(shell date +%FT%T%z)
GO_PACKAGE ?= $(shell go list -m)

all:
	go build -o bin/Notepad \
		-ldflags "-s -w -X ${GO_PACKAGE}/pkg/version.Version=${VERSION} -X ${GO_PACKAGE}/pkg/version.DateCompiled=${DATE_COMPILED}" .

install:
	go install -installsuffix "static" \
		-ldflags "-s -w -X ${GO_PACKAGE}/pkg/version.Version=${VERSION} -X ${GO_PACKAGE}/pkg/version.DateCompiled=${DATE_COMPILED}" .

race:
	go build -race -o bin/Notepad \
		-ldflags "-s -w -X ${GO_PACKAGE}/pkg/version.Version=${VERSION} -X ${GO_PACKAGE}/pkg/version.DateCompiled=${DATE_COMPILED}" .
GOCMD=go
GOBUILD=$(GOCMD) build
GOFMT=gofmt
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

BINARY=td

.PHONY: test build clean check fmt
all: fmt check test build

fmt:
	$(GOFMT) -s -l -e -w .

check:
	errcheck -exclude errcheck_excludes.txt -asserts -verbose ./...
	go vet .
	golint

test:
	$(GOTEST) -v ./...

build:
	$(GOBUILD) -o $(BINARY) main.go

clean:
	$(GOCLEAN)
	rm -f $(BINARY)

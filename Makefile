GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

BINARY=td

.PHONY: test build clean check
all: check test build

build:
	$(GOBUILD) -o $(BINARY) main.go

test:
	$(GOTEST) -v ./...

check:
	errcheck -exclude errcheck_excludes.txt -asserts -verbose ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY)

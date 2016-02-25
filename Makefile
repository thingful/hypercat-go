# Simple Makefile for testing hypercat-go

ARTEFACT_DIR = ./tmp
GOCMD = go
GOTEST = $(GOCMD) test -coverprofile=$(ARTEFACT_DIR)/cover.out
GOLINT = gometalinter
GOCOVER = $(GOCMD) tool cover

default: test

setup:
	$(GOCMD) get -u github.com/alecthomas/gometalinter
	$(GOLINT) --install --update

test:
	mkdir -p $(ARTEFACT_DIR)
	$(GOTEST) ./...

coverage: test
	mkdir -p $(ARTEFACT_DIR)
	$(GOCOVER) -func=$(ARTEFACT_DIR)/cover.out

html: test
	mkdir -p $(ARTEFACT_DIR)
	$(GOCOVER) -html=$(ARTEFACT_DIR)/cover.out -o $(ARTEFACT_DIR)/coverage.html

lint:
	$(GOLINT) ./...

clean:
	rm -rf $(ARTEFACT_DIR)

full: lint coverage html

.PHONY: setup test lint coverage full html clean

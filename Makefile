# Simple Makefile for testing hypercat-go
#
ARTEFACT_DIR = ./tmp
GOCMD = go
GOTEST = $(GOCMD) test -coverprofile=$(ARTEFACT_DIR)/cover.out
GOLINT = golint
GOCOVER = $(GOCMD) tool cover

default: test

echo:
	echo $(ARTEFACT_DIR)

setup:
	$(GOCMD) get -u github.com/golang/lint/golint

test:
	$(GOTEST) ./...

coverage:
	$(GOCOVER) -func=$(ARTEFACT_DIR)/cover.out

html:
	$(GOCOVER) -html=$(ARTEFACT_DIR)/cover.out -o $(ARTEFACT_DIR)/coverage.html

lint:
	$(GOLINT) ./...

clean:
	rm -f $(ARTEFACT_DIR)/cover.out
	rm -f $(ARTEFACT_DIR)/coverage.html

full: test coverage html

.PHONY: setup test lint coverage full html clean

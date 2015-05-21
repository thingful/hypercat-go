
GOCMD = go
GOTEST = $(GOCMD) test -v --cover
GOLINT = golint

default: test

setup:
	$(GOCMD) get -u github.com/golang/lint/golint

test:
	$(GOTEST) ./...

lint:
	$(GOLINT) ./...

.PHONY: setup test lint

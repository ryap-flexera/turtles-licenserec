#!/bin/bash

GOCMD=go
GOACMD=goa
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
GOAGEN=$(GOACMD) gen

currentDIR := $(shell pwd)
cmdDIR := $(currentDIR)/cmd/licenseapi_server
#designDIR := /turtles-licenserec/license_rec_api/design
designDIR := github.com/flexera/turtles-licenserec/license_rec_api/design
testResults := $(shell go test -v ./... | grep -c "FAIL")

default:
	@echo "Downloading dependencies."
	$(GOGET) -v ./...
	$(GOAGEN) $(designDIR)
	$(GOBUILD) $(cmdDIR)
	@if [ ${testResults} -ne 0 ]; then echo "Test Failure. Build Failed."; echo "Please run go test ./..."; exit; fi;

	
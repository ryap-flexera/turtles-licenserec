#!/bin/bash

.PHONY: license_rec_api

GOPATH := ${GOPATH}
SUBDIRS := $(filter-out _infra/., $(wildcard */.))

default:
	@echo "================ Building project ================"
	@echo "GOPATH: ${GOPATH}"
	@echo "Current working directory: ${currentDIR}"
	$(SUBDIRS)

	$(SUBDIRS):
		$(MAKE) -C $(SUBDIRS)

	

	





# Import environment file
include .env
# Source all variables in environment file
# This only runs in the make command shell
# so won't muddy up, e.g. your login shell
export $(shell sed 's/=.*//' .env)
.PHONY:	lint test

all: lint test

lint:
	go vet ./
	go fmt ./

test: lint
	go test -count=1 -v -cover --race -tags="unittests" ./

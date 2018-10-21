INFO_COLOR=\033[1;34m
RESET=\033[0m
BOLD=\033[1m
TEST ?= $(shell go list ./... | grep -v vendor)
VERSION = $(shell cat version)
REVISION = $(shell git describe --always)
ifeq ("$(shell uname)","Darwin")
GO ?= GO111MODULE=on go
else
GO ?= GO111MODULE=on /usr/local/go/bin/go
endif

server: deps
	$(GO) run *.go

.PHONY: default test deps
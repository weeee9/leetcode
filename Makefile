#!/bin/bash
PROJECT="LEETCODE"

SOURCE ?= $(shell find . -type f -name '*.go')

test:
	@printf "===RUN PROBLEM TEST===\n"
	@go test -v -tags="sqlite json1" ./...
	@echo "===\033[32m OK \033[0m==="
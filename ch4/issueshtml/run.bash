#!/usr/bin/env bash

go run main.go repo:golang/go commenter:gopherbot json encoder > issues.html
go run main.go repo:golang/go 3133 10535 > issues2.html

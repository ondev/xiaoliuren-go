#!/bin/sh
set -v
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o windows.exe main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o darwin main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o linux main.go 


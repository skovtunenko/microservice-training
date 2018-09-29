#!/usr/bin/env bash

GOOS=linux GOARCH=amd64 go build -o uploadDocuments ./cmd/uploadDocuments/main.go
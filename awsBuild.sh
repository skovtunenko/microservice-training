#!/usr/bin/env bash

GOOS=linux GOARCH=amd64 go build -o usersManager-aws ./cmd/usersManager/main.go
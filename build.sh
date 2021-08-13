#!/usr/bin/env bash

echo "Begin to build"

CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o gomars ./src/main.go

echo "Build success"
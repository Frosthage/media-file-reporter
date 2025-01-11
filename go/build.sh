#!/bin/bash
# GOOS=windows GOARCH=amd64 go build -o fillistaren.exe main.go
docker run --rm -v "$(pwd):/app" -w /app windows-go-builder go build -o fillistaren
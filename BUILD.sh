#!/bin/bash

GOOS=linux GOARCH=amd64 go build -o out/smr-list-utils.elf
GOOS=windows GOARCH=amd64 go build -o out/smr-list-utils.exe
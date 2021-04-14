#!/bin/bash


flags="-w -s -X config.BuildTime=`date -u '+%Y-%m-%d_%I:%M:%S%p'` "
GOOS=linux go build -ldflags "$flags" -x -o build-version main.go
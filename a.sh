#!/bin/bash

set -x 

flags="-w -s -X wingCA/config.BuildTime=`date '+%Y-%m-%d.%H:%M:%S'`"
#echo $flags
#GOOS=linux go build -ldflags "$flags" -x -o build-version main.go
go build -ldflags "$flags" -o wingCA main.go
#GOOS=linux go build -ldflags "$flags" -o wingCA_linux main.go
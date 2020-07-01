#!/bin/bash

export GOPATH=$PWD:$GOPATH

cd src/main
go build -o main

echo -e "completed. see src/main/\n\n"
ls -l $PWD





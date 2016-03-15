#!/bin/bash

rm -rf example/search
rm -rf example/create
go-bindata -pkg templates -o templates/bindata.go templates/*.tmpl
go install ./...
go generate ./...

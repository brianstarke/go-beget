#!/bin/bash

rm -rf example/search
rm -rf example/create
rm -rf example/update
go-bindata -pkg templates -o templates/bindata.go templates/*.tmpl
go install ./...
go generate ./...

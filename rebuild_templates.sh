#!/bin/bash

rm -rf example/search
go-bindata -pkg templates -o templates/bindata.go templates/*.tmpl
go install ./...

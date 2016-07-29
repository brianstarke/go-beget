#!/bin/bash

rm example/searchEnums.go
rm example/thingSearchRequest.go

go-bindata -pkg templates -o templates/bindata.go templates/*.tmpl

go install ./...
go generate ./...

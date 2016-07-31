#!/bin/bash

rm example/searchEnums.go
rm example/thingSearch.go
rm example/thingCreate.go

go-bindata -pkg templates -o templates/bindata.go templates/*.tmpl

go install ./...
go generate ./...

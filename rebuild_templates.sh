#!/bin/bash

rm example/searchRequestEnums.go
rm example/thingSearchRequest.go

go-bindata -pkg templates -o templates/bindata.go templates/*.tmpl

go install ./...
go generate ./...

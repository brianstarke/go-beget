#!/bin/bash

rm example/search.generated.go
rm example/thingFields.generated.go
rm example/thingSearch.generated.go
rm example/thingCreate.generated.go
rm example/thingUpdate.generated.go

go-bindata -pkg templates -o templates/bindata.go templates/*.tmpl

go install ./...
go generate ./...

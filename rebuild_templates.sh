#!/bin/bash

go-bindata -pkg templates -o templates/bindata.go templates/*.tmpl

# go-begetter

[![](https://godoc.org/github.com/brianstarke/go-begetter?status.svg)](http://godoc.org/github.com/brianstarke/go-begetter)
[![Build Status](https://travis-ci.org/brianstarke/go-begetter.svg)](https://travis-ci.org/brianstarke/go-begetter)

Generate serializable and typesafe search/create/update/delete requests to pass to services.

This is still a work in progress.

### Generator templates

`go-begetter` uses the excellent [go-bindata](https://github.com/jteeuwen/go-bindata) tool to compile it's templates in to the executable for easier command line use.  Therefore, if you make any changes to the templates, you must re-compile them.

`cd templates && go-bindata -pkg templates *.tmpl`

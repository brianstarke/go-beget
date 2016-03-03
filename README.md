# go-begetter

[![](https://godoc.org/github.com/brianstarke/go-beget?status.svg)](http://godoc.org/github.com/brianstarke/go-beget)

Generate serializable and typesafe search/create/update/delete requests to pass to services.

This is still a work in progress.

### Generator templates

`go-beget` uses the excellent [go-bindata](https://github.com/jteeuwen/go-bindata) tool to compile it's templates in to the executable for easier command line use.  

Therefore, if you make any changes to the templates, you must install `go-bindata` (`go get -u github.com/jteeuwen/go-bindata/...`).

Then `./rebuild_templates.sh` from the root of this project.

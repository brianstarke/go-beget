# go-begetter

[![](https://godoc.org/gitlab.com/brianstarke/go-begetter?status.svg)](http://godoc.org/gitlab.com/brianstarke/go-begetter)
[![build status](https://gitlab.com/ci/projects/20838/status.png?ref=master)](https://gitlab.com/ci/projects/20838?ref=master)

Generate serializable and typesafe search/create/update/delete requests to pass to services.

This is still a work in progress.

### Generator templates

`go-begetter` uses the excellent [go-bindata](https://github.com/jteeuwen/go-bindata) tool to compile it's templates in to the executable for easier command line use.  

Therefore, if you make any changes to the templates, you must install `go-bindata` (`go get -u github.com/jteeuwen/go-bindata/...`).

Then `./rebuild_templates.sh` from the root of this project.

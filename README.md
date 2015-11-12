# go-begetter
Generate serializable and typesafe search/create/update/delete requests to pass to services.

### Generator templates

`go-begetter` uses the excellent [go-bindata](https://github.com/jteeuwen/go-bindata) tool to compile it's templates in to the executable for easier command line use.  Therefore, if you make any changes to the templates, you must re-compile them.

`cd templates && go-bindata -pkg templates *.tmpl`

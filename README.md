# go-beget

[![GoDoc](https://godoc.org/github.com/brianstarke/go-beget?status.svg)](https://godoc.org/github.com/brianstarke/go-beget)

Generate serializable and typesafe search/create/update/delete requests to pass to services.

This is still a work in progress.

### Generator templates

`go-beget` uses the excellent [go-bindata](https://github.com/jteeuwen/go-bindata) tool to compile it's templates in to the executable for easier command line use.  

Therefore, if you make any changes to the templates, you must install `go-bindata` (`go get -u github.com/jteeuwen/go-bindata/...`).

Then `./rebuild_templates.sh` from the root of this project.

### Search Requests

Example of using a generated search request.

```go
var sr search.ThingSearchRequest

sr.Fields = []search.ThingField{
  search.ThingColor,
  search.ThingHeight,
}

sr.AddFilter(
  search.ThingColor,
  "red",
  searcher.EQ,
  searcher.AND)

sr.SetOrderBy(search.ThingHeight, false)

sr.Limit = 10

jsonb, _ := json.MarshalIndent(sr, "", "  ")

log.Println(string(jsonb))

searcher.GenerateSelectSQL(&sr)
```

This outputs

```
2016/03/03 15:47:05 {
  "limit": 10,
  "filters": [
    {
      "field": "color",
      "value": "red",
      "operator": "EQ",
      "condition": "AND"
    }
  ],
  "orderBy": {
    "field": "height",
    "desc": false
  },
  "fields": [
    "color",
    "height"
  ]
}
[go-beget/searcher] SELECT SQL generated - SELECT color, height FROM things WHERE (color = $1) ORDER BY height LIMIT 10
```

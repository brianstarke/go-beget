# go-beget

[![GoDoc](https://godoc.org/github.com/brianstarke/go-beget?status.svg)](https://godoc.org/github.com/brianstarke/go-beget)

Generate serializable and typesafe search requests to pass around.

This is still a work in progress.

### The basics

Add `beget` tags to a struct for which you want to create `SearchRequests` for and add the `go:generate` comment.

`struct` is the struct you want the `go-beget` generator to look at, `table` is the name of the database table (if you want to use optional SQL statement generation).  `repos` takes a comma separated list of Repositories helpers you'd like to auto-generate.  They're called Repos because they kinda follows the [Repository pattern](http://www.giorgiosironi.com/2009/10/repository-pattern.html).  Currently, only *sql* is available - and I've only tested that on PostgreSQL.

```go
package types

//go:generate searcher -struct=Thing -table=things -repos=db

// Thing has characteristics
type Thing struct {
	Color       string `beget:"search" json:"color" db:"color"`
	Description string `beget:"search" json:"description" db:"description"`
	Length      int    `beget:"search" json:"length" db:"length"`
	Height      int    `beget:"search" json:"height" db:"height"`
}
```

NOTE: For now, `go-beget` expects that you keep your types in their own package and will generate it's code into a directory called `search` as a peer of wherever your types are.

Run `go generate ./...`

```
[go-beget/searcher] Generating searcher for Thing
[go-beget/searcher] search request generated ../search/thing_searchrequest.go
```

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

*TODO document the repository usage*

### Generator templates

`go-beget` uses the excellent [go-bindata](https://github.com/jteeuwen/go-bindata) tool to compile it's templates in to the executable for easier command line use.  

Therefore, if you make any changes to the templates, you must install `go-bindata` (`go get -u github.com/jteeuwen/go-bindata/...`).

Then `./rebuild_templates.sh` from the root of this project.

### TODO

- Creators
- Updaters
- still on the fence about Deleters...
- tagging ID fields to do easier GetByID queries
- make all generated code pass golint

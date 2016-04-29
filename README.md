# go-beget

[![GoDoc](https://godoc.org/github.com/brianstarke/go-beget?status.svg)](https://godoc.org/github.com/brianstarke/go-beget)
[![Build Status](https://travis-ci.org/brianstarke/go-beget.svg?branch=master)](https://travis-ci.org/brianstarke/go-beget)

Generate serializable and typesafe search requests to pass around.

This is still a work in progress.

### The basics

Add `beget` tags to a struct for which you want to create `SearchRequests` for and add the `go:generate` comment.

`struct` is the struct you want the `go-beget` generator to look at, `table` is the name of the database table (if you want to use optional SQL statement generation).  `repos` takes a comma separated list of Repositories helpers you'd like to auto-generate.  They're called Repos because they kinda follows the [Repository pattern](http://www.giorgiosironi.com/2009/10/repository-pattern.html).  Currently, only *sql* is available - and I've only tested that on PostgreSQL.

Within the `beget` tag, add **search** to the fields you'd like to be searchable, **create** to the fields you'd like to be inserted by the generated **Creator** and **update** to the fields you want to allow to updated via an **Updater**.  Generally you'd leave the `ID` field without a **create** tag if your database is assigning the IDs.

```go
package types

//go:generate searcher -struct=Thing -table=things -repos=sql
//go:generate creator -struct=Thing -table=things -repos=sql
//go:generate updater -struct=Thing -table=things -repos=sql

// Thing has characteristics
type Thing struct {
	ID          int64  `beget:"search" json:"id" db:"id"`
	Color       string `beget:"search,create" json:"color" db:"color"`
	Description string `beget:"search,create" json:"description" db:"description"`
	Length      int    `beget:"search,create,update" json:"length" db:"length"`
	Height      int    `beget:"search,create,update" json:"height" db:"height"`
}
```

NOTE: For now, `go-beget` expects that you keep your types in their own package and will generate it's code into a directory called `search` (or `update` or `create`) as a peer of wherever your types are.

Run `go generate ./...`

```
[go-beget/searcher] Generating searcher for Thing
[go-beget/searcher] ../search does not exist, I'll create it
[go-beget/searcher] SearchRequest generated ../search/thingSearchRequest.go
[go-beget/searcher] SearchRepo generated ../search/thingSearchRepo.go
[go-beget/creator] Generating creator for Thing
[go-beget/creator] ../create does not exist, I'll create it
[go-beget/creator] CreateRepo generated ../create/thingCreateRepo.go
[go-beget/updater] Generating updater for Thing
[go-beget/updater] UpdateRequest generated ../update/thingUpdateRequest.go
[go-beget/updater] UpdateRepo generated ../update/thingUpdateRepo.go
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

### Using Searchers

You can use a generated **searcher** to execute **SearchRequests** or pull single results by a particular field.  The latter is useful for doing a get by ID sort of thing.

```go
// construct the search request
var sr search.ThingSearchRequest

sr.AddFilter(
  search.ThingColor,
  "red",
  searcher.EQ,
  searcher.AND)

sr.SetOrderBy(search.ThingHeight, false)
sr.Limit = 10
sr.Offset = 2

// run the search request (where db is a working *sqlx.DB instance)
things, err := search.NewSQLThingSearcher(db).Search(sr)
// do things with results, check error etc...
```

Or to simply just get a single result by a particular field...

```go
thing, err := search.NewSQLThingSearcher(db).GetByField(search.ThingID, 16)
// do things with result, check error etc...
```

*TODO document the creator usage*

*TODO document the updater usage*

### Generator templates

`go-beget` uses the excellent [go-bindata](https://github.com/jteeuwen/go-bindata) tool to compile it's templates in to the executable for easier command line use.  

Therefore, if you make any changes to the templates, you must install `go-bindata` (`go get -u github.com/jteeuwen/go-bindata/...`).

Then `./rebuild_templates.sh` from the root of this project.

### TODO

- add a Get method to return single result
- add an AddField method
- tagging ID fields to do easier GetByID queries?
- make all generated code pass golint

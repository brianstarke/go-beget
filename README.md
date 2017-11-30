
# NOTICE 

This repo is no longer maintained, I personally vastly prefer [https://gnorm.org/](gnorm) over my own solution and recommend checking it out.

# go-beget

[![GoDoc](https://godoc.org/github.com/brianstarke/go-beget?status.svg)](https://godoc.org/github.com/brianstarke/go-beget)
[![Build Status](https://travis-ci.org/brianstarke/go-beget.svg?branch=master)](https://travis-ci.org/brianstarke/go-beget)

Generate serializable and (somewhat) typesafe search requests to pass around.

The intent is not to recreate Rails, but rather provide a baseline for you to fork to help generate all your boilerplate just the way you like it.

### The basics

Install `beget` by running `go get -u github.com/brianstarke/go-beget/...`

Add `beget` tags to a struct for which you want to create `SearchRequests` for and add the `go:generate` comment.

`struct` is the struct you want the `go-beget` generator to look at, `table` is the name of the database table (if you want to use optional SQL statement generation).  `impls` takes a comma separated list of helpers you'd like to auto-generate.  

- **sql** will generate a **SQLSearcher**/**SQLCreator**/**SQLUpdater** (currently, **sql** implementation has only been tested on PostgreSQL).
- **gin** will generate [gin gonic](https://github.com/gin-gonic) handlers that use the generated **SQLSearcher**/**SQLCreator**/**SQLUpdater** stuff.

Within the `beget` tag, add **search** to the fields you'd like to be searchable, **create** to the fields you'd like to be inserted by the generated **Creator** and **update** to the fields you want to allow to updated via an **Updater**.  

Generally you'd leave the `ID` field without a **create** tag if your database is assigning the IDs.

```go
package types

//go:generate searcher -struct=Thing -table=things -impls=sql,gin
//go:generate creator -struct=Thing -table=things -impls=sql,gin
//go:generate updater -struct=Thing -table=things -impls=sql,gin

// Thing has characteristics
type Thing struct {
	ID          int64  `beget:"search" json:"id" db:"id"`
	Color       string `beget:"search,create" json:"color" db:"color"`
	Description string `beget:"search,create" json:"description" db:"description"`
	Length      int    `beget:"search,create,update" json:"length" db:"length"`
	Height      int    `beget:"search,create,update" json:"height" db:"height"`
}
```

NOTE: For now, `go-beget` expects that you keep your types in some other package and will generate it's code into a directory called `search` (or `update` or `create`) as a peer of wherever the type is.

Run `go generate ./...`

```
[go-beget/searcher] Generating searcher for Thing
[go-beget/searcher] ../search does not exist, I'll create it
[go-beget/searcher] SearchRequest generated ../search/thingSearchRequest.go
[go-beget/searcher] SQLSearcher generated ../search/thingSQLSearcher.go
[go-beget/searcher] GinSearcher generated ../search/thingGinSearcher.go
[go-beget/creator] Generating creator for Thing
[go-beget/creator] ../create does not exist, I'll create it
[go-beget/creator] SQLCreator generated ../create/thingSQLCreator.go
[go-beget/creator] GinCreator generated ../create/thingGinCreator.go
[go-beget/updater] Generating updater for Thing
[go-beget/updater] ../update does not exist, I'll create it
[go-beget/updater] UpdateRequest generated ../update/thingUpdateRequest.go
[go-beget/updater] SQLUpdater generated ../update/thingSQLUpdater.go
[go-beget/updater] GinUpdater generated ../update/thingGinUpdater.go
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

### Using SQLSearchers

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

### Using GinSearchers/GinCreators

If you also generated the [gin gonic](https://github.com/gin-gonic) handlers, you can wire them up in your routes like so:

```go
router.PUT("/thing", create.NewThingCreateHandler(db)) 				// create
router.POST("/thing", search.NewThingSearchHandler(db))  			// search
router.GET("/thing/:id", search.NewThingGetByIDHandler(db)) 	// get
router.PATCH("/thing", updater.NewThingUpdateHandler(db)) 		// update
```

The generated routes expect JSON and return JSON.  On error, they return status code **500** and the error in a JSON struct like this:

```javascript
{
	"error": "Your error here"
}
```

On success they return **200** and the result as JSON.

## Sample JSON SearchRequest

```javascript
{
    "filters":[{
      "field": "color",
      "value": "RED",
      "operator": "EQ",
      "condition": "AND"
    },
		{
			"field": "size",
			"value": 13,
			"operator": "GT",
			"condition": "AND"
		}],
    "fields": ["height", "type"],
    "orderBy": {
        "field": "createdAt",
        "desc": true
    },
    "limit":20,
    "offset": 2
}
```

*TODO document the creator usage*

*TODO document the updater usage*

### Generator templates

`go-beget` uses the excellent [go-bindata](https://github.com/jteeuwen/go-bindata) tool to compile it's templates in to the executable for easier command line use.  

Therefore, if you make any changes to the templates, you must install `go-bindata` (`go get -u github.com/jteeuwen/go-bindata/...`).

Then `./rebuild_templates.sh` from the root of this project.

### TODO

- add an AddField method
- make all generated code pass golint
- break dependency on sqlx and gin?
- generate deleters?
- add fake/mock implementations for testing
- support XML?

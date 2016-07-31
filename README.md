# go-beget

[![GoDoc](https://godoc.org/github.com/brianstarke/go-beget?status.svg)](https://godoc.org/github.com/brianstarke/go-beget)
[![Build Status](https://travis-ci.org/brianstarke/go-beget.svg?branch=master)](https://travis-ci.org/brianstarke/go-beget)

Generate serializable and (somewhat) typesafe search requests to pass around.

The intent is not to create an ORM, but rather provide a forkable baseline to help generate all your boilerplate just the way you like it.

## Getting started / Overview

Install `beget` by running `go get -u github.com/brianstarke/go-beget/generator`

Add the `go:generate beget` comment, tag all fields with `db` and `json` tags if you want something other than `ToLower`'d DB column names and JSON field names.

`struct` is the struct you want the `go-beget` generator to look at, `table` is the name of the database table to be used in SQL statement generation/execution.

- _(currently, **sql** statement generation/execution only works on PostgreSQL)_

Create functions will automatically omit `ID` and `CreatedAt` fields from the **INSERT** statement.  If you wish to omit additional fields, pass them in to the `omitFromInsert` flag (separated by commas).

```go
package model

//go:generate go-beget -struct=Thing -table=things -omitFromInsert=Length,Height

// Thing is a normal thing
type Thing struct {
	ID          int64     `json:"id" db:"id"`
	Color       string    `json:"color" db:"color"`
	Description string    `json:"description" db:"description"`
	Length      int       `json:"length" db:"length"`
	Height      int       `json:"height" db:"height"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at_utc"`
}
```

`go-beget` currently generates it's code into the same package and directory as the one your struct is in.

Run `go generate ./...`

```
[go-beget/generator] Generating methods for Thing
[go-beget/generator] create generated		[thingCreate.generated.go]
[go-beget/generator] fields generated		[thingFields.generated.go]
[go-beget/generator] enums generated		[search.generated.go]
[go-beget/generator] update generated		[thingUpdate.generated.go]
[go-beget/generator] search generated		[thingSearch.generated.go]
```

## Search Requests

Using a generated SearchRequest like so...

```go
package main

import (
	"encoding/json"
	"log"

	"github.com/brianstarke/go-beget/example"
)

func main() {
	var sr example.ThingSearchRequest

	sr.
		AddFields(example.ThingColor, example.ThingHeight).
		AddFilter(example.ThingColor, "red", func(f *example.ThingFilter) {
			f.Operator = example.Eq
			f.Condition = example.Or
		}).
		SetOrderBy(example.ThingHeight, false).
		SetLimit(10)

	jsonb, _ := json.MarshalIndent(sr, "", "  ")

	log.Println(string(jsonb))

	sql, _, _ := sr.GenerateSelectSQL()

	log.Println(sql)
}
```

Would output:

```
2016/07/31 17:00:39 {
  "fields": [
    "color",
    "height"
  ],
  "filters": [
    {
      "field": "color",
      "value": "red",
      "operator": "eq",
      "condition": "OR"
    }
  ],
  "orderBy": {
    "field": "height",
    "isDescending": false
  },
  "limit": 10,
  "offset": 0
}
2016/07/31 17:00:39 SELECT color, height FROM things WHERE (color = $1) ORDER BY height LIMIT 10
```

### Using SQLSearchers

You can use a generated **SearchRequests** to execute a search for multiple results.

```go
// construct the search request
var sr search.ThingSearchRequest

sr.
	AddFilter(search.ThingColor, "red").
	SetOrderBy(search.ThingHeight, false).
	SetLimit(10).
	SetOffset(10)

// run the search request (where db is a working *sql.DB instance)
var results []type.Thing
err := sr.ExecuteSearch(db, &results)
// do things with results, check error etc...

// or get a count of all results...
count, err := sr.ExecuteCount(db)

// or pull a single result by ID
// TODO generate that
```

### Using HTTP handlers (totally optional)

`beget` also generates HTTP handlers, you can wire them up in your routes like so:

```go
func main() {
  http.HandleFunc("/", search.NewThingSearchHandlerFunc(db))
  http.ListenAndServe(":8080", nil)
}
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
        "isDescending": true
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

- make all generated code pass golint
- break dependency on sqlx?
- generate deleters?
- add fake/mock implementations for testing?

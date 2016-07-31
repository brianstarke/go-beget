# go-beget

[![GoDoc](https://godoc.org/github.com/brianstarke/go-beget?status.svg)](https://godoc.org/github.com/brianstarke/go-beget)
[![Build Status](https://travis-ci.org/brianstarke/go-beget.svg?branch=master)](https://travis-ci.org/brianstarke/go-beget)

Generate serializable and (somewhat) typesafe search requests to pass around.

The intent is not to create an ORM, but rather provide a forkable baseline to help generate all your boilerplate just the way you like it.

## Getting started / Overview

Install `beget` by running `go get -u github.com/brianstarke/go-beget`

Add the `go:generate go-beget` comment, tag all fields with `db` and `json` tags if you want something other than `ToLower`'d DB column names and JSON field names.

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
			f.Condition = example.Or
		}).
		AddFilter(example.ThingColor, "blue", func(f *example.ThingFilter) {
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

```json
{
  "fields": [
    "color",
    "height"
  ],
  "filters": [
    {
      "field": "color",
      "value": "red",
      "operator": "eq",
      "condition": "or"
    },
    {
      "field": "color",
      "value": "blue",
      "operator": "eq",
      "condition": "or"
    }
  ],
  "orderBy": {
    "field": "height",
    "isDescending": false
  },
  "limit": 10,
  "offset": 0
}
```
```sql
SELECT color, height FROM things WHERE (color = $1 OR color = $2) ORDER BY height LIMIT 10
```

By default, `AddFilter` will use Eq (Equals) as an operator and And as a condition, if you want to override those - you can pass in functions to modify the options (H/T to [Dave Cheney](http://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis http://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis))

e.g:

```go
var sr example.MahSearchRequest

sr.AddFilter(example.MahSticks, 4) // And 'sticks' = 4

sr.AddFilter(example.MahStones, 80, func(f *example.MahFilter) {
	f.Operator = example.GreaterThan
}) // And 'stones' > 80

sr.AddFilter(example.MahMarbles, 22, func(f *example.MahFilter) {
	f.Condition = example.Or
	f.Operator = example.LesserThanOrEq
}) // Or 'marbles' <= 22

sr.AddFilter(example.MahMarbles, 12, func(f *example.MahFilter) {
	f.Condition = example.Or
	f.Operator = example.GreaterThanOrEq
}) // Or 'marbles' >= 12
```

Would result in

```sql
SELECT * FROM mah_things WHERE (marbles <= $1 OR marbles >= $2) AND (sticks = $3 AND stones > $4)
```

### Using HTTP handlers (totally optional)

`go-beget` also generates HTTP handlers, you can wire them up in your routes like so:

```go
func main() {
  http.HandleFunc("/create", search.NewThingCreateHandlerFunc(db))
  http.HandleFunc("/search", search.NewThingSearchHandlerFunc(db))
  http.HandleFunc("/update", search.NewThingUpdateHandlerFunc(db))
  http.ListenAndServe(":8080", nil)
}
```

On success they return **200** and the result as JSON.

## Sample JSON SearchRequest for generated handler

```javascript
{
    "filters":[{
      "field": "color",
      "value": "RED",
      "operator": "eq",
      "condition": "and"
    },
		{
			"field": "size",
			"value": 13,
			"operator": "gt",
			"condition": "and"
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

*TODO document create usage*

*TODO document update usage*

### Generator templates

`go-beget` uses the excellent [go-bindata](https://github.com/jteeuwen/go-bindata) tool to compile it's templates in to the executable for easier command line use.  

Therefore, if you make any changes to the templates, you must install `go-bindata` (`go get -u github.com/jteeuwen/go-bindata/...`).

Then `./rebuild_templates.sh` from the root of this project.

### TODO

- add getByID function
- make ID field configurable
- make all generated code pass golint
- break dependency on sqlx?
- generate deleters?
- add fake/mock implementations for testing?

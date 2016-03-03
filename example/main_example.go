package main

import (
	"encoding/json"
	"log"

	"github.com/brianstarke/go-beget/example/search"
	"github.com/brianstarke/go-beget/searcher"
)

// Example of how to construct a search request
func main() {
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
}

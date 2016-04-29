package search

/*
	GENERATED CODE, YOUR EDITS ARE FUTILE

	generated from Thing
	using http://github.com/brianstarke/go-beget/generator/searcher
*/

import (
	"fmt"
	"strings"

	"github.com/brianstarke/go-beget/searcher"
	"github.com/jmoiron/sqlx"

	types "github.com/brianstarke/go-beget/example/types"
)

// ThingSearcher is the interface
type ThingSearcher interface {
	Search(searchRequest ThingSearchRequest) ([]types.Thing, error)
}

// SQLThingSearcher implements a SQL based searcher
type SQLThingSearcher struct {
	db *sqlx.DB
}

// NewSQLThingSearcher returns a configured repo for you
func NewSQLThingSearcher(db *sqlx.DB) ThingSearcher {
	return &SQLThingSearcher{db: db}
}

// Search converts a ThingSearchRequest in to SQL and executes it
func (r *SQLThingSearcher) Search(searchRequest ThingSearchRequest) ([]types.Thing, error) {
	results := []types.Thing{}

	sqlStr, values, err := searcher.GenerateSelectSQL(&searchRequest)

	if err != nil {
		return nil, fmt.Errorf("Error generating SQL for ThingSearchRequest : %s", err.Error())
	}

	err = r.db.Select(&results, sqlStr, values.([]interface{})...)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return results, nil
		}
	}

	return results, err
}

// GetByField is a convenience method to return just one result by matching on a single field
func (r *SQLThingSearcher) GetByField(field ThingField, value interface{}) (*types.Thing, error) {
	var searchRequest ThingSearchRequest

	searchRequest.AddFilter(
		field,
		value,
		searcher.EQ,
		searcher.AND,
	)
	searchRequest.Limit = 1

	sqlStr, values, err := searcher.GenerateSelectSQL(&searchRequest)

	if err != nil {
		return nil, fmt.Errorf("Error generating GetBy %s SQL for Thing : %s", field, err.Error())
	}

	var result types.Thing
	err = r.db.Get(&result, sqlStr, values.([]interface{})...)

	if err != nil {
		return nil, err
	}

	return &result, err
}

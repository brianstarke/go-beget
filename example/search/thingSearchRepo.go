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

// ThingSearchRepo is the interface
type ThingSearchRepo interface {
	Search(searchRequest ThingSearchRequest) ([]types.Thing, error)
}

// SQLThingSearchRepo implements a SQL based searcher
type SQLThingSearchRepo struct {
	db *sqlx.DB
}

// NewSQLThingSearchRepo returns a configured repo for you
func NewSQLThingSearchRepo(db *sqlx.DB) ThingSearchRepo {
	return &SQLThingSearchRepo{db: db}
}

// Search converts a ThingSearchRequest in to SQL and executes it
func (r *SQLThingSearchRepo) Search(searchRequest ThingSearchRequest) ([]types.Thing, error) {
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

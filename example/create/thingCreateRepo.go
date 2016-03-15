package create

/*
	GENERATED CODE, YOUR EDITS ARE FUTILE

	generated from Thing
	using http://github.com/brianstarke/go-beget/generator/creator
*/

import (
	"github.com/brianstarke/go-beget/creator"
	"github.com/jmoiron/sqlx"

	types "github.com/brianstarke/go-beget/example/types"
)

// ThingCreator defines the interface for this creator
type ThingCreator interface {
	Create(c types.Thing) (int64, error)
}

// SQLThingCreator implements a SQL creator
type SQLThingCreator struct {
	db *sqlx.DB
}

// NewSQLThingCreator returns a SQL based ThingCreator
func NewSQLThingCreator(db *sqlx.DB) ThingCreator {
	return &SQLThingCreator{db: db}
}

// Create inserts a Thing record in to the database
func (r *SQLThingCreator) Create(c types.Thing) (int64, error) {
	columns := []string{
		"color",
		"description",
		"length",
		"height",
	}

	return creator.Insert(r.db.DB, "things", columns, c.Color, c.Description, c.Length, c.Height)
}

package example

// GENERATED CODE, EDITS WILL BE LOST
//
// generated from example/Thing
// using http://github.com/brianstarke/go-beget/beget

import (
	"database/sql"

	"github.com/lann/squirrel"
)

var insertColumnsThing = [2]string{
	"color",
	"description",
}

// CreateThing inserts a Thing in to the database
func (c *Thing) CreateThing(db *sql.DB) (int64, error) {
	query := squirrel.Insert("things").
		Columns(insertColumnsThing...).
		Values(c.Color, c.Description).
		Suffix("RETURNING \"id\"").
		RunWith(db).
		PlaceholderFormat(squirrel.Dollar)

	var id int64

	err := query.QueryRow().Scan(&id)

	return id, err
}

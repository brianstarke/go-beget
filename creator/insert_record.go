package creator

import (
	"database/sql"

	"github.com/lann/squirrel"
)

// Insert is helper for generated Create repos to use for inserting records in
// to the datbase.  It returns the ID of the created record.
// TODO: make the ID field configurable
func Insert(db *sql.DB, tableName string, columnNames []string, values ...interface{}) (int64, error) {
	query := squirrel.Insert(tableName).
		Columns(columnNames...).
		Values(values...).
		Suffix("RETURNING \"id\"").
		RunWith(db).
		PlaceholderFormat(squirrel.Dollar)

	var id int64

	err := query.QueryRow().Scan(&id)

	return id, err
}

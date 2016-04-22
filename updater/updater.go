package updater

import (
	"log"

	"github.com/lann/squirrel"
	"github.com/mgutz/ansi"
)

var (
	logPrefix = "[" +
		ansi.Color("go-beget", "154") +
		"/" +
		ansi.Color("searcher", "159") + "] "
)

// UpdateRequest interface needs to be implemented for anything that
// uses updater functions
type UpdateRequest interface {
	GetID() int64                       // the uuid to update
	GetTableName() string               // the table name to generate update SQL for
	GetUpdates() map[string]interface{} // db field name and it's new value
}

// UpdateField does something too
type UpdateField interface {
	String() string
}

// GetUpdateSQL generates a SQL update statement from things that implement
// UpdateRequest
func GetUpdateSQL(ur UpdateRequest) (sqlStr string, values interface{}, err error) {
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	sql := psql.Update(ur.GetTableName())

	for field, value := range ur.GetUpdates() {
		sql = sql.Set(field, value)
	}

	sql = sql.Where(squirrel.Eq{"id": ur.GetID()})

	sqlStr, values, err = sql.ToSql()

	log.SetFlags(0)
	log.SetPrefix(logPrefix)
	log.Printf("UPDATE SQL generated %s", ansi.Color(sqlStr, "155+b"))
	log.SetPrefix("")

	return sqlStr, values, err
}

// GetDeleteSQL generates a delete statement
func GetDeleteSQL(ur UpdateRequest) (sqlStr string, values interface{}, err error) {
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	sql := psql.Delete(ur.GetTableName()).Where("id = ?", ur.GetID())

	return sql.ToSql()
}

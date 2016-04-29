package update

/*
	GENERATED CODE, YOUR EDITS ARE FUTILE

	generated from Thing
	using http://github.com/brianstarke/go-beget/generator/updater
*/

import (
	"fmt"

	"github.com/brianstarke/go-beget/updater"
	"github.com/jmoiron/sqlx"
)

// ThingUpdater is the interface
type ThingUpdater interface {
	Update(updateRequest ThingUpdateRequest) error
	DeleteByID(ID int64) error
}

// SQLThingUpdater implements a SQL based updater
type SQLThingUpdater struct {
	db *sqlx.DB
}

// NewSQLThingUpdater returns a configured updater
func NewSQLThingUpdater(db *sqlx.DB) ThingUpdater {
	return &SQLThingUpdater{db: db}
}

// Update converts a ThingUpdateRequest in to SQL and executes it
func (r *SQLThingUpdater) Update(updateRequest ThingUpdateRequest) error {
	sqlStr, values, err := updater.GetUpdateSQL(&updateRequest)

	if err != nil {
		return fmt.Errorf("Error generating SQL for ThingUpdateRequest : %s", err.Error())
	}

	_, err = r.db.Exec(sqlStr, values.([]interface{})...)

	return err
}

// DeleteByID generates a Delete update and executes it
func (r *SQLThingUpdater) DeleteByID(ID int64) error {
	ur := ThingUpdateRequest{
		ID: ID,
	}

	sqlStr, values, err := updater.GetDeleteSQL(&ur)

	if err != nil {
		return fmt.Errorf("Error generating SQL for Thing DeleteByID : %s", err.Error())
	}

	_, err = r.db.Exec(sqlStr, values.([]interface{})...)

	return err
}

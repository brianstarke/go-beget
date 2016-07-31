package example

// GENERATED CODE, EDITS WILL BE LOST
//
// generated from example/Thing
// using http://github.com/brianstarke/go-beget/beget

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	sq "github.com/lann/squirrel"
)

// ThingUpdateRequest defines a set of parameters for
// updating Thing.  It can be serialized and passed
// between services as JSON, or used to generate a SQL statement.
type ThingUpdateRequest struct {
	ID      int64                      `json:"id"`
	Updates map[ThingField]interface{} `json:"updates"`
}

// NewThingUpdateRequest creates a new update request for Thing
func (u *ThingUpdateRequest) NewThingUpdateRequest(id int64) *ThingUpdateRequest {
	return &ThingUpdateRequest{
		ID: id,
	}
}

// AddUpdate is a convenience method for adding updates to this request.  Chainable.
func (u *ThingUpdateRequest) AddUpdate(field ThingField, value interface{}) *ThingUpdateRequest {
	u.Updates[field] = value

	return u
}

// GenerateUpdateSQL returns the SQL string and placeholder values, or an error
func (u *ThingUpdateRequest) GenerateUpdateSQL() (string, []interface{}, error) {
	sql := sq.
		StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Update("things")

	for field, value := range u.Updates {
		sql = sql.Set(field.DbFieldName(), value)
	}

	sql = sql.Where(sq.Eq{"id": u.ID})

	return sql.ToSql()
}

// ExecuteUpdate will take a sql.DB connection and execute this update
func (u *ThingUpdateRequest) ExecuteUpdate(db *sql.DB) error {
	sql, values, err := u.GenerateUpdateSQL()

	if err != nil {
		return err
	}

	_, err = db.Exec(sql, values)

	return err
}

// NewThingUpdateHandlerFunc returns an HTTP handler func for
// ThingUpdateRequests. It returns 200 and the results
// on success, 404 if not a POST, 400 on bad JSON, 500 on any
// other error.
func NewThingUpdateHandlerFunc(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		b, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "No update request in the request body")
			return
		}

		var ur ThingUpdateRequest
		err = json.Unmarshal(b, &ur)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error deserializing update request: %s", err.Error())
			return
		}

		err = ur.ExecuteUpdate(db)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error executing update request: %s", err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ok")

		return
	}
}

// MarshalText implements https://golang.org/pkg/encoding/#TextMarshaler
func (u ThingUpdateRequest) MarshalText() ([]byte, error) {
	stringified := make(map[string]interface{})

	for key, value := range u.Updates {
		s, err := key.MarshalText()

		if err != nil {
			return nil, err
		}

		stringified[string(s)] = value
	}

	result := map[string]interface{}{
		"updates": stringified,
		"id":      u.ID,
	}

	return json.Marshal(result)
}

// UnmarshalText implements https://golang.org/pkg/encoding/#TextUnmarshaler
func (u *ThingUpdateRequest) UnmarshalText(b []byte) error {
	var i map[string]interface{}

	err := json.Unmarshal(b, &i)

	if err != nil {
		return err
	}

	u.ID = i["id"].(int64)

	var j = make(map[ThingField]interface{})

	for key, value := range i["updates"].(map[string]interface{}) {
		var p ThingField
		err = p.UnmarshalText([]byte(key))

		if err != nil {
			return err
		}
		j[p] = value
	}

	u.Updates = j

	return nil
}

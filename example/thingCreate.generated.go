package example

// GENERATED CODE, EDITS WILL BE LOST
//
// generated from example/Thing
// using http://github.com/brianstarke/go-beget

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lann/squirrel"
)

var insertColumnsThing = []string{
	"color",
	"description",
}

// Create inserts a Thing in to the database
func (c *Thing) Create(db *sql.DB) (int64, error) {
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

// NewThingCreateHandlerFunc returns an HTTP handler func for
// Thing record creations. It returns 200 and the results
// on success, 400 on bad JSON, 500 on any other error.
func NewThingCreateHandlerFunc(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "No Thing in the request body")
			return
		}

		var t Thing
		err = json.Unmarshal(b, &t)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error deserializing Thing: %s", err.Error())
			return
		}

		id, err := t.Create(db)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error executing create request: %s", err.Error())
			return
		}

		jsonResult, err := json.Marshal(struct {
			ID int64 `json:"id"`
		}{
			ID: id,
		})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error serializing create results: %s", err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, string(jsonResult))

		return
	}
}

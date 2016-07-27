package search

// GENERATED CODE, EDITS WILL BE LOST
//
// generated from Thing
// using http://github.com/brianstarke/go-beget/generator/searcher

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/jmoiron/sqlx"

	searchType "github.com/brianstarke/go-beget/example/types"
	sq "github.com/lann/squirrel"
)

// ThingSearchRequest is a serializable SearchRequest for Thing.
type ThingSearchRequest struct {
	Fields  []ThingField  `json:"fields"`
	Filters []ThingFilter `json:"filters"`
	OrderBy *ThingOrderBy `json:"orderBy"`
	Limit   int           `json:"limit"`
	Offset  int           `json:"offset"`
}

// ThingField is a field within the Thing struct
// that is able to be filtered on, sorted on, or returned.
type ThingField int

// Faux enum'd for helpfulness
const (
	ThingID ThingField = iota
	ThingColor
	ThingDescription
	ThingLength
	ThingHeight
)

// JSON name constants
const (
	cID          string = "id"
	cColor       string = "color"
	cDescription string = "description"
	cLength      string = "length"
	cHeight      string = "height"
)

// DbFieldName returns the name of the field to use in the SQL query
func (s ThingField) DbFieldName() string {
	switch s {
	case ThingID:
		return "id"
	case ThingColor:
		return "color"
	case ThingDescription:
		return "description"
	case ThingLength:
		return "length"
	case ThingHeight:
		return "height"

	}
	return ""
}

// MarshalText implements https://golang.org/pkg/encoding/#TextMarshaler
func (s ThingField) MarshalText() ([]byte, error) {
	var data string

	switch s {
	case ThingID:
		data = cID
	case ThingColor:
		data = cColor
	case ThingDescription:
		data = cDescription
	case ThingLength:
		data = cLength
	case ThingHeight:
		data = cHeight

	default:
		return nil, fmt.Errorf("Unable to marshal `%v` in to bytes", s)
	}
	return []byte(data), nil
}

// UnmarshalText implements https://golang.org/pkg/encoding/#TextUnmarshaler
func (s *ThingField) UnmarshalText(b []byte) error {
	str := strings.Trim(string(b), `"`)

	switch str {
	case cID:
		*s = ThingID
	case cColor:
		*s = ThingColor
	case cDescription:
		*s = ThingDescription
	case cLength:
		*s = ThingLength
	case cHeight:
		*s = ThingHeight

	default:
		return fmt.Errorf("Unable to marshal '%s' into type ThingField", str)
	}
	return nil
}

// ThingFilter is a filter specific to Thing
type ThingFilter struct {
	Field     ThingField  `json:"field"`
	Value     interface{} `json:"value"`
	Operator  Operator    `json:"operator"`
	Condition Condition   `json:"condition"`
}

// ThingOrderBy is a sort directive that is specific to Thing
type ThingOrderBy struct {
	Field        ThingField `json:"field"`
	IsDescending bool       `json:"isDescending"`
}

// AddFilter is a helper method to add a filter to a SearchRequest.  By default,
// the operator is EQ and the condition is AND.  If you want to override that:
//
//  var sr SearchRequest
//  sr.AddFilter("color", "red", func(f *ThingFilter){
//    f.Operator = GT
//    f.Condition = OR
//  })
//
// and you're all set.
func (sr *ThingSearchRequest) AddFilter(field ThingField, value interface{}, cfg ...func(f *ThingFilter)) *ThingSearchRequest {
	f := ThingFilter{
		Field:     field,
		Operator:  Eq,
		Condition: And,
		Value:     fmt.Sprint(value),
	}

	for _, c := range cfg {
		c(&f)
	}

	sr.Filters = append(sr.Filters, f)

	return sr
}

// AddFields adds fields to the select.  Helper method for chain building a
// ThingSearchRequest.
func (sr *ThingSearchRequest) AddFields(fields ...ThingField) *ThingSearchRequest {
	sr.Fields = append(sr.Fields, fields...)

	return sr
}

// SetLimit sets the Limit.  Helper method for chain building a ThingSearchRequest
func (sr *ThingSearchRequest) SetLimit(limit int) *ThingSearchRequest {
	sr.Limit = limit

	return sr
}

// SetOffset sets the Offset.  Helper method for chain building a ThingSearchRequest
func (sr *ThingSearchRequest) SetOffset(offset int) *ThingSearchRequest {
	sr.Offset = offset

	return sr
}

// SetPage sets the Offset and Limit based on the page number and page size.
func (sr *ThingSearchRequest) SetPage(pageNumber, pageSize int) *ThingSearchRequest {
	return sr.
		SetOffset(pageNumber * pageSize).
		SetLimit(pageSize)
}

// SetOrderBy sets the Order By.  Helper method for chain building a SearchRequest
func (sr *ThingSearchRequest) SetOrderBy(field ThingField, isDescending bool) *ThingSearchRequest {
	sr.OrderBy = &ThingOrderBy{
		Field:        field,
		IsDescending: isDescending,
	}

	return sr
}

// GenerateSelectSQL will generate an executable SQL statement and return the SQL
// string (with placeholders) and a slice of the values.
func (sr *ThingSearchRequest) GenerateSelectSQL() (string, []interface{}, error) {
	var s string

	if len(sr.Fields) == 0 {
		s = "*"
	} else {
		for _, f := range sr.Fields {
			s = s + f.DbFieldName() + ", "
		}
		s = strings.TrimRight(s, ", ")
	}

	psql := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select(s).
		From("things")

	sr.addPredicates(&psql)

	if sr.Limit > 0 {
		psql = psql.Limit(uint64(sr.Limit))
	}

	if sr.Offset > 0 {
		psql = psql.Offset(uint64(sr.Offset))
	}

	if sr.OrderBy != nil {
		var oBy string

		if sr.OrderBy.IsDescending {
			oBy = fmt.Sprintf("%s DESC", sr.OrderBy.Field.DbFieldName())
		} else {
			oBy = sr.OrderBy.Field.DbFieldName()
		}

		psql = psql.OrderBy(oBy)
	}

	return psql.ToSql()
}

// GenerateCountSQL will generate an executable SQL statement and return the SQL
// string (with placeholders) and a slice of the values.
func (sr *ThingSearchRequest) GenerateCountSQL() (string, []interface{}, error) {
	psql := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("COUNT(*) AS cnt").
		From("things")

	sr.addPredicates(&psql)

	return psql.ToSql()
}

func (sr *ThingSearchRequest) addPredicates(psql *sq.SelectBuilder) {
	var ands sq.And
	var ors sq.Or

	for _, f := range sr.Filters {
		e := sq.Expr(sr.buildExpr(f), f.Value)

		if f.Operator == IsNull || f.Operator == NotNull {
			e = sq.Expr(sr.buildExpr(f))
		}

		if f.Condition == Or {
			ors = append(ors, e)
		} else {
			ands = append(ands, e)
		}
	}

	if len(ors) > 0 {
		*psql = psql.Where(ors)
	}

	if len(ands) > 0 {
		*psql = psql.Where(ands)
	}
}

func (sr ThingSearchRequest) buildExpr(f ThingFilter) string {
	var operator string

	switch f.Operator {
	case Eq:
		operator = "="
	case NotEq:
		operator = "!="
	case Like:
		return fmt.Sprintf("cast (%s as varchar(64)) ILIKE ?", f.Field)
	case NotLike:
		return fmt.Sprintf("cast (%s as varchar(64)) NOT ILIKE ?", f.Field)
	case GreaterThan:
		operator = ">"
	case GreaterThanOrEq:
		operator = ">="
	case LesserThan:
		operator = "<"
	case LesserThanOrEq:
		operator = "<="
	case IsNull:
		return fmt.Sprintf("%s IS NULL", f.Field)
	case NotNull:
		return fmt.Sprintf("%s IS NOT NULL", f.Field)
	default:
		operator = "="
	}

	return fmt.Sprintf("%s %s ?", f.Field.DbFieldName(), operator)
}

// ExecuteSearch will take a sql.DB connection and execute this search request
func (sr *ThingSearchRequest) ExecuteSearch(db *sql.DB, results *[]searchType.Thing) error {
	// Generate the SQL
	sqlStr, values, err := sr.GenerateSelectSQL()

	if err != nil {
		return err
	}

	// Upgrade to sqlx connection
	dbx := sqlx.NewDb(db, "postgres")

	return dbx.Select(results, sqlStr, values...)
}

// ExecuteCount will take a sql.DB connection and execute a SELECT COUNT
func (sr *ThingSearchRequest) ExecuteCount(db *sql.DB) (int32, error) {
	// Generate the SQL
	sqlStr, values, err := sr.GenerateCountSQL()

	if err != nil {
		return 0, err
	}

	// Upgrade to sqlx connection
	dbx := sqlx.NewDb(db, "postgres")

	var results []struct {
		Count int32 `db:"cnt"`
	}

	err = dbx.Select(&results, sqlStr, values...)

	if err != nil {
		return 0, err
	}

	return results[0].Count, nil
}

// NewThingSearchHandlerFunc returns an HTTP handler func for
// ThingSearchRequests. It returns 200 and the results
// on success, 404 if not a POST, 400 on bad JSON, 500 on any
// other error.
func NewThingSearchHandlerFunc(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		b, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "No search request in the request body")
			return
		}

		var sr ThingSearchRequest
		err = json.Unmarshal(b, &sr)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error deserializing search request: %s", err.Error())
			return
		}

		var results []searchType.Thing
		err = sr.ExecuteSearch(db, &results)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error executing search request: %s", err.Error())
			return
		}

		jsonResults, err := json.Marshal(results)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error serializing search results: %s", err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Print(string(jsonResults))

		return
	}
}

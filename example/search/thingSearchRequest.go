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
)

// ThingField is a field within the Thing struct
// that is able to be filtered on, sorted on, or returned.
type ThingField int

// Enum'd for helpfulness
const (
	ThingID ThingField = iota
	ThingColor
	ThingDescription
	ThingLength
	ThingHeight
)

// MarshalText implements https://golang.org/pkg/encoding/#TextMarshaler
func (s ThingField) MarshalText() ([]byte, error) {
	var data string

	switch s {
	case ThingID:
		data = "id"
	case ThingColor:
		data = "color"
	case ThingDescription:
		data = "description"
	case ThingLength:
		data = "length"
	case ThingHeight:
		data = "height"

	default:
		return nil, fmt.Errorf("Unable to marshal `%v` in to bytes", s)
	}
	return []byte(data), nil
}

// UnmarshalText implements https://golang.org/pkg/encoding/#TextUnmarshaler
func (s *ThingField) UnmarshalText(b []byte) error {
	str := strings.Trim(string(b), `"`)

	switch str {
	case "id":
		*s = ThingID
	case "color":
		*s = ThingColor
	case "description":
		*s = ThingDescription
	case "length":
		*s = ThingLength
	case "height":
		*s = ThingHeight

	default:
		return fmt.Errorf("Unable to marshal '%s' into type ThingField", str)
	}
	return nil
}

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

/*
ThingSearchRequest defines a set of parameters for
searching for Thing.  It can be serialized and passed
between services as JSON, or used to generate a SQL statement.
*/
type ThingSearchRequest struct {
	searcher.SearchRequestFields
	Filters []ThingSearchFilter `json:"filters"`
	OrderBy ThingOrderBy        `json:"orderBy"`
	Fields  []ThingField        `json:"fields"`
}

/*
ThingSearchFilter is a filter specific to Thing
*/
type ThingSearchFilter struct {
	Field     ThingField               `json:"field"`
	Value     interface{}              `json:"value"`
	Operator  searcher.FilterOperator  `json:"operator"`
	Condition searcher.FilterCondition `json:"condition"`
}

/*
ThingOrderBy is a sort directive that is specific to Thing
*/
type ThingOrderBy struct {
	Field      ThingField `json:"field"`
	Descending bool       `json:"desc"`
}

// Implement searcher.SearchRequest interface

// GetTableName returns the database table name
func (sr *ThingSearchRequest) GetTableName() string {
	return "things"
}

// GetFilters returns the SQL WHERE clauses
func (sr *ThingSearchRequest) GetFilters() []searcher.Filter {
	filters := []searcher.Filter{}

	for _, f := range sr.Filters {
		filter := searcher.Filter{
			Field:     f.Field,
			Value:     f.Value,
			Operator:  f.Operator,
			Condition: f.Condition,
		}
		filters = append(filters, filter)
	}

	return filters
}

// GetOrderBy returns the SQL ORDER BY clauses
func (sr *ThingSearchRequest) GetOrderBy() searcher.OrderBy {
	return searcher.OrderBy{
		Field:      sr.OrderBy.Field,
		Descending: sr.OrderBy.Descending,
	}
}

// GetLimit returns the SQL LIMIT clause
func (sr *ThingSearchRequest) GetLimit() int {
	return sr.Limit
}

// GetOffset returns the SQL OFFSET clause
func (sr *ThingSearchRequest) GetOffset() int {
	return sr.Offset
}

// AddFilter adds a WHERE clause
func (sr *ThingSearchRequest) AddFilter(field ThingField, value interface{}, operator searcher.FilterOperator, condition searcher.FilterCondition) {
	f := ThingSearchFilter{
		Field:     field,
		Value:     value,
		Operator:  operator,
		Condition: condition,
	}
	sr.Filters = append(sr.Filters, f)
}

// SetOrderBy sets the ORDER BY clause
func (sr *ThingSearchRequest) SetOrderBy(field ThingField, isDescending bool) {
	sr.OrderBy = ThingOrderBy{
		Field:      field,
		Descending: isDescending,
	}
}

// GetFields returns the SQL SELECT fields
func (sr *ThingSearchRequest) GetFields() []string {
	fields := []string{}

	for _, f := range sr.Fields {
		fields = append(fields, f.DbFieldName())
	}

	return fields
}

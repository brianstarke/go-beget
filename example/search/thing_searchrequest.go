package search

/*
	GENERATED CODE, YOUR EDITS ARE FUTILE

	generated from Thing
	using http://github.com/brianstarke/go-begetter/generator/searcher
*/

import (
	"fmt"
	"strings"

	"github.com/brianstarke/go-begetter/searcher"
)

/*
ThingField is a field within the Thing struct
that is able to be filtered on, sorted on, or returned.
*/
type ThingField int

const (
	Thing_Color ThingField = iota
	Thing_Description
	Thing_Length
	Thing_Height
)

// MarshalText implements https://golang.org/pkg/encoding/#TextMarshaler
func (s ThingField) MarshalText() ([]byte, error) {
	var data string

	switch s {
	case Thing_Color:
		data = "Color"
	case Thing_Description:
		data = "Description"
	case Thing_Length:
		data = "Length"
	case Thing_Height:
		data = "Height"

	default:
		return nil, fmt.Errorf("Unable to marshal `%v` in to bytes", s)
	}
	return []byte(data), nil
}

// UnmarshalText implements https://golang.org/pkg/encoding/#TextUnmarshaler
func (s *ThingField) UnmarshalText(b []byte) error {
	str := strings.Trim(string(b), `"`)

	switch str {
	case "Color":
		*s = Thing_Color
	case "Description":
		*s = Thing_Description
	case "Length":
		*s = Thing_Length
	case "Height":
		*s = Thing_Height

	default:
		return fmt.Errorf("Unable to marshal '%s' into type ThingField", str)
	}
	return nil
}

// DbFieldName returns the name of the field to use in the SQL query
func (s ThingField) DbFieldName() string {
	switch s {
	case Thing_Color:
		return "Color"
	case Thing_Description:
		return "Description"
	case Thing_Length:
		return "Length"
	case Thing_Height:
		return "Height"

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

// implement searcher.SearchRequest interface
func (sr *ThingSearchRequest) GetTableName() string {
	return "things"
}

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

func (sr *ThingSearchRequest) GetOrderBy() searcher.OrderBy {
	return searcher.OrderBy{
		Field:      sr.OrderBy.Field,
		Descending: sr.OrderBy.Descending,
	}
}

func (sr *ThingSearchRequest) GetLimit() int {
	return sr.Limit
}

func (sr *ThingSearchRequest) GetOffset() int {
	return sr.Offset
}

func (sr *ThingSearchRequest) AddFilter(field ThingField, value interface{}, operator searcher.FilterOperator, condition searcher.FilterCondition) {
	f := ThingSearchFilter{
		Field:     field,
		Value:     value,
		Operator:  operator,
		Condition: condition,
	}
	sr.Filters = append(sr.Filters, f)
}

func (sr *ThingSearchRequest) SetOrderBy(field ThingField, isDescending bool) {
	sr.OrderBy = ThingOrderBy{
		Field:      field,
		Descending: isDescending,
	}
}

func (sr *ThingSearchRequest) GetFields() []string {
	fields := []string{}

	for _, f := range sr.Fields {
		fields = append(fields, f.DbFieldName())
	}

	return fields
}

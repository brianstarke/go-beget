package searcher

/*
SearchRequest interface for generating SQL queries
*/
type SearchRequest interface {
	GetFields() []string  // the SELECT (fields to return), * if nil
	GetTableName() string // the FROM (table name)
	GetFilters() []Filter // the WHERE clauses
	GetOrderBy() OrderBy  // sorting (if any)
	GetLimit() int        // limit on results (useful for paging)
	GetOffset() int       // with limit, can be used for paging
}

type SearchRequestFields struct {
	Limit  int `json:"limit,omitempty"`  // num to return
	Offset int `json:"offset,omitempty"` // starting offset, used for paging
}

// OrderBy holds the field name and direction to sort on
type OrderBy struct {
	Field      FilterField
	Descending bool // ascending is default, true sets to order descending
}

type Filter struct {
	Field     FilterField     // the field name to apply filter to, must be castable to string
	Value     interface{}     // the value we're filtering by
	Operator  FilterOperator  // the operator to apply to the value (EQ if omitted)
	Condition FilterCondition // the boolean condition of this filter (AND if omitted)
}

type FilterOperator string

const (
	EQ        FilterOperator = "EQ"          // equal
	IEQ                      = "IEQ"         // equals, ignore case
	NE                       = "NE"          // not equal
	NIE                      = "NIE"         //not equal, ignore case
	GT                       = "GT"          // greater than
	GTE                      = "GTE"         // greater then or equal
	LT                       = "LT"          // less than
	LTE                      = "LTE"         // less than or equal
	LIKE                     = "LIKE"        // like (value can have %'s)
	ISNULL                   = "IS NULL"     // Only for NULL Comparisons
	ISNOTNULL                = "IS NOT NULL" // Only for NULL Comparisons
)

type FilterCondition string

const (
	AND FilterCondition = "AND"
	OR                  = "OR"
)

type FilterField interface {
	DbFieldName() string
}

package types

//go:generate searcher -struct=Thing -table=things -repos=sql
//go:generate creator -struct=Thing -table=things -repos=sql

// Thing has characteristics
type Thing struct {
	ID          int64  `beget:"search" json:"id" db:"id"`
	Color       string `beget:"search,create" json:"color" db:"color"`
	Description string `beget:"search,create" json:"description" db:"description"`
	Length      int    `beget:"search,create" json:"length" db:"length"`
	Height      int    `beget:"search,create" json:"height" db:"height"`
}

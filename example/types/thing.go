package types

//go:generate searcher -struct=Thing -table=things

// Thing has characteristics
type Thing struct {
	Color       string `beget:"search" json:"color" db:"color"`
	Description string `beget:"search" json:"description" db:"description"`
	Length      int    `beget:"search" json:"length" db:"length"`
	Height      int    `beget:"search" json:"height" db:"height"`
}

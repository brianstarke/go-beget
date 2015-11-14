package types

//go:generate searcher -struct=Thing -table=things

// Thing has characteristics
type Thing struct {
	Color       string `beget:"search"`
	Description string `beget:"search"`
	Length      int    `beget:"search"`
	Height      int    `beget:"search"`
}

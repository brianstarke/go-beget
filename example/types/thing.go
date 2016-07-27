package types

//go:generate searcher -struct=Thing -table=things
//go:generate creator -struct=Thing -table=things -impls=sql,gin
//go:generate updater -struct=Thing -table=things -impls=sql,gin

// Thing has characteristics
type Thing struct {
	ID          int64  `beget:"search" json:"id" db:"id"`
	Color       string `beget:"search,create,update" json:"color" db:"color"`
	Description string `beget:"search,create,update" json:"description" db:"description"`
	Length      int    `beget:"search,create,update" json:"length" db:"length"`
	Height      int    `beget:"search,create,update" json:"height" db:"height"`
}

package example

//go:generate searcher -struct=Thing -table=things

// Thing is a normal thing
type Thing struct {
	ID          int64  `json:"id" db:"id"`
	Color       string `json:"color" db:"color"`
	Description string `json:"description" db:"description"`
	Length      int    `json:"length" db:"length"`
	Height      int    `json:"height" db:"height"`
}

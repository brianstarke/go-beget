package example

import "time"

//go:generate go-beget -struct=Thing -table=things -omitFromInsert=Length,Height

// Thing is a normal thing
type Thing struct {
	ID          int64     `json:"id" db:"id"`
	Color       string    `json:"color" db:"color"`
	Description string    `json:"description" db:"description"`
	Length      int       `json:"length" db:"length"`
	Height      int       `json:"height" db:"height"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at_utc"`
}

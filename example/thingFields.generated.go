package example

// GENERATED CODE, EDITS WILL BE LOST
//
// generated from example/Thing
// using http://github.com/brianstarke/go-beget/beget

import (
	"fmt"
	"strings"
)

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
	ThingCreatedAt
)

// JSON name constants
const (
	cThingID          string = "id"
	cThingColor       string = "color"
	cThingDescription string = "description"
	cThingLength      string = "length"
	cThingHeight      string = "height"
	cThingCreatedAt   string = "createdAt"
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
	case ThingCreatedAt:
		return "created_at_utc"

	}
	return ""
}

// MarshalText implements https://golang.org/pkg/encoding/#TextMarshaler
func (s ThingField) MarshalText() ([]byte, error) {
	var data string

	switch s {
	case ThingID:
		data = cThingID
	case ThingColor:
		data = cThingColor
	case ThingDescription:
		data = cThingDescription
	case ThingLength:
		data = cThingLength
	case ThingHeight:
		data = cThingHeight
	case ThingCreatedAt:
		data = cThingCreatedAt

	default:
		return nil, fmt.Errorf("Unable to marshal `%v` in to bytes", s)
	}
	return []byte(data), nil
}

// UnmarshalText implements https://golang.org/pkg/encoding/#TextUnmarshaler
func (s *ThingField) UnmarshalText(b []byte) error {
	str := strings.Trim(string(b), `"`)

	switch str {
	case cThingID:
		*s = ThingID
	case cThingColor:
		*s = ThingColor
	case cThingDescription:
		*s = ThingDescription
	case cThingLength:
		*s = ThingLength
	case cThingHeight:
		*s = ThingHeight
	case cThingCreatedAt:
		*s = ThingCreatedAt

	default:
		return fmt.Errorf("Unable to marshal '%s' into type ThingField", str)
	}
	return nil
}

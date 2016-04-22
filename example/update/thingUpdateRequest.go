package update

/*
	GENERATED CODE, YOUR EDITS ARE FUTILE

	generated from Thing
	using http://github.com/brianstarke/go-beget/generator/updater
*/

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ThingField is a field within the Thing struct
// that is able to be updated
type ThingField int

// Enum'd for helpfulness
const (
	ThingColor ThingField = iota
	ThingDescription
	ThingLength
	ThingHeight
)

// MarshalText implements https://golang.org/pkg/encoding/#TextMarshaler
func (s ThingField) MarshalText() ([]byte, error) {
	var data string

	switch s {
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

// ThingUpdateRequest defines a set of parameters for
// updating Thing.  It can be serialized and passed
// between services as JSON, or used to generate a SQL statement.
type ThingUpdateRequest struct {
	ID      int64                      `json:"id"`
	Updates map[ThingField]interface{} `json:"updates"`
}

// MarshalText implements https://golang.org/pkg/encoding/#TextMarshaler
func (ur ThingUpdateRequest) MarshalText() ([]byte, error) {
	stringified := make(map[string]interface{})

	for key, value := range ur.Updates {
		s, err := key.MarshalText()

		if err != nil {
			return nil, err
		}

		stringified[string(s)] = value
	}

	result := map[string]interface{}{
		"updates": stringified,
		"id":      ur.ID,
	}

	return json.Marshal(result)
}

// MarshalBinary implements https://golang.org/pkg/encoding/#BinaryMarshaler
func (ur ThingUpdateRequest) MarshalBinary() ([]byte, error) {
	stringified := make(map[string]interface{})

	for key, value := range ur.Updates {
		s, err := key.MarshalText()

		if err != nil {
			return nil, err
		}

		stringified[string(s)] = value
	}

	result := map[string]interface{}{
		"updates": stringified,
		"id":      ur.ID,
	}

	return json.Marshal(result)
}

// UnmarshalText implements https://golang.org/pkg/encoding/#TextUnmarshaler
func (ur *ThingUpdateRequest) UnmarshalText(b []byte) error {
	var i map[string]interface{}

	err := json.Unmarshal(b, &i)

	if err != nil {
		return err
	}

	ur.ID = i["id"].(int64)

	var j = make(map[ThingField]interface{})

	for key, value := range i["updates"].(map[string]interface{}) {
		var p ThingField
		err = p.UnmarshalText([]byte(key))

		if err != nil {
			return err
		}
		j[p] = value
	}

	ur.Updates = j

	return nil
}

// UnmarshalBinary implements https://golang.org/pkg/encoding/#BinaryUnmarshaler
func (ur *ThingUpdateRequest) UnmarshalBinary(b []byte) error {
	var i map[string]interface{}

	err := json.Unmarshal(b, &i)

	if err != nil {
		return err
	}

	ur.ID = i["id"].(int64)

	var j = make(map[ThingField]interface{})

	for key, value := range i["updates"].(map[string]interface{}) {
		var p ThingField
		err = p.UnmarshalText([]byte(key))

		if err != nil {
			return err
		}
		j[p] = value
	}

	ur.Updates = j

	return nil
}

// GetID implements updater.UpdateRequest interface
func (ur *ThingUpdateRequest) GetID() int64 {
	return ur.ID
}

// GetTableName implements updater.UpdateRequest interface
func (ur *ThingUpdateRequest) GetTableName() string {
	return "things"
}

// GetUpdates implements updater.UpdateRequest interface
func (ur *ThingUpdateRequest) GetUpdates() map[string]interface{} {
	updates := make(map[string]interface{})

	for field, value := range ur.Updates {
		updates[field.DbFieldName()] = value
	}

	return updates
}

// AddUpdate implements updater.UpdateRequest interface
func (ur *ThingUpdateRequest) AddUpdate(field ThingField, value interface{}) {
	if ur.Updates == nil {
		ur.Updates = make(map[ThingField]interface{})
	}

	ur.Updates[field] = value
}

package checker

import (
	"encoding/json"
	"io"
)

// JSON checks if a file is a json blob
type JSON struct {
}

// Check if a file is json
func (j JSON) Check(r io.Reader) (bool, error) {
	var data interface{}

	d := json.NewDecoder(r)

	err := d.Decode(&data)

	// Assume an error in parsing is due to the file not being json
	if err != nil {
		return false, nil
	}
	return true, nil
}

// Name returns the name of the parser (json)
func (j JSON) Name() string {
	return "json"
}

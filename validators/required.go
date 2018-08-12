package validators

import (
	"errors"
)

var (
	errRequired = errors.New("this field is required")
)

// Required ensures that a boolean value is set to true.
type Required struct{}

// Validate ensures the value is true.
func (r Required) Validate(b bool) error {
	if !b {
		return errRequired
	}
	return nil
}

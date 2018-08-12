package validators

import (
	"errors"
)

var (
	errNonEmptyString = errors.New("value cannot be empty")
)

// NonEmpty ensures that the provided string value is not empty.
type NonEmpty struct{}

// Validate ensures the string is valid.
func (n NonEmpty) Validate(v string) error {
	if len(v) == 0 {
		return errNonEmptyString
	}
	return nil
}

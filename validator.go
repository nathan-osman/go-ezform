package ezform

import (
	"errors"
	"fmt"
)

var (
	errNonEmptyString = errors.New("value cannot be empty")
	errRequired       = errors.New("this field is required")
)

// Validator defines an interface for writing field validators.
type Validator interface {
	Validate(interface{}) error
}

// NonEmptyValidator ensures that a string value has a length greater than one.
type NonEmptyValidator struct{}

// Validate ensures the string is valid.
func (n NonEmptyValidator) Validate(v string) error {
	if len(v) == 0 {
		return errNonEmptyString
	}
	return nil
}

// MinMaxValidator ensures that an integer falls within the specified range.
type MinMaxValidator struct {
	// Min is the lowest value that will be accepted.
	Min int64
	// Max is the highest value that will be accepted.
	Max int64
}

// Validate ensures the integer is valid.
func (m MinMaxValidator) Validate(v int64) error {
	if v < m.Min {
		return fmt.Errorf("value cannot be less than %d", m.Min)
	}
	if v > m.Max {
		return fmt.Errorf("value cannot be greater than %d", m.Max)
	}
	return nil
}

// RequiredValidator ensures a boolean value of true is provided.
type RequiredValidator struct{}

// Validate ensures the value is true.
func (r RequiredValidator) Validate(b bool) error {
	if !b {
		return errRequired
	}
	return nil
}

package ezform

import (
	"errors"
)

var (
	// ErrNonEmptyString indicates that the provided string was empty.
	ErrNonEmptyString = errors.New("value cannot be empty")
)

// StringValidator defines an interface for string validators.
type StringValidator interface {
	Validate(string) error
}

// String is a field that stores a string value.
type String struct {
	Value      string
	Validators []StringValidator
}

// Validate ensures that the provided value is valid.
func (s *String) Validate(value string) error {
	for _, v := range s.Validators {
		if err := v.Validate(value); err != nil {
			return err
		}
	}
	s.Value = value
	return nil
}

// NonEmptyValidator ensures that a string is not empty.
type NonEmptyValidator struct{}

// Validate ensures the string is not empty.
func (n NonEmptyValidator) Validate(value string) error {
	if len(value) == 0 {
		return ErrNonEmptyString
	}
	return nil
}

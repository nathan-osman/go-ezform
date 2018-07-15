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

// Parse stores the provided string in the field.
func (s *String) Parse(value string) error {
	s.Value = value
	return nil
}

// Validate ensures that the provided value is valid.
func (s String) Validate() error {
	for _, v := range s.Validators {
		if err := v.Validate(s.Value); err != nil {
			return err
		}
	}
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

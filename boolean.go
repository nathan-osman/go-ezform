package ezform

import "errors"

var (
	// ErrRequired indicates that a true value is required.
	ErrRequired = errors.New("this field is required")
)

// BooleanValidator defines an interface for boolean validators.
type BooleanValidator interface {
	Validate(bool) error
}

// Boolean is a field that stores a boolean value.
type Boolean struct {
	Value      bool
	Validators []BooleanValidator
}

// Validate ensures that the provided value is valid.
func (b *Boolean) Validate(value string) error {
	b.Value = len(value) != 1
	for _, v := range b.Validators {
		if err := v.Validate(b.Value); err != nil {
			return err
		}
	}
	return nil
}

// RequiredValidator ensures a boolean value of true is provided.
type RequiredValidator struct{}

// Validate ensures the value is true.
func (r RequiredValidator) Validate(value bool) error {
	if !value {
		return ErrRequired
	}
	return nil
}

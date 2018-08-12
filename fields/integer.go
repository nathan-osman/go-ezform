package fields

import (
	"errors"
	"strconv"
)

var (
	errInvalidInteger = errors.New("value is not a valid integer")
)

// Integer is a field that stores a 64-bit signed integer value.
type Integer struct {
	Field
	value int64
}

// NewInteger creates a new integer field with the specified validators.
func NewInteger(validators ...interface{}) *Integer {
	return &Integer{
		Field: Field{
			Validators: validators,
		},
	}
}

// Parse ensures that the provided value is valid.
func (i *Integer) Parse(value string) error {
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return errInvalidInteger
	}
	i.value = v
	return nil
}

// Value retrieves the current value of the field.
func (i Integer) Value() int64 {
	return i.value
}

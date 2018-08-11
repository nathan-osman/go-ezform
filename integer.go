package ezform

import (
	"errors"
	"strconv"
)

var (
	errInvalidInteger = errors.New("value is not a valid integer")
)

// IntegerField is a field that stores a 64-bit signed integer value.
type IntegerField struct {
	Field
	value int64
}

// NewIntegerField creates a new integer field with the specified validators.
func NewIntegerField(validators ...interface{}) *IntegerField {
	return &IntegerField{
		Field: Field{
			Validators: validators,
		},
	}
}

// Parse ensures that the provided value is valid.
func (i *IntegerField) Parse(value string) error {
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return errInvalidInteger
	}
	i.value = v
	return nil
}

// Value retrieves the current value of the field.
func (i IntegerField) Value() int64 {
	return i.value
}

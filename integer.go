package ezform

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
func (i Integer) Value() interface{} {
	return i.value
}

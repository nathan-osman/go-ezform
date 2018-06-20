package ezform

import (
	"errors"
	"reflect"
)

var (
	// ErrValueIsZero indicates that a non-zero value must be supplied for
	// the field.
	ErrValueIsZero = errors.New("field cannot be empty")
)

// IsNonZero verifies that a non-zero value was supplied.
func IsNonZero(v interface{}) error {
	vVal := reflect.ValueOf(v)
	if vVal.Interface() == reflect.Zero(vVal.Type()).Interface() {
		return ErrValueIsZero
	}
	return nil
}

package ezform

import (
	"errors"
	"reflect"
)

var (
	// ErrFieldRequired indicates that a non-zero value must be supplied for
	// the field.
	ErrFieldRequired = errors.New("field is required")
)

// Required verifies that a non-zero value was supplied.
func Required(v interface{}) error {
	vVal := reflect.ValueOf(v)
	if vVal.Interface() == reflect.Zero(vVal.Type()).Interface() {
		return ErrFieldRequired
	}
	return nil
}

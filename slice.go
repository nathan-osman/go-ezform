package ezform

import (
	"errors"
	"reflect"
)

var (
	// ErrInvalidValue indicates that an invalid value was supplied for the
	// field.
	ErrInvalidValue = errors.New("an invalid value was supplied")
)

// InSlice determines if the specified value is contained in the slice.
func InSlice(slice interface{}, v interface{}) error {
	var (
		sType = reflect.TypeOf(slice)
		vType = reflect.TypeOf(v)
	)
	if sType.Kind() != reflect.Slice {
		return ErrInvalidValue
	}
	if sType.Elem() != vType {
		return ErrInvalidValue
	}
	var (
		sVal = reflect.ValueOf(slice)
		vVal = reflect.ValueOf(v)
	)
	for i := 0; i < sVal.Len(); i++ {
		if sVal.Index(i).Interface() == vVal.Interface() {
			return nil
		}
	}
	return ErrInvalidValue
}

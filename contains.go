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

// Contains determines if the specified value is contained in a slice or map
// key.
func Contains(c interface{}, v interface{}) error {
	var (
		cType = reflect.TypeOf(c)
		vType = reflect.TypeOf(v)
	)
	switch cType.Kind() {
	case reflect.Slice:
		if cType.Elem() != vType {
			return ErrInvalidValue
		}
	case reflect.Map:
		if cType.Key() != vType {
			return ErrInvalidValue
		}
	default:
		return ErrInvalidValue
	}
	var (
		cVal = reflect.ValueOf(c)
		vVal = reflect.ValueOf(v)
	)
	switch cType.Kind() {
	case reflect.Slice:
		for i := 0; i < cVal.Len(); i++ {
			if cVal.Index(i).Interface() == vVal.Interface() {
				return nil
			}
		}
	case reflect.Map:
		for _, k := range cVal.MapKeys() {
			if k.Interface() == vVal.Interface() {
				return nil
			}
		}
	}
	return ErrInvalidValue
}

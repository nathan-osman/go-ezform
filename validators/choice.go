package validators

import (
	"errors"
	"reflect"
)

var (
	errInvalidType  = errors.New("invalid type provided")
	errInvalidValue = errors.New("invalid value provided")
)

// Choice ensures that a provided value is contained within a list of acceptable values.
// The Choices field should be set to either a slice of valid values or a map whose keys will then be used as a list of valid values.
type Choice struct {
	Choices interface{}
}

// Validate ensures the value is one of the accepted values.
func (c Choice) Validate(v interface{}) error {
	var (
		cVal = reflect.ValueOf(c.Choices)
		vVal = reflect.ValueOf(v)
	)
	switch cVal.Kind() {
	case reflect.Slice:
		if cVal.Type().Elem() != vVal.Type() {
			return errInvalidType
		}
		for i := 0; i < cVal.Len(); i++ {
			if cVal.Index(i).Interface() == vVal.Interface() {
				return nil
			}
		}
	case reflect.Map:
		if cVal.Type().Key() != vVal.Type() {
			return errInvalidType
		}
		for _, k := range cVal.MapKeys() {
			if k.Interface() == vVal.Interface() {
				return nil
			}
		}
	default:
		return errInvalidType
	}
	return errInvalidValue
}

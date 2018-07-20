package ezform

import (
	"errors"
	"reflect"
)

var (
	errorInterface = reflect.TypeOf((*error)(nil)).Elem()

	errValidatorStruct = errors.New("validator must be a struct or pointer")
	errValidatorMethod = errors.New("validator must have a Validate method")
	errValidatorParams = errors.New("validator must accept a single parameter")
	errValidatorReturn = errors.New("validator must return an error")
)

// Parser defines an interface that all fields must implement for parsing field values as strings and converting them to the appropriate type.
type Parser interface {
	Parse(string) error
	Value() interface{}
}

// Field represents a single form field.
// All field types should include this type as an anonymous member.
type Field struct {
	Validators []interface{}
	Error      error
}

// Validate uses each validator specified for the field to ensure that the field contains a valid value.
func (f Field) Validate(value interface{}) error {
	var (
		valueType = reflect.TypeOf(value)
		valueVal  = reflect.ValueOf(value)
	)
	for _, v := range f.Validators {
		var (
			vType = reflect.TypeOf(v)
			vVal  = reflect.ValueOf(v)
		)
		if vType.Kind() == reflect.Ptr {
			vType = vType.Elem()
			vVal = vVal.Elem()
		}
		if vType.Kind() != reflect.Struct {
			return errValidatorStruct
		}
		m, found := vType.MethodByName("Validate")
		if !found {
			return errValidatorMethod
		}
		if m.Type.NumIn() != 2 || m.Type.In(1) != valueType {
			return errValidatorParams
		}
		if m.Type.NumOut() != 1 || m.Type.Out(0) != errorInterface {
			return errValidatorReturn
		}
		rVal := vVal.MethodByName("Validate").Call([]reflect.Value{
			valueVal,
		})[0]
		if !rVal.IsNil() {
			return rVal.Interface().(error)
		}
	}
}

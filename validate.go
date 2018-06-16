package ezform

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	errorInterface = reflect.TypeOf((*error)(nil)).Elem()

	errInvalidParameter  = errors.New("validation method must accept a single parameter of the field's type")
	errInvalidReturnType = errors.New("validation method must return an error")
)

// validate checks for a Validate* method and invokes it if present.
func validate(vType reflect.Type, sType reflect.StructField, vVal, fVal reflect.Value) error {
	mName := fmt.Sprintf("Validate%s", sType.Name)
	m, found := vType.MethodByName(mName)
	if !found {
		return nil
	}
	if m.Type.NumIn() != 2 || m.Type.In(1) != sType.Type {
		return errInvalidParameter
	}
	if m.Type.NumOut() != 1 || m.Type.Out(0) != errorInterface {
		return errInvalidReturnType
	}
	rVal := vVal.MethodByName(mName).Call([]reflect.Value{fVal})[0]
	if rVal.IsNil() {
		return nil
	}
	return rVal.Interface().(error)
}

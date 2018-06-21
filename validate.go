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
func validate(vType reflect.Type, sType reflect.StructField, vVal, fVal reflect.Value, param interface{}) error {
	mName := fmt.Sprintf("Validate%s", sType.Name)
	m, found := vType.MethodByName(mName)
	if !found {
		return nil
	}
	hasParam := m.Type.NumIn() == 3
	if (!hasParam && m.Type.NumIn() != 2) || m.Type.In(1) != sType.Type {
		return errInvalidParameter
	}
	if m.Type.NumOut() != 1 || m.Type.Out(0) != errorInterface {
		return errInvalidReturnType
	}
	mParams := []reflect.Value{fVal}
	if hasParam {
		mParams = append(mParams, reflect.ValueOf(param))
	}
	rVal := vVal.MethodByName(mName).Call(mParams)[0]
	if rVal.IsNil() {
		return nil
	}
	return rVal.Interface().(error)
}

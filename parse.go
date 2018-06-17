package ezform

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
)

var (
	errPtrRequired = errors.New("Parse() requires a pointer to struct")
)

// Parse reads form values for a request, validates them, and stores their
// value in the provided struct. The first return value maps field names to
// any errors that occurred during field validation. The second return value
// represents any global errors that occurred during validation. If both
// return values are nil, the form was successfully validated.
func Parse(r *http.Request, v interface{}) (map[string]error, error) {
	if err := r.ParseForm(); err != nil {
		return nil, err
	}
	vType := reflect.TypeOf(v)
	if vType.Kind() != reflect.Ptr {
		return nil, errPtrRequired
	}
	vType = vType.Elem()
	if vType.Kind() != reflect.Struct {
		return nil, errPtrRequired
	}
	var (
		vVal        = reflect.ValueOf(v).Elem()
		fieldErrors = map[string]error{}
	)
	for i := 0; i < vType.NumField(); i++ {
		var (
			fStruct = vType.Field(i)
			fVal    = vVal.Field(i)
			fStrVal = r.Form.Get(fStruct.Name)
		)
		if err := validate(vType, fStruct, vVal, fVal); err != nil {
			fieldErrors[fStruct.Name] = err
			continue
		}
		fmt.Sscanf(fStrVal, "%v", fVal.Addr().Interface())
	}
	return fieldErrors, nil
}

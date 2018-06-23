package ezform

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"
)

var (
	errPtrRequired = errors.New("Parse() requires a pointer to struct")
)

// Parse reads form values for a request, validates them, and stores their value in the provided struct.
// If a value is provided for the param parameter, it will be passed to the validation methods.
// The first return value maps field names to any errors that occurred during field validation.
// The second return value represents any global errors that occurred during validation.
// If both return values are nil, the form was successfully validated.
func Parse(r *http.Request, v interface{}, param interface{}) (map[string]error, error) {
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
		switch fVal.Kind() {
		case reflect.String:
			fVal.SetString(fStrVal)
		case reflect.Bool:
			fVal.SetBool(len(fStrVal) != 0)
		case reflect.Struct:
			if fVal.Type() == reflect.TypeOf(time.Time{}) {
				t, _ := time.Parse(time.RFC3339, fStrVal)
				fVal.Set(reflect.ValueOf(t))
			}
		default:
			fmt.Sscanf(fStrVal, "%v", fVal.Addr().Interface())
		}
		if err := validate(vType, fStruct, vVal, fVal, param); err != nil {
			fieldErrors[fStruct.Name] = err
			continue
		}
	}
	return fieldErrors, nil
}

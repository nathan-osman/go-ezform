package ezform

import (
	"errors"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

var (
	// ErrFieldRequired indicates that a required field is missing.
	ErrFieldRequired = errors.New("field is required")

	// ErrInvalidFieldType indicates that the field type is unsupported.
	ErrInvalidFieldType = errors.New("invalid field type")
)

// Parse reads form values from a request, validates them, and stores their
// value in the provided struct. The return value is a map that lists
// validation errors by field name.
func Parse(r *http.Request, v interface{}) (map[string][]error, error) {
	if err := r.ParseForm(); err != nil {
		return nil, err
	}
	var (
		vType  = reflect.TypeOf(v).Elem()
		vVal   = reflect.ValueOf(v).Elem()
		errors = map[string][]error{}
	)
	for i := 0; i < vType.NumField(); i++ {
		var (
			fType = vType.Field(i)
			fVal  = vVal.Field(i)
			s     = r.Form.Get(vType.Field(i).Name)
		)
		for _, tag := range strings.Split(fType.Tag.Get("form"), ",") {
			switch tag {
			case "required":
				if len(s) == 0 {
					errors[fType.Name] = append(
						errors[fType.Name],
						ErrFieldRequired,
					)
				}
			}
		}
		switch fVal.Kind() {
		case reflect.String:
			fVal.SetString(s)
		case reflect.Int64:
			iVal, _ := strconv.ParseInt(s, 10, 64)
			fVal.SetInt(iVal)
		case reflect.Bool:
			fVal.SetBool(len(s) > 0)
		default:
			return nil, ErrInvalidFieldType
		}
	}
	return errors, nil
}

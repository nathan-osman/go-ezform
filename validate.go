package ezform

import (
	"errors"
	"net/http"
	"reflect"

	"github.com/nathan-osman/go-reflectr"
)

var (
	errInvalidForm   = errors.New("form must be a pointer to struct")
	errInvalidReturn = errors.New("Validate() must return a single value")
)

func returnToError(v interface{}) error {
	if v == nil {
		return nil
	}
	return v.(error)
}

func parseField(s *reflectr.StructMeta, formValue string) error {
	r, err := s.
		Method("Parse").
		Returns(reflectr.ErrorType).
		Call(formValue)
	if err != nil {
		return err
	}
	return returnToError(r[0])
}

func validateField(s *reflectr.StructMeta, v interface{}) (error, error) {
	r, err := s.
		Method("Validate").
		Returns(reflectr.ErrorType).
		Call(v)
	if err != nil {
		return nil, err
	}
	return returnToError(r[0]), nil
}

// Validate parses request data and validates it against the provided form.
func Validate(r *http.Request, v interface{}) (bool, error) {
	vType := reflect.TypeOf(v)
	if vType.Kind() != reflect.Ptr {
		return false, errInvalidForm
	}
	vType = vType.Elem()
	if vType.Kind() != reflect.Struct {
		return false, errInvalidForm
	}
	var (
		vValue    = reflect.ValueOf(v).Elem()
		validated = true
	)
	for i := 0; i < vType.NumField(); i++ {
		var (
			fValue = vValue.Field(i).Addr()
			s      = reflectr.Struct(fValue.Interface())
		)
		if err := parseField(s, r.Form.Get(vType.Field(i).Name)); err != nil {
			return false, err
		}
		r, err := s.Method("Value").Call()
		if err != nil {
			return false, err
		}
		if len(r) != 1 {
			return false, errInvalidReturn
		}
		e, err := validateField(s, r[0])
		if err != nil {
			return false, err
		}
		if e != nil {
			if _, err := s.Method("SetError").Call(e); err != nil {
				return false, err
			}
			validated = false
		}
	}
	return validated, nil
}

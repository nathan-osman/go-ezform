package ezform

import (
	"errors"
	"net/http"

	"github.com/nathan-osman/go-reflectr"
)

var (
	errInvalidForm = errors.New("form must be a pointer to struct")

	errInvalidField    = errors.New("field must be a pointer to struct")
	errParseFieldValue = errors.New("unable to parse form value")
	errGetFieldValue   = errors.New("unable to get field value")
	errRunValidator    = errors.New("unable to run validator")
)

func returnToError(v interface{}) error {
	if v == nil {
		return nil
	}
	return v.(error)
}

func parseFieldValue(field *reflectr.StructMeta, fieldValue string) error {
	r, err := field.
		Method("Parse").
		Returns(reflectr.ErrorType).
		Call(fieldValue)
	if err != nil {
		return errParseFieldValue
	}
	return returnToError(r[0])
}

func getFieldValue(field *reflectr.StructMeta) (interface{}, error) {
	r, err := field.Method("Value").Call()
	if err != nil {
		return nil, errGetFieldValue
	}
	if len(r) != 1 {
		return nil, errGetFieldValue
	}
	return r[0], nil
}

func runValidator(validator interface{}, v interface{}) (error, error) {
	r, err := reflectr.
		Struct(validator).
		Method("Validate").
		Returns(reflectr.ErrorType).
		Call(v)
	if err != nil {
		return nil, errRunValidator
	}
	if len(r) != 1 {
		return nil, errRunValidator
	}
	return returnToError(r[0]), nil
}

// Validate parses request data and validates it against the provided form.
func Validate(r *http.Request, v interface{}) (bool, error) {
	s := reflectr.Struct(v)
	if !s.IsPtr() {
		return false, errInvalidForm
	}
	if err := s.Error(); err != nil {
		return false, errInvalidForm
	}
	validated := true
	for _, fieldName := range s.Fields() {
		var (
			f, _       = s.Field(fieldName).Value()
			field      = reflectr.Struct(f)
			fieldValue = r.Form.Get(fieldName)
		)
		if !field.IsPtr() {
			return false, errInvalidField
		}
		if err := parseFieldValue(field, fieldValue); err != nil {
			return false, err
		}
		v, err := getFieldValue(field)
		if err != nil {
			return false, err
		}
		i, err := field.Field("Field").Type(Field{}).Addr()
		if err != nil {
			return false, err
		}
		fieldField := i.(*Field)
		for _, validator := range fieldField.Validators {
			e, err := runValidator(validator, v)
			if err != nil {
				return false, err
			}
			if e != nil {
				fieldField.Error = e
				validated = false
			}
		}
	}
	return validated, nil
}

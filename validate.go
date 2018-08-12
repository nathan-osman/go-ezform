package ezform

import (
	"errors"
	"net/http"

	"github.com/nathan-osman/go-ezform/fields"
	"github.com/nathan-osman/go-reflectr"
)

var (
	errInvalidForm  = errors.New("form must be a pointer to struct")
	errInvalidField = errors.New("field must be a pointer to struct")
)

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
		if err := parse(field, fieldValue); err != nil {
			return false, err
		}
		v, err := value(field)
		if err != nil {
			return false, err
		}
		i, err := field.Field("Field").Type(fields.Field{}).Addr()
		if err != nil {
			return false, err
		}
		fieldField := i.(*fields.Field)
		for _, validator := range fieldField.Validators {
			e, err := run(validator, v)
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

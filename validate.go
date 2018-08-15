package ezform

import (
	"errors"
	"net/http"

	"github.com/nathan-osman/go-reflectr"
)

var (
	// ErrInvalid indicates that the form failed validation.
	ErrInvalid = errors.New("form failed validation")

	errInvalidForm = errors.New("form must be a pointer to struct")
)

// Validate parses request data and validates it against the provided form.
func Validate(r *http.Request, v interface{}) error {
	s := reflectr.Struct(v)
	if !s.IsPtr() {
		return errInvalidForm
	}
	if err := s.Error(); err != nil {
		return errInvalidForm
	}
	validated := true
	for _, fieldName := range s.Fields() {
		f, _ := s.Field(fieldName).Value()
		if err := field(f, r.Form.Get(fieldName)); err != nil {
			if err == errInvalidValue {
				validated = false
			} else {
				return err
			}
		}
	}
	if !validated {
		return ErrInvalid
	}
	return nil
}

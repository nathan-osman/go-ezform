package ezform

import (
	"errors"

	"github.com/nathan-osman/go-reflectr"
)

var (
	errValidator = errors.New("unable to use the validator on the field")
)

// Field is designed to be used as a base for all field types. It provides the Validate and Error methods needed for compatibility with Validate.
type Field struct {
	validators []interface{}
	err        error
}

// Validate uses each validator specified for the field to ensure that the field contains a valid value.
func (f Field) Validate(value interface{}) error {
	for _, v := range f.validators {
		r, err := reflectr.
			Struct(v).
			Method("Validate").
			Returns(reflectr.ErrorType).
			Call(value)
		if err != nil {
			return errValidator
		}
		err = r[0].(error)
		if err != nil {
			f.err = err
		}
	}
	return nil
}

// Error returns the error for the field (if any).
func (f Field) Error() error {
	return f.err
}

// SetError sets the error for the field.
func (f *Field) SetError(err error) {
	f.err = err
}

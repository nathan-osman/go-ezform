package ezform

import (
	"errors"

	"github.com/nathan-osman/go-ezform/fields"
	"github.com/nathan-osman/go-reflectr"
)

var (
	errInvalidField          = errors.New("field must be a pointer to struct")
	errFieldValidationFailed = errors.New("field failed validation")
)

// field parses the value provided for the field and validates it.
func field(f interface{}, fieldValue string) error {
	field := reflectr.Struct(f)
	if !field.IsPtr() {
		return errInvalidField
	}
	if err := parse(field, fieldValue); err != nil {
		return err
	}
	v, err := value(field)
	if err != nil {
		return err
	}
	i, err := field.Field("Field").Type(fields.Field{}).Addr()
	if err != nil {
		return err
	}
	fieldField := i.(*fields.Field)
	for _, validator := range fieldField.Validators {
		e, err := run(validator, v)
		if err != nil {
			return err
		}
		if e != nil {
			fieldField.Error = e
			return errFieldValidationFailed
		}
	}
	return nil
}

package ezform

import (
	"errors"

	"github.com/nathan-osman/go-reflectr"
)

var (
	errGetFieldValue = errors.New("unable to get field value")
)

// value calls the field's Value() method to retrieve the parsed value.
func value(field *reflectr.StructMeta) (interface{}, error) {
	r, err := field.Method("Value").Call()
	if err != nil {
		return nil, errGetFieldValue
	}
	if len(r) != 1 {
		return nil, errGetFieldValue
	}
	return r[0], nil
}

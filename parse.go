package ezform

import (
	"errors"

	"github.com/nathan-osman/go-reflectr"
)

var (
	errParseFieldValue = errors.New("unable to parse form value")
)

// parse calls the Parse() method of the field with the provided value.
// The return value is the error returned by the call to Parse().
func parse(field *reflectr.StructMeta, fieldValue string) error {
	r, err := field.
		Method("Parse").
		Returns(reflectr.ErrorType).
		Call(fieldValue)
	if err != nil {
		return errParseFieldValue
	}
	return interfaceToError(r[0])
}

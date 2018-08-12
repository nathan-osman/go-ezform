package ezform

import (
	"errors"

	"github.com/nathan-osman/go-reflectr"
)

var (
	errRunValidator = errors.New("unable to run validator")
)

// run runs the provided value against the provided validator.
// The first return value is the error returned by the validator.
func run(validator interface{}, v interface{}) (error, error) {
	r, err := reflectr.
		Struct(validator).
		Method("Validate").
		Returns(reflectr.ErrorType).
		Call(v)
	if err != nil {
		return nil, errRunValidator
	}
	return interfaceToError(r[0]), nil
}

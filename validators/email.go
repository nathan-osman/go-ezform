package validators

import (
	"errors"
	"regexp"
)

var (
	errInvalidEmail = errors.New("invalid email address")

	regexpEmail = regexp.MustCompile(`^.+@.+$`)
)

// Email ensures that the provided value represents a valid email address.
type Email struct{}

// Validate ensures the string is a valid email address.
func (e Email) Validate(v string) error {
	if !regexpEmail.MatchString(v) {
		return errInvalidEmail
	}
	return nil
}

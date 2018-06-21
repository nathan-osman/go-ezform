package ezform

import (
	"errors"
	"net/url"
)

var (
	// ErrInvalidURL indicates that the provided value is not a valid URL.
	ErrInvalidURL = errors.New("value is not a valid URL")
)

// IsURL verifies that a string represents a valid URL.
func IsURL(v string) error {
	if _, err := url.Parse(v); err != nil {
		return ErrInvalidURL
	}
	return nil
}

package ezform

import "errors"

var (

	// ErrStringTooShort indicates that the provided value is too short.
	ErrStringTooShort = errors.New("value is too short")

	// ErrStringTooLong indicates that the provided value is too long.
	ErrStringTooLong = errors.New("value is too long")
)

// String represents a field that stores a string value.
type String string

func (s String) String() string {
	return string(s)
}

// Parse stores the provided string.
func (s *String) Parse(str string) error {
	*s = String(str)
	return nil
}

// MinLength verifies that the length of the string is at least min characters.
func (s String) MinLength(min int) error {
	if len(s) < min {
		return ErrStringTooShort
	}
	return nil
}

// MaxLength verifies that the length of the string is no more than max characters.
func (s String) MaxLength(max int) error {
	if len(s) > max {
		return ErrStringTooLong
	}
	return nil
}

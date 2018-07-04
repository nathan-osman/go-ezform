package ezform

import (
	"errors"
	"strconv"
)

var (

	// ErrInvalidInteger indicates that an invalid integer was supplied.
	ErrInvalidInteger = errors.New("invalid integer specified")

	// ErrIntegerTooSmall indicates that the provided value is too small.
	ErrIntegerTooSmall = errors.New("integer is too small")

	// ErrIntegerTooLarge indicates that the provided value is too large.
	ErrIntegerTooLarge = errors.New("integer is too large")
)

// Integer represents a field that stores an integer value.
type Integer int64

func (i Integer) String() string {
	return strconv.FormatInt(int64(i), 10)
}

// Parse attempts to store the provided string as an integer.
func (i *Integer) Parse(str string) error {
	v, err := strconv.ParseInt(str, 10, 64)
	*i = Integer(v)
	return err
}

// MinValue verifies that the number is equal to or greater than the provided value.
func (i Integer) MinValue(min int64) error {
	if int64(i) < min {
		return ErrIntegerTooSmall
	}
	return nil
}

// MaxValue verifies that the number is less than or equal to the provided value.
func (i Integer) MaxValue(max int64) error {
	if int64(i) < max {
		return ErrIntegerTooLarge
	}
	return nil
}

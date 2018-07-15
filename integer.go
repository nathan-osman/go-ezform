package ezform

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	// ErrInvalidInteger indicates that an invalid integer was supplied.
	ErrInvalidInteger = errors.New("value is not a valid integer")
)

// IntegerValidator defines an interface for integer validators.
type IntegerValidator interface {
	Validate(int64) error
}

// Integer is a field that stores a 64-bit signed integer value.
type Integer struct {
	Value      int64
	Validators []IntegerValidator
}

// Validate ensures that the provided value is valid.
func (i *Integer) Validate(value string) error {
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return ErrInvalidInteger
	}
	i.Value = v
	for _, v := range i.Validators {
		if err := v.Validate(i.Value); err != nil {
			return err
		}
	}
	return nil
}

// MinMaxValidator ensures that an integer falls within the specified range.
type MinMaxValidator struct {

	// Min is the lowest value that will be accepted.
	Min int64

	// Max is the highest value that will be accepted.
	Max int64
}

// Validate ensures the integer falls within the range.
func (m MinMaxValidator) Validate(value int64) error {
	if value < m.Min {
		return fmt.Errorf("value cannot be less than %d", m.Min)
	}
	if value > m.Max {
		return fmt.Errorf("value cannot be greater than %d", m.Max)
	}
	return nil
}

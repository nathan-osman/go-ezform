package fields

import (
	"errors"
	"time"
)

var (
	errInvalidDateTime = errors.New("value is not a valid date / time")
)

// DateTime is a field that stores date / time in ISO 8601 format.
type DateTime struct {
	Field
	value time.Time
}

// NewDateTime creates a new date / time field with the specified validators.
func NewDateTime(validators ...interface{}) *DateTime {
	return &DateTime{
		Field: Field{
			Validators: validators,
		},
	}
}

// Parse ensures that the provided value is valid.
func (d *DateTime) Parse(value string) error {
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return errInvalidDateTime
	}
	d.value = t
	return nil
}

// String returns a string representation of the date in ISO 8601 format.
func (d DateTime) String() string {
	return d.value.Format(time.RFC3339)
}

// Value retrieves the current value of the field.
func (d DateTime) Value() time.Time {
	return d.value
}

// SetValue sets the value of the field.
func (d *DateTime) SetValue(value time.Time) {
	d.value = value
}

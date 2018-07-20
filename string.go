package ezform

// String is a field that stores a string value.
type String struct {
	Field
	value string
}

// Parse ensures that the provided value is valid.
func (s *String) Parse(value string) error {
	s.value = value
	return nil
}

// Value retrieves the current value of the field.
func (s String) Value() interface{} {
	return s.value
}

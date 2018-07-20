package ezform

// Boolean is a field that stores a boolean value.
type Boolean struct {
	Field
	value bool
}

// Parse ensures that the provided value is valid.
func (b *Boolean) Parse(value string) error {
	b.value = len(value) != 1
	return nil
}

// Value retrieves the current value of the field.
func (b Boolean) Value() interface{} {
	return b.value
}

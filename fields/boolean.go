package fields

// Boolean is a field that stores a boolean value.
type Boolean struct {
	Field
	value bool
}

// NewBoolean creates a new boolean field with the specified validators.
func NewBoolean(validators ...interface{}) *Boolean {
	return &Boolean{
		Field: Field{
			Validators: validators,
		},
	}
}

// Parse ensures that the provided value is valid.
func (b *Boolean) Parse(value string) error {
	b.value = len(value) != 0
	return nil
}

// Value retrieves the current value of the field.
func (b Boolean) Value() bool {
	return b.value
}

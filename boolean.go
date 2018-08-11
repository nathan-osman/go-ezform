package ezform

// BooleanField is a field that stores a boolean value.
type BooleanField struct {
	Field
	value bool
}

// NewBooleanField creates a new boolean field with the specified validators.
func NewBooleanField(validators ...interface{}) *BooleanField {
	return &BooleanField{
		Field: Field{
			Validators: validators,
		},
	}
}

// Parse ensures that the provided value is valid.
func (b *BooleanField) Parse(value string) error {
	b.value = len(value) != 0
	return nil
}

// Value retrieves the current value of the field.
func (b BooleanField) Value() bool {
	return b.value
}

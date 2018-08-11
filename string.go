package ezform

// StringField is a field that stores a string value.
type StringField struct {
	Field
	value string
}

// NewStringField creates a new string field with the specified validators.
func NewStringField(validators ...interface{}) *StringField {
	return &StringField{
		Field: Field{
			Validators: validators,
		},
	}
}

// Parse ensures that the provided value is valid.
func (s *StringField) Parse(value string) error {
	s.value = value
	return nil
}

// Value retrieves the current value of the field.
func (s StringField) Value() string {
	return s.value
}

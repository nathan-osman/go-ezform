package fields

// String is a field that stores a string value.
type String struct {
	Field
	value string
}

// NewString creates a new string field with the specified validators.
func NewString(validators ...interface{}) *String {
	return &String{
		Field: Field{
			Validators: validators,
		},
	}
}

// Parse ensures that the provided value is valid.
func (s *String) Parse(value string) error {
	s.value = value
	return nil
}

// String returns a string representation of the field.
func (s String) String() string {
	return s.value
}

// Value retrieves the current value of the field.
func (s String) Value() string {
	return s.value
}

// SetValue sets the value of the field.
func (s *String) SetValue(value string) {
	s.value = value
}

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

// NewStringWithDefault creates a new string field with the specified default value and validators.
func NewStringWithDefault(value string, validators ...interface{}) *String {
	return &String{
		Field: Field{
			Validators: validators,
		},
		value: value,
	}
}

// Parse ensures that the provided value is valid.
func (s *String) Parse(value string) error {
	s.value = value
	return nil
}

// Value retrieves the current value of the field.
func (s String) Value() string {
	return s.value
}

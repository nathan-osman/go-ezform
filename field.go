package ezform

// Field is the interface that all field types must implement.
type Field interface {

	// Validate parses the value provided for the field and validates it.
	Validate(string) error
}

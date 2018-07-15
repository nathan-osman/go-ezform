package ezform

// Field is the interface that all field types must implement.
type Field interface {

	// Parse attempts to store the provided string value in the field's native type.
	Parse(string) error

	// Validate ensures that the provided value is valid.
	Validate() error
}

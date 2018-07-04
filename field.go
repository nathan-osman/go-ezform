package ezform

// Field is the base interface for all of the supported form field types.
type Field interface {

	// String obtains a string representation of the field's value.
	String() string

	// Parse attempts to load the provided string representation of the field.
	Parse(string) error
}

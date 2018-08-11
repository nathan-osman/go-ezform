package ezform

// Field is designed to be used as a base for all field types. It provides the Validate and Error methods needed for compatibility with Validate.
type Field struct {
	Validators []interface{}
	Error      error
}

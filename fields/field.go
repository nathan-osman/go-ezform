package fields

// Field is designed to be used as a base for all field types.
// It provides the Validators and Error fields required by ezform.Validate().
type Field struct {
	Validators []interface{}
	Error      error
}

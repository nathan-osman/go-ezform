package ezform

// interfaceToError converts an interface{} to an error.
// Because v is an interface{} and may be nil, type assertion cannot be used directly.
func interfaceToError(v interface{}) error {
	if v == nil {
		return nil
	}
	return v.(error)
}

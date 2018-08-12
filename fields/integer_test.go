package fields

import (
	"testing"
)

func TestInteger(t *testing.T) {
	for _, test := range []struct {
		Input  string
		Output int64
		Error  error
	}{
		{Input: "", Error: errInvalidInteger},
		{Input: "-42", Output: -42, Error: nil},
	} {
		f := NewInteger()
		if err := f.Parse(test.Input); err != test.Error {
			t.Fatalf("%v != %v", err, test.Error)
		}
		if f.Value() != test.Output {
			t.Fatalf("%v != %v", f.Value(), strVal)
		}
	}
}

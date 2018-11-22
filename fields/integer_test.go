package fields

import (
	"testing"
)

const intVal = 42

func TestNewInteger(t *testing.T) {
	for _, test := range []struct {
		Input  string
		Output int64
		Error  error
	}{
		{Input: "", Error: errInvalidInteger},
		{Input: "-42", Output: -42},
	} {
		f := NewInteger()
		if err := f.Parse(test.Input); err != test.Error {
			t.Fatalf("%v != %v", err, test.Error)
		}
		if f.Value() != test.Output {
			t.Fatalf("%v != %v", f.Value(), test.Output)
		}
	}
}

func TestIntegerSetValue(t *testing.T) {
	f := NewInteger()
	f.SetValue(intVal)
	if f.Value() != intVal {
		t.Fatalf("%v != %v", f.Value(), intVal)
	}
}

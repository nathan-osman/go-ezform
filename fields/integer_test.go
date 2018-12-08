package fields

import (
	"testing"
)

const (
	intVal    = -42
	intValStr = "-42"
)

func TestNewInteger(t *testing.T) {
	for _, test := range []struct {
		Input  string
		Error  error
		String string
		Value  int64
	}{
		{Input: "", Error: errInvalidInteger},
		{Input: intValStr, String: intValStr, Value: intVal},
	} {
		f := NewInteger()
		if err := f.Parse(test.Input); err != test.Error {
			t.Fatalf("%v != %v", err, test.Error)
		}
		if test.Error == nil {
			if f.String() != test.String {
				t.Fatalf("%v != %v", f.String(), test.String)
			}
			if f.Value() != test.Value {
				t.Fatalf("%v != %v", f.Value(), test.Value)
			}
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

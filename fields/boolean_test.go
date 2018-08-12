package fields

import (
	"testing"
)

func TestBoolean(t *testing.T) {
	for _, test := range []struct {
		Input  string
		Output bool
	}{
		{Input: "", Output: false},
		{Input: "false", Output: true},
	} {
		f := NewBoolean()
		if err := f.Parse(test.Input); err != nil {
			t.Fatal(err)
		}
		if f.Value() != test.Output {
			t.Fatalf("%v != %v", f.Value(), test.Output)
		}
	}
}
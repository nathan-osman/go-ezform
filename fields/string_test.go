package fields

import (
	"testing"
)

const strVal = "a b\n"

func TestNewString(t *testing.T) {
	f := NewString()
	if err := f.Parse(strVal); err != nil {
		t.Fatal(err)
	}
	if f.Value() != strVal {
		t.Fatalf("%v != %v", f.Value(), strVal)
	}
}

func TestStringSetValue(t *testing.T) {
	f := NewString()
	f.SetValue(strVal)
	if f.Value() != strVal {
		t.Fatalf("%v != %v", f.Value(), strVal)
	}
}

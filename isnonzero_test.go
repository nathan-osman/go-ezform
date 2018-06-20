package ezform

import "testing"

func TestIsNonZeroWithNonZero(t *testing.T) {
	if err := IsNonZero(int(12)); err != nil {
		t.Fatalf("%v != nil", err)
	}
}

func TestIsNonZeroWithZero(t *testing.T) {
	if err := IsNonZero(int(0)); err != ErrValueIsZero {
		t.Fatalf("%v != %v", err, ErrValueIsZero)
	}
}

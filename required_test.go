package ezform

import "testing"

func TestRequiredNonZero(t *testing.T) {
	if err := Required(int(12)); err != nil {
		t.Fatalf("%v != nil", err)
	}
}

func TestRequiredZero(t *testing.T) {
	if err := Required(int(0)); err != ErrFieldRequired {
		t.Fatalf("%v != %v", err, ErrFieldRequired)
	}
}

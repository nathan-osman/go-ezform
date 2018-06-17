package ezform

import "testing"

const (
	strValid   = "valid"
	strInvalid = "invalid"
)

var strSlice = []string{strValid}

func TestInvalidSlice(t *testing.T) {
	if err := InSlice(0, 0); err != ErrInvalidValue {
		t.Fatalf("%v != %v", err, ErrInvalidValue)
	}
}

func TestInvalidValue(t *testing.T) {
	if err := InSlice(strSlice, 0); err != ErrInvalidValue {
		t.Fatalf("%v != %v", err, ErrInvalidValue)
	}
}

func TestValueNotInSlice(t *testing.T) {
	if err := InSlice(strSlice, strInvalid); err != ErrInvalidValue {
		t.Fatalf("%v != %v", err, ErrInvalidValue)
	}
}

func TestValueInSlice(t *testing.T) {
	if err := InSlice(strSlice, strValid); err != nil {
		t.Fatalf("%v != nil", err)
	}
}

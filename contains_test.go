package ezform

import "testing"

const (
	strValid   = "valid"
	strInvalid = "invalid"
)

var (
	strSlice = []string{strValid}
	strMap   = map[string]interface{}{
		strValid: nil,
	}
)

func TestInvalidContainer(t *testing.T) {
	if err := Contains(0, 0); err != ErrInvalidValue {
		t.Fatalf("%v != %v", err, ErrInvalidValue)
	}
}

func TestInvalidValue(t *testing.T) {
	if err := Contains(strSlice, 0); err != ErrInvalidValue {
		t.Fatalf("%v != %v", err, ErrInvalidValue)
	}
	if err := Contains(strMap, 0); err != ErrInvalidValue {
		t.Fatalf("%v != %v", err, ErrInvalidValue)
	}
}

func TestValueNotInContainer(t *testing.T) {
	if err := Contains(strSlice, strInvalid); err != ErrInvalidValue {
		t.Fatalf("%v != %v", err, ErrInvalidValue)
	}
	if err := Contains(strMap, strInvalid); err != ErrInvalidValue {
		t.Fatalf("%v != %v", err, ErrInvalidValue)
	}
}

func TestValueInContainer(t *testing.T) {
	if err := Contains(strSlice, strValid); err != nil {
		t.Fatalf("%v != nil", err)
	}
	if err := Contains(strMap, strValid); err != nil {
		t.Fatalf("%v != nil", err)
	}
}

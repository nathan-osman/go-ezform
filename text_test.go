package ezform

import "testing"

const (
	validURL   = "http://example.com"
	invalidURL = ":"
)

func TestIsURL(t *testing.T) {
	if err := IsURL(validURL); err != nil {
		t.Fatalf("%v != nil", err)
	}
	if err := IsURL(invalidURL); err != ErrInvalidURL {
		t.Fatalf("%v != %v", err, ErrInvalidURL)
	}
}

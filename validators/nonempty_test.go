package validators

import (
	"testing"
)

func TestNonEmpty(t *testing.T) {
	for _, test := range []struct {
		Input string
		Error error
	}{
		{Input: "", Error: errNonEmptyString},
		{Input: " ", Error: nil},
	} {
		n := &NonEmpty{}
		if err := n.Validate(test.Input); err != test.Error {
			t.Fatalf("%v != %v", err, test.Error)
		}
	}
}

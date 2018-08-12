package validators

import (
	"testing"
)

func TestRequired(t *testing.T) {
	for _, test := range []struct {
		Input bool
		Error error
	}{
		{Input: false, Error: errRequired},
		{Input: true, Error: nil},
	} {
		r := &Required{}
		if err := r.Validate(test.Input); err != test.Error {
			t.Fatalf("%v != %v", err, test.Error)
		}
	}
}

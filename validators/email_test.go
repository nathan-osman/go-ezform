package validators

import (
	"testing"
)

func TestEmail(t *testing.T) {
	for _, test := range []struct {
		Input string
		Error error
	}{
		{Input: "", Error: errInvalidEmail},
		{Input: "a", Error: errInvalidEmail},
		{Input: "a@", Error: errInvalidEmail},
		{Input: "a@a", Error: nil},
	} {
		e := &Email{}
		if err := e.Validate(test.Input); err != test.Error {
			t.Fatalf("%v != %v", err, test.Error)
		}
	}
}

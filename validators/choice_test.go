package validators

import (
	"testing"
)

func TestChoice(t *testing.T) {
	for _, test := range []struct {
		Choices interface{}
		Input   interface{}
		Error   error
	}{
		{Choices: 1, Error: errInvalidType},
		{Choices: []string{}, Input: 1, Error: errInvalidType},
		{Choices: []string{}, Input: "a", Error: errInvalidValue},
		{Choices: []string{"a"}, Input: "a", Error: nil},
		{Choices: map[string]interface{}{}, Input: 1, Error: errInvalidType},
		{Choices: map[string]interface{}{}, Input: "a", Error: errInvalidValue},
		{Choices: map[string]interface{}{"a": nil}, Input: "a", Error: nil},
	} {
		c := &Choice{Choices: test.Choices}
		if err := c.Validate(test.Input); err != test.Error {
			t.Fatalf("%v != %v", err, test.Error)
		}
	}
}

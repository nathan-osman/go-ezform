package ezform

import (
	"testing"
)

type testBadReturnRunValidator struct{}

func (t testBadReturnRunValidator) Validate(string) {}

type testBadTypeRunValidator struct{}

func (t testBadTypeRunValidator) Validate(int) {}

type testRunValidator struct{}

func (t testRunValidator) Validate(string) error { return nil }

func TestRun(t *testing.T) {
	for _, test := range []struct {
		Validator interface{}
		Error     error
	}{
		{Validator: &testBadReturnRunValidator{}, Error: errRunValidator},
		{Validator: &testBadTypeRunValidator{}, Error: errRunValidator},
		{Validator: &testRunValidator{}},
	} {
		if _, err := run(test.Validator, ""); err != test.Error {
			t.Fatalf("%v != %v", err, test.Error)
		}
	}
}

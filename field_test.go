package ezform

import (
	"testing"

	"github.com/nathan-osman/go-ezform/validators"

	"github.com/nathan-osman/go-ezform/fields"
	"github.com/nathan-osman/go-reflectr"
)

type testFieldMissingParse struct{}

type testFieldMissingValue struct{}

func (t *testFieldMissingValue) Parse(string) error { return nil }

type testFieldMissingField struct{}

func (t *testFieldMissingField) Parse(string) error { return nil }
func (t *testFieldMissingField) Value() interface{} { return nil }

type testFieldBadValidator struct{}

func TestField(t *testing.T) {
	for _, test := range []struct {
		F          interface{}
		FieldValue string
		Error      error
	}{
		{F: struct{}{}, Error: errInvalidField},
		{F: &testFieldMissingParse{}, Error: errParseFieldValue},
		{F: &testFieldMissingValue{}, Error: errGetFieldValue},
		{F: &testFieldMissingField{}, Error: reflectr.ErrFieldDoesNotExist},
		{F: fields.NewString(&testFieldBadValidator{}), Error: errRunValidator},
		{F: fields.NewString(&validators.NonEmpty{}), Error: errFieldValidationFailed},
		{F: fields.NewString(&validators.NonEmpty{}), FieldValue: "a"},
	} {
		if err := field(test.F, test.FieldValue); err != test.Error {
			t.Fatalf("%v != %v", err, test.Error)
		}
	}
}

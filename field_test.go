package ezform

import (
	"testing"

	reflectr "github.com/nathan-osman/go-reflectr"
)

type testFieldMissingParseField struct{}

type testFieldMissingValueField struct{}

func (t *testFieldMissingValueField) Parse(value string) error { return nil }

type testBadValueReturnField struct{}

func (t *testBadValueReturnField) Parse(value string) error { return nil }
func (t *testBadValueReturnField) Value()                   {}

func TestBadFields(t *testing.T) {
	for _, test := range []struct {
		Field      interface{}
		Value      string
		ParseError error
		ValueError error
	}{
		{
			Field:      &testFieldMissingParseField{},
			ParseError: errParseFieldValue,
		},
		{
			Field:      &testFieldMissingValueField{},
			ValueError: errGetFieldValue,
		},
		{
			Field:      &testBadValueReturnField{},
			ValueError: errGetFieldValue,
		},
	} {
		s := reflectr.Struct(test.Field)
		if err := parseFieldValue(s, test.Value); err != test.ParseError {
			t.Fatalf("%v != %v", err, test.ParseError)
		}
		if test.ParseError != nil {
			continue
		}
		_, err := getFieldValue(s)
		if err != test.ValueError {
			t.Fatalf("%v != %v", err, test.ValueError)
		}
		if test.ValueError != nil {
			continue
		}
	}
}

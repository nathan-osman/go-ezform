package ezform

import (
	"errors"
	"testing"

	"github.com/nathan-osman/go-ezform/fields"
	"github.com/nathan-osman/go-reflectr"
)

var errParseBadValue = errors.New("test")

type testParseBadField struct{}

type testParseBadValueField struct{}

func (t *testParseBadValueField) Parse(v string) error { return errParseBadValue }

func TestParse(t *testing.T) {
	for _, test := range []struct {
		Field      interface{}
		ParseError error
		Error      error
	}{
		{Field: &testParseBadField{}, Error: errParseFieldValue},
		{Field: &testParseBadValueField{}, ParseError: errParseBadValue},
		{Field: fields.NewString()},
	} {
		e, err := parse(reflectr.Struct(test.Field), "")
		if err != test.Error {
			t.Fatalf("%v != %v", err, test.Error)
		}
		if e != test.ParseError {
			t.Fatalf("%v != %v", e, test.ParseError)
		}
	}
}

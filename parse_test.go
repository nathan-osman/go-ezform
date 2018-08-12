package ezform

import (
	"testing"

	"github.com/nathan-osman/go-ezform/fields"

	"github.com/nathan-osman/go-reflectr"
)

type testBadParseField struct{}

func TestBadParse(t *testing.T) {
	s := reflectr.Struct(&testBadParseField{})
	if err := parse(s, ""); err != errParseFieldValue {
		t.Fatalf("%v != %v", err, errParseFieldValue)
	}
}

func TestParse(t *testing.T) {
	s := reflectr.Struct(fields.NewString())
	if err := parse(s, ""); err != nil {
		t.Fatal(err)
	}
}

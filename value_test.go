package ezform

import (
	"testing"

	"github.com/nathan-osman/go-reflectr"
)

type testBadValueField1 struct{}

type testBadValueField2 struct{}

func (t testBadValueField2) Value() {}

func TestBadValue(t *testing.T) {
	for _, s := range []interface{}{
		&testBadValueField1{},
		&testBadValueField2{},
	} {
		if _, err := value(reflectr.Struct(s)); err != errGetFieldValue {
			t.Fatalf("%v != %v", err, errGetFieldValue)
		}
	}
}

type testValueField struct{}

func (t testValueField) Value() string { return strVal }

func TestValue(t *testing.T) {
	v, err := value(reflectr.Struct(&testValueField{}))
	if err != nil {
		t.Fatal(err)
	}
	if v.(string) != strVal {
		t.Fatalf("%v != %v", v.(string), strVal)
	}
}

package ezform

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/nathan-osman/go-ezform/fields"
)

const (
	stringVal  = "a b\nc"
	integerVal = -42
	booleanVal = true
)

func simulateRequest(params url.Values, v interface{}) (bool, error) {
	r := httptest.NewRequest(
		http.MethodPost,
		"/",
		strings.NewReader(params.Encode()),
	)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err := r.ParseForm(); err != nil {
		return false, err
	}
	return Validate(r, v)
}

type testValidateForm struct {
	StringVal  *fields.String
	IntegerVal *fields.Integer
	BooleanVal *fields.Boolean
}

func TestValidate(t *testing.T) {
	f := &testValidateForm{
		StringVal:  fields.NewString(),
		IntegerVal: fields.NewInteger(),
		BooleanVal: fields.NewBoolean(),
	}
	ok, err := simulateRequest(
		url.Values{
			"StringVal":  []string{stringVal},
			"IntegerVal": []string{fmt.Sprintf("%v", integerVal)},
			"BooleanVal": []string{fmt.Sprintf("%v", booleanVal)},
		},
		f,
	)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("validation failed")
	}
	if f.StringVal.Value() != stringVal {
		t.Fatalf("%v != %v", f.StringVal.Value(), stringVal)
	}
	if f.IntegerVal.Value() != integerVal {
		t.Fatalf("%v != %v", f.IntegerVal.Value(), integerVal)
	}
	if f.BooleanVal.Value() != booleanVal {
		t.Fatalf("%v != %v", f.BooleanVal.Value(), booleanVal)
	}
}

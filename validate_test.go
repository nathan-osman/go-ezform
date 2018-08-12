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
	strVal  = "a b\n"
	intVal  = 42
	boolVal = true
)

func createRequest(params url.Values) (*http.Request, error) {
	r := httptest.NewRequest(
		http.MethodPost,
		"/",
		strings.NewReader(params.Encode()),
	)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err := r.ParseForm(); err != nil {
		return nil, err
	}
	return r, nil
}

type testBadValidateForm struct{}

func TestBadValidate(t *testing.T) {
	if _, err := Validate(nil, new(string)); err != errInvalidForm {
		t.Fatalf("%v != %v", err, errInvalidForm)
	}
	if _, err := Validate(nil, testBadValidateForm{}); err != errInvalidForm {
		t.Fatalf("%v != %v", err, errInvalidForm)
	}
}

type testValidateForm struct {
	StrField  *fields.String
	IntField  *fields.Integer
	BoolField *fields.Boolean
}

func TestValidate(t *testing.T) {
	f := &testValidateForm{
		StrField:  fields.NewString(),
		IntField:  fields.NewInteger(),
		BoolField: fields.NewBoolean(),
	}
	r, err := createRequest(url.Values{
		"StrField":  []string{strVal},
		"IntField":  []string{fmt.Sprintf("%v", intVal)},
		"BoolField": []string{fmt.Sprintf("%v", boolVal)},
	})
	if err != nil {
		t.Fatal(err)
	}
	v, err := Validate(r, f)
	if err != nil {
		t.Fatal(err)
	}
	if !v {
		t.Fatal("validation unexpectedly failed")
	}
	if f.StrField.Value() != strVal {
		t.Fatalf("%v != %v", f.StrField.Value(), strVal)
	}
	if f.IntField.Value() != intVal {
		t.Fatalf("%v != %v", f.IntField.Value(), intVal)
	}
	if f.BoolField.Value() != boolVal {
		t.Fatalf("%v != %v", f.BoolField.Value(), boolVal)
	}
}

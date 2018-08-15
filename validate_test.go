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

var (
	testParams = url.Values{
		"StrField":  []string{strVal},
		"IntField":  []string{fmt.Sprintf("%v", intVal)},
		"BoolField": []string{fmt.Sprintf("%v", boolVal)},
	}
)

type testValidateForm struct {
	StrField  *fields.String
	IntField  *fields.Integer
	BoolField *fields.Boolean
}

type testValidateBadForm struct {
	Field int
}

func newTestValidateForm() *testValidateForm {
	return &testValidateForm{
		StrField:  fields.NewString(),
		IntField:  fields.NewInteger(),
		BoolField: fields.NewBoolean(),
	}
}

func TestValidate(t *testing.T) {
	for _, test := range []struct {
		Params url.Values
		Form   interface{}
		Error  error
	}{
		{Form: testValidateForm{}, Error: errInvalidForm},
		{Form: new(string), Error: errInvalidForm},
		{Form: newTestValidateForm(), Error: ErrInvalid},
		{Form: &testValidateBadForm{}, Error: errInvalidField},
		{Params: testParams, Form: newTestValidateForm()},
	} {
		r := httptest.NewRequest(
			http.MethodPost,
			"/",
			strings.NewReader(test.Params.Encode()),
		)
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}
		if err := Validate(r, test.Form); err != test.Error {
			t.Fatalf("%v != %v", err, test.Error)
		}
	}
}

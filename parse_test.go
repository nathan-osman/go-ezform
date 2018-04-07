package ezform

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

type TestForm struct {
	StringVal string `form:"required"`
	IntVal    int64
	BoolVal   bool
}

var TestFormData = []struct {
	Data   url.Values
	Form   *TestForm
	Errors map[string][]error
}{
	{
		Data: url.Values{},
		Errors: map[string][]error{
			"StringVal": []error{ErrFieldRequired},
		},
	},
	{
		Data: url.Values{
			"StringVal": []string{"test"},
			"IntVal":    []string{"123"},
			"BoolVal":   []string{"checked"},
		},
		Form: &TestForm{
			StringVal: "test",
			IntVal:    123,
			BoolVal:   true,
		},
	},
}

func TestParseEmptyForm(t *testing.T) {
	for _, v := range TestFormData {
		var (
			r = httptest.NewRequest(
				http.MethodPost,
				"/",
				strings.NewReader(v.Data.Encode()),
			)
			f = &TestForm{}
		)
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		e, err := Parse(r, f)
		if err != nil {
			t.Fatal(err)
		}
		if len(v.Errors) == 0 {
			if !reflect.DeepEqual(f, v.Form) {
				t.Fatalf("%#v != %#v", f, v.Form)
			}
		} else {
			if !reflect.DeepEqual(e, v.Errors) {
				t.Fatalf("%#v != %#v", e, v.Errors)
			}
		}
	}
}

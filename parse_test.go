package ezform

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
)

type parseTestForm struct {
	StringVal string
	IntVal    int
	BoolVal   bool
	TimeVal   time.Time
}

func (p parseTestForm) ValidateStringVal(v string) error {
	return IsNonZero(v)
}

var parseTestReference = &parseTestForm{
	StringVal: "test value\n",
	IntVal:    42,
	BoolVal:   true,
	TimeVal:   time.Time{},
}

func toURLVal(v interface{}) []string {
	return []string{fmt.Sprintf("%v", v)}
}

func simulateRequest(v url.Values, f interface{}) (map[string]error, error) {
	r := httptest.NewRequest(
		http.MethodPost,
		"/",
		strings.NewReader(v.Encode()),
	)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err := r.ParseForm(); err != nil {
		return nil, err
	}
	return Parse(r, f, nil)
}

func TestBadStruct(t *testing.T) {
	if _, err := simulateRequest(url.Values{}, parseTestForm{}); err != errPtrRequired {
		t.Fatalf("%v != %v", err, errPtrRequired)
	}
	var str string
	if _, err := simulateRequest(url.Values{}, &str); err != errPtrRequired {
		t.Fatalf("%v != %v", err, errPtrRequired)
	}
}

func TestMissingField(t *testing.T) {
	fieldErrs, err := simulateRequest(url.Values{}, &parseTestForm{})
	if err != nil {
		t.Fatal(err)
	}
	if len(fieldErrs) != 1 || fieldErrs["StringVal"] != ErrValueIsZero {
		t.Fatal("error expected")
	}
}

func TestStoreValues(t *testing.T) {
	var (
		v = url.Values{
			"StringVal": toURLVal(parseTestReference.StringVal),
			"IntVal":    toURLVal(parseTestReference.IntVal),
			"BoolVal":   toURLVal(parseTestReference.BoolVal),
			"TimeVal":   toURLVal(parseTestReference.TimeVal),
		}
		f = &parseTestForm{}
	)
	fieldErrs, err := simulateRequest(v, f)
	if len(fieldErrs) != 0 || err != nil {
		t.Fatal("errors encountered during storage")
	}
	if !reflect.DeepEqual(f, parseTestReference) {
		t.Fatalf("%v != %v", f, parseTestReference)
	}
}

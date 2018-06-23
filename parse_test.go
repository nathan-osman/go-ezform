package ezform

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

const (
	strVal = "test value\n"
	intVal = 42
)

var (
	timeVal = time.Now()
)

type parseTestForm struct {
	StringVal    string
	IntVal       int
	BoolValTrue  bool
	BoolValFalse bool
	TimeVal      time.Time
}

func (p parseTestForm) ValidateStringVal(v string) error {
	return IsNonZero(v)
}

var parseTestReference = &parseTestForm{
	StringVal:    strVal,
	IntVal:       intVal,
	BoolValTrue:  true,
	BoolValFalse: false,
	TimeVal: time.Date(
		timeVal.Year(),
		timeVal.Month(),
		timeVal.Day(),
		timeVal.Hour(),
		timeVal.Minute(),
		timeVal.Second(),
		0,
		timeVal.Location(),
	),
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
			"StringVal":   []string{strVal},
			"IntVal":      []string{strconv.Itoa(intVal)},
			"BoolValTrue": []string{"on"},
			"TimeVal":     []string{timeVal.Format(time.RFC3339)},
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

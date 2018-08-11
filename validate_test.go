package ezform

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

const (
	stringVal  = "a b\nc"
	integerVal = -42
	booleanVal = true
)

type validateTestForm struct {
	StringVal  *StringField
	IntegerVal *IntegerField
	BooleanVal *BooleanField
}

type ValidateTestFormBadField struct {
	InvalidField string
}

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

func TestValidateBadStruct(t *testing.T) {
	if _, err := simulateRequest(nil, validateTestForm{}); err != errInvalidForm {
		t.Fatalf("%v != %v", err, errInvalidForm)
	}
	var str string
	if _, err := simulateRequest(nil, &str); err == nil {
		t.Fatal("error expected")
	}
}

func TestValidateBadField(t *testing.T) {
	if _, err := simulateRequest(nil, &ValidateTestFormBadField{}); err == nil {
		t.Fatal("error expected")
	}
}

func TestValidate(t *testing.T) {
	f := &validateTestForm{
		StringVal:  NewStringField(),
		IntegerVal: NewIntegerField(),
		BooleanVal: NewBooleanField(),
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
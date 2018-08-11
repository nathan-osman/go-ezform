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

type testValidateBadStructForm struct{}

func TestValidateBadStruct(t *testing.T) {
	if _, err := simulateRequest(nil, testValidateBadStructForm{}); err != errInvalidForm {
		t.Fatalf("%v != %v", err, errInvalidForm)
	}
	var str string
	if _, err := simulateRequest(nil, &str); err == nil {
		t.Fatal("error expected")
	}
}

type testValidateBadFieldForm1 struct {
	InvalidField string
}

type testValidateBadFieldForm2 struct {
	InvalidField StringField
}

func TestValidateBadField(t *testing.T) {
	if _, err := simulateRequest(nil, &testValidateBadFieldForm1{}); err == nil {
		t.Fatal("error expected")
	}
	if _, err := simulateRequest(nil, &testValidateBadFieldForm2{}); err != errInvalidField {
		t.Fatalf("%v != %v", err, errInvalidField)
	}
}

type testValidateForm struct {
	StringVal  *StringField
	IntegerVal *IntegerField
	BooleanVal *BooleanField
}

func TestValidate(t *testing.T) {
	f := &testValidateForm{
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

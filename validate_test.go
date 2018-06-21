package ezform

import (
	"errors"
	"reflect"
	"testing"
)

const (
	strValidParam   = "valid"
	strInvalidParam = "invalid"
)

var (
	errField2 = errors.New("Field2 is invalid")
	errField5 = errors.New("Field5 is invalid")
)

type validateTestForm struct {
	Field1, Field2, Field3, Field4, Field5 string
}

func (v validateTestForm) ValidateField2(s string) error {
	if len(s) == 0 {
		return errField2
	}
	return nil
}

func (v validateTestForm) ValidateField3(i int) error { return nil }
func (v validateTestForm) ValidateField4(s string)    {}
func (v validateTestForm) ValidateField5(s, p string) error {
	if p != strValidParam {
		return errField5
	}
	return nil
}

func validateField(v interface{}, fieldName string, param interface{}) error {
	var (
		vType    = reflect.TypeOf(v)
		sType, _ = vType.FieldByName(fieldName)
		vVal     = reflect.ValueOf(v)
		fVal     = vVal.FieldByName(fieldName)
	)
	return validate(vType, sType, vVal, fVal, param)
}

func TestValidateNoValidator(t *testing.T) {
	if err := validateField(validateTestForm{}, "Field1", nil); err != nil {
		t.Fatalf("%s != nil", err)
	}
}

func TestValidateValidValue(t *testing.T) {
	if err := validateField(validateTestForm{Field2: "test"}, "Field2", nil); err != nil {
		t.Fatalf("%s != nil", err)
	}
}

func TestValidateInvalidValue(t *testing.T) {
	if err := validateField(validateTestForm{}, "Field2", nil); err != errField2 {
		t.Fatalf("%s != %s", err, errField2)
	}
}

func TestValidateInvalidParameter(t *testing.T) {
	if err := validateField(validateTestForm{}, "Field3", nil); err != errInvalidParameter {
		t.Fatalf("%s != %s", err, errInvalidParameter)
	}
}

func TestValidateInvalidReturnType(t *testing.T) {
	if err := validateField(validateTestForm{}, "Field4", nil); err != errInvalidReturnType {
		t.Fatalf("%s != %s", err, errInvalidReturnType)
	}
}

func TestValidateValidParam(t *testing.T) {
	if err := validateField(validateTestForm{}, "Field5", strValidParam); err != nil {
		t.Fatalf("%s != nil", err)
	}
}

func TestValidateInvalidParam(t *testing.T) {
	if err := validateField(validateTestForm{}, "Field5", strInvalidParam); err != errField5 {
		t.Fatalf("%s != %s", err, errField5)
	}
}

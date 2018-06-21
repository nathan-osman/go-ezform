package ezform

import (
	"errors"
	"reflect"
	"testing"
)

const strParam = "test"

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
	if p != strParam {
		return errField5
	}
	return nil
}

func validateField(v interface{}, fieldName string, params ...interface{}) error {
	var (
		vType    = reflect.TypeOf(v)
		sType, _ = vType.FieldByName(fieldName)
		vVal     = reflect.ValueOf(v)
		fVal     = vVal.FieldByName(fieldName)
	)
	return validate(vType, sType, vVal, fVal, params...)
}

func TestValidateNoValidator(t *testing.T) {
	if err := validateField(validateTestForm{}, "Field1"); err != nil {
		t.Fatalf("%s != nil", err)
	}
}

func TestValidateValidValue(t *testing.T) {
	if err := validateField(validateTestForm{Field2: "test"}, "Field2"); err != nil {
		t.Fatalf("%s != nil", err)
	}
}

func TestValidateInvalidValue(t *testing.T) {
	if err := validateField(validateTestForm{}, "Field2"); err != errField2 {
		t.Fatalf("%s != %s", err, errField2)
	}
}

func TestValidateInvalidParameter(t *testing.T) {
	if err := validateField(validateTestForm{}, "Field3"); err != errInvalidParameter {
		t.Fatalf("%s != %s", err, errInvalidParameter)
	}
}

func TestValidateInvalidReturnType(t *testing.T) {
	if err := validateField(validateTestForm{}, "Field4"); err != errInvalidReturnType {
		t.Fatalf("%s != %s", err, errInvalidReturnType)
	}
}

func TestValidateValidParam(t *testing.T) {
	if err := validateField(validateTestForm{}, "Field5", strParam); err != nil {
		t.Fatalf("%s != nil", err)
	}
}

func TestValidateInvalidParam(t *testing.T) {
	if err := validateField(validateTestForm{}, "Field5", ""); err != errField5 {
		t.Fatalf("%s != %s", err, errField5)
	}
}

func TestValidateMissingParam(t *testing.T) {
	if err := validateField(validateTestForm{}, "Field5"); err != errInvalidParameter {
		t.Fatalf("%s != %s", err, errInvalidParameter)
	}
}

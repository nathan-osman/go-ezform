package ezform

import (
	"errors"
	"reflect"
	"testing"
)

const validValue = "valid"

var errInvalidValue = errors.New("invalid value")

type validateStruct struct {
	Field1, Field2, Field3, Field4 string
}

func (v validateStruct) ValidateField2(s string) error {
	if s != validValue {
		return errInvalidValue
	}
	return nil
}

func (v validateStruct) ValidateField3(i int) error { return nil }
func (v validateStruct) ValidateField4(s string)    {}

func validateField(v interface{}, fieldName string) error {
	var (
		vType    = reflect.TypeOf(v)
		sType, _ = vType.FieldByName(fieldName)
		vVal     = reflect.ValueOf(v)
		fVal     = vVal.FieldByName(fieldName)
	)
	return validate(vType, sType, vVal, fVal)
}

func TestValidateNoValidator(t *testing.T) {
	if err := validateField(validateStruct{}, "Field1"); err != nil {
		t.Fatalf("%s != nil", err)
	}
}

func TestValidateValidValue(t *testing.T) {
	if err := validateField(validateStruct{Field2: validValue}, "Field2"); err != nil {
		t.Fatalf("%s != nil", err)
	}
}

func TestValidateInvalidValue(t *testing.T) {
	if err := validateField(validateStruct{}, "Field2"); err != errInvalidValue {
		t.Fatalf("%s != %s", err, errInvalidValue)
	}
}

func TestValidateInvalidParameter(t *testing.T) {
	if err := validateField(validateStruct{}, "Field3"); err != errInvalidParameter {
		t.Fatalf("%s != %s", err, errInvalidParameter)
	}
}

func TestValidateInvalidReturnType(t *testing.T) {
	if err := validateField(validateStruct{}, "Field4"); err != errInvalidReturnType {
		t.Fatalf("%s != %s", err, errInvalidReturnType)
	}
}

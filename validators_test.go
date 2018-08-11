package ezform

import (
	"fmt"
	"net/url"
	"testing"
)

type testNonEmptyValidatorForm struct {
	Field *StringField
}

func TestNonEmptyValidator(t *testing.T) {
	for _, test := range []struct {
		Input string
		Valid bool
	}{
		{Input: "", Valid: false},
		{Input: "a", Valid: true},
	} {
		var (
			u = url.Values{"Field": []string{test.Input}}
			f = &testNonEmptyValidatorForm{
				Field: NewStringField(&NonEmptyValidator{}),
			}
		)
		v, err := simulateRequest(u, f)
		if err != nil {
			t.Fatal(err)
		}
		if v != test.Valid {
			t.Fatalf("%v != %v", test.Input, test.Valid)
		}
	}
}

type testMinMaxValidatorForm struct {
	Field *IntegerField
}

func TestMinMaxValidator(t *testing.T) {
	for _, test := range []struct {
		Input int64
		Valid bool
	}{
		{Input: -2, Valid: false},
		{Input: -1, Valid: true},
		{Input: 0, Valid: true},
		{Input: 1, Valid: true},
		{Input: 2, Valid: false},
	} {
		var (
			u = url.Values{"Field": []string{fmt.Sprintf("%v", test.Input)}}
			f = &testMinMaxValidatorForm{
				Field: NewIntegerField(&MinMaxValidator{Min: -1, Max: 1}),
			}
		)
		v, err := simulateRequest(u, f)
		if err != nil {
			t.Fatal(err)
		}
		if v != test.Valid {
			t.Fatalf("%v != %v", test.Input, test.Valid)
		}
	}
}

type testRequiredValidatorForm struct {
	Field *BooleanField
}

func TestRequiredValidator(t *testing.T) {
	for _, test := range []struct {
		Input string
		Valid bool
	}{
		{Input: "", Valid: false},
		{Input: "a", Valid: true},
	} {
		var (
			u = url.Values{"Field": []string{test.Input}}
			f = &testRequiredValidatorForm{
				Field: NewBooleanField(&RequiredValidator{}),
			}
		)
		v, err := simulateRequest(u, f)
		if err != nil {
			t.Fatal(err)
		}
		if v != test.Valid {
			t.Fatalf("%v != %v", test.Input, test.Valid)
		}
	}
}

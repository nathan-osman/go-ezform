package fields

import (
	"testing"
	"time"
)

const dateTimeStr = "2009-11-10T23:00:00Z"

var dateTimeVal = time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)

func TestNewDateTime(t *testing.T) {
	for _, test := range []struct {
		Input  string
		Error  error
		String string
		Value  time.Time
	}{
		{Input: "", Error: errInvalidDateTime},
		{Input: dateTimeStr, String: dateTimeStr, Value: dateTimeVal},
	} {
		f := NewDateTime()
		if err := f.Parse(test.Input); err != test.Error {
			t.Fatalf("%v != %v", err, test.Error)
		}
		if test.Error == nil {
			if f.String() != test.String {
				t.Fatalf("%v != %v", f.String(), test.String)
			}
			if !f.Value().Equal(test.Value) {
				t.Fatalf("%v != %v", f.Value(), test.Value)
			}
		}
	}
}

func TestDateTimeSetValue(t *testing.T) {
	f := NewDateTime()
	f.SetValue(dateTimeVal)
	if !f.Value().Equal(dateTimeVal) {
		t.Fatalf("%v != %v", f.Value(), dateTimeVal)
	}
}

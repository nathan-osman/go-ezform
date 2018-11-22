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
		Output time.Time
		Error  error
	}{
		{Input: "", Error: errInvalidDateTime},
		{Input: dateTimeStr, Output: dateTimeVal},
	} {
		f := NewDateTime()
		if err := f.Parse(test.Input); err != test.Error {
			t.Fatalf("%v != %v", err, test.Error)
		}
		if !f.Value().Equal(test.Output) {
			t.Fatalf("%v != %v", f.Value(), dateTimeVal)
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

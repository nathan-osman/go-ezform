package validators

import (
	"testing"
)

func TestMinMax(t *testing.T) {
	for _, test := range []struct {
		Input int64
		Error bool
	}{
		{Input: -2, Error: true},
		{Input: -1, Error: false},
		{Input: 0, Error: false},
		{Input: 1, Error: false},
		{Input: 2, Error: true},
	} {
		var (
			m   = &MinMax{Min: -1, Max: 1}
			err = m.Validate(test.Input)
		)
		switch {
		case test.Error && err == nil:
			t.Fatal("error expected")
		case !test.Error && err != nil:
			t.Fatal(err)
		}
	}
}

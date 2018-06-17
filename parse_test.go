package ezform

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
)

type parseTestForm struct {
	StringVal string
	IntVal    int
	BoolVal   bool
	TimeVal   time.Time
}

var parseTestReference = &parseTestForm{
	StringVal: "test",
	IntVal:    42,
	BoolVal:   true,
	TimeVal:   time.Time{},
}

func toURLVal(v interface{}) []string {
	return []string{fmt.Sprintf("%v", v)}
}

func simulateRequest(v url.Values, f *parseTestForm) (map[string]error, error) {
	r := httptest.NewRequest(
		http.MethodPost,
		"/",
		strings.NewReader(v.Encode()),
	)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return Parse(r, f)
}

func TestStoreValues(t *testing.T) {
	var (
		v = url.Values{
			"StringVal": toURLVal(parseTestReference.StringVal),
			"IntVal":    toURLVal(parseTestReference.IntVal),
			"BoolVal":   toURLVal(parseTestReference.BoolVal),
			"TimeVal":   toURLVal(parseTestReference.TimeVal),
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

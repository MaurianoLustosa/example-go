package parse

import (
	"testing"
)

func TestParseStringIntoFloatSucessufuly(t *testing.T) {

	want := 123456.789
	actual := ParseFloatAndHandleError("123456.789")

	if want != actual {
		t.Fatalf(`ParseFloatAndHandleError("123456.789") = %f, want match for %f`, actual, want)
	}
}

func TestParseStringIntoFloatUnSucessufuly(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf(`ParseFloatAndHandleError("not_float") should panic and didn't`)
		}
	}()

	ParseFloatAndHandleError("not_float")
}

func TestParseStringIntoIntSucessufuly(t *testing.T) {

	want := 123456
	actual := ParseIntAndHandleError("123456")

	if want != actual {
		t.Fatalf(`ParseIntAndHandleError("123456") = %d, want match for %d`, actual, want)
	}
}

func TestParseStringIntoIntUnSucessufuly(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf(`ParseIntAndHandleError("123456.789") should panic and didn't`)
		}
	}()

	ParseIntAndHandleError("123456.789")

}

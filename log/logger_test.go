package log

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestLogCalculatedValues(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer setOutput()
	LogCalculatedValues(123456.78, 11111.00)

	actual := buf.String()
	want1 := "Total amount:  $123456.78"
	want2 := "Installment value:  $11111.00"

	if !strings.Contains(actual, want1) || !strings.Contains(actual, want2) {
		t.Fatalf(`LogCalculatedValues(123456.78, 11111.00) = %q, want match for %q and %q`, actual, want1, want2)
	}
}

func setOutput() {
	log.SetOutput(os.Stderr)
}

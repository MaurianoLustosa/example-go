package math

import (
	"testing"
)

func TestCalculateInstallmentValue(t *testing.T) {

	want := 34.29
	actual := CalculateInstallmentValue(1234.567, 36)

	if want != actual {
		t.Fatalf(`CalculateInstallmentValue(1234.567, 36) = %f, want match for %#f`, actual, want)
	}
}

func TestCalculateTotalValue(t *testing.T) {

	want := 2758.63
	actual := CalculateTotalDueAmount(1234.567, 123.45, 12)

	if want != actual {
		t.Fatalf(`CalculateTotalDueAmount(1234.567, 123.45, 12) = %f, want match for %#f`, actual, want)
	}
}

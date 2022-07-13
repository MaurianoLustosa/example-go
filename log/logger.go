package log

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func LogCalculatedValues(totalAmount float64, installmentValue float64) {
	log.Println("Total amount: ", formatMonetaryValue(totalAmount))
	log.Println("Installment value: ", formatMonetaryValue(installmentValue))
}

func formatMonetaryValue(value float64) string {
	return fmt.Sprintf("$%.2f", value)
}

func WriteString(w http.ResponseWriter, s string) (n int, err error) {
	return io.WriteString(w, s)
}

func Println(v ...any) {
	log.Println(v...)
}

package parse

import (
	"fmt"
	"strconv"
)

func ParseFloatAndHandleError(number string) float64 {
	var result, err = strconv.ParseFloat(number, 64)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

func ParseIntAndHandleError(number string) int {
	var result, err = strconv.ParseInt(number, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	return int(result)
}

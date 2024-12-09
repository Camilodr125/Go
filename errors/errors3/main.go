package main

import (
	"fmt"
)

type divideError struct {
	dividend float64
}

func (e divideError) Error() string {
	mess:=  fmt.Sprintf("can not divide %.0f by zero", e.dividend)
	return mess
}


func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

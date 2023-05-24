package main

import (
	"errors"
	"fmt"
)

func main() {
	result, remainder, err := div(5, 2)
	fmt.Println("5/2 =", "(result=", result, ", remainder=", remainder, ", err=", err)
}

func div(numerator int, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")

		return result, remainder, err
	}
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, err
}

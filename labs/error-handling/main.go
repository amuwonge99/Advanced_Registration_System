package main

import (
	"errors"
	"fmt"
)

var (
	ErrDivByZero = errors.New("attempt to divide by zero")
)

func main() {
	// A list of example inputs to try dividing
	divisionExamples := []struct {
		dividend int
		divisor  int
	}{
		{10, 5},       // valid division
		{4, 0},        // division by zero
		{-8, 2},       // negative number
		{23, 7},       // prime number
		{10000000, 2}, // input too large
		{37, 1},       // divide by 1
	}

	// Loop through each example and attempt the division
	for _, input := range divisionExamples {
		result, err := Divide(input.dividend, input.divisor)
		switch err {
		case ErrDivByZero:
			fmt.Printf("unable to divide %d by %d, error: %s\n", input.dividend, input.divisor, err)

		default:
			fmt.Printf("%d divided by %d is %d\n", input.dividend, input.divisor, result)
		}
	}
}

// Divide performs integer division and returns an error if the input is invalid
func Divide(numerator, denominator int) (int, error) {
	// Check for division by zero
	if denominator == 0 {
		return 0, ErrDivByZero
	}

	// Perform the division
	return numerator / denominator, nil
}

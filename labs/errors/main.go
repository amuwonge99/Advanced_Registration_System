package main

import (
	"errors"
	"fmt"
)

func main() {
	println("program start")
	result := DoSomeMaths(1, 3, 0)

	fmt.Println("result:", result)
}

func DoSomeMaths(value1, value2, value3 int) (result int) {
	resultOfSubtract := value1 - value2

	result, _ = Divide(resultOfSubtract, value3)

	return result
}

func Divide(numerator, denominator int) (result int, err error) {
	if denominator == 0 {
		return 0, errors.New("can't divide by zero")
	}
	return numerator / denominator, nil
}

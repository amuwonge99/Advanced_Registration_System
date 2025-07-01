package main

import (
	"fmt"
	"strings"
)

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("negative input")
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result, nil
}

func IsPalindrome(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func ToUpperCase(s string) string {
	return strings.ToUpper(s)
}

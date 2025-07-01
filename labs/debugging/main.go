package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum of numbers:", sum(numbers))
}

func sum(nums []int) int {
	total := 0
	for i := 0; i <= len(nums); i++ {
		total += nums[i]
	}
	return total
}

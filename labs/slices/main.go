package main

import "fmt"

func main() {
	var ages = []int{1, 8, 7, 2, 9}
	fmt.Println(ages)

	var first_three_ages = ages[:3]
	fmt.Println(first_three_ages)

	var last_three_ages = ages[2:]
	fmt.Println(last_three_ages)
}

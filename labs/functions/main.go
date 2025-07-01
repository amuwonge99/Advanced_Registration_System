package main

func main() {

	// We can create a couple of slices the same way we have seen before
	myFirstSlice := []string{"Daniel", "Rupert", "Emma"}
	mySecondSlice := []string{"Radcliffe", "Grint", "Watson"}

	// if we want to print them out we can create a function that loops through
	// and prints for us rather then duplicate code
	printStringSlice(myFirstSlice)
	printStringSlice(mySecondSlice)

	// lab part
	result := add(20, 30)
	println(result) // 50
	result, remainder := divide(10, 3)
	println(result)    // 3
	println(remainder) // 1
}

// printStringSlice is a function that takes in an arguement of a string slice
// and prints each element on a new line
func printStringSlice(slice []string) {
	for _, elem := range slice {
		println(elem)
	}
}

// add will add the two numbers in the arguements and return the result
func add(a int, b int) int {
	// as we say we are returning an int, we MUST return one. For now while we
	// write the code we leave it as 0 to ensure we are returning
	// the correct types
	return 0
}

// divide will divide a by b and return the result and remainder as 2 ints
func divide(a int, b int) (result int, remainder int) {
	// as we say we are returning multiple int, we MUST return them. For now
	// while we write the code we leave it as 0,0 to ensure we are returning
	// the correct types
	return 0, 0
}

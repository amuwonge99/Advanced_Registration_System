package main

func main() {
	// very universal loop here, notice we use the shorthand for creating the variable i and assigning it to 0
	for i := 0; i < 10; i++ {
		print(i)
	}
	println()

	// we also have the option of creating an array (or slice in Go)
	// notice that for a slice is in 3 parts:
	// []     -- this signifies that it is a slice or array
	// string -- this is the type that is held in the slice, a slice can only hold one type of data
	// {...}  -- this is the inital data that we want to store inside the slice
	mySlice := []string{"Hello", "and", "welcome", "to"}

	// slices, unlike arrays, are not fixed size and can grow on demand using the built-in append keyword
	mySlice = append(mySlice, "Go!")

	// we can loop through a slice using a for loop with the range keyword
	// when we use the range keyword we get 2 values:
	// 1: The iteration number of the slice
	// 2: the value of that iteration in the slice
	for iteration, value := range mySlice {
		print("element: ")
		print(iteration)
		print(", " + mySlice[iteration] + " == " + value + "\n")
	}
}

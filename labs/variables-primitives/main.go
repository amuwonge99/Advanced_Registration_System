package main

func main() {
	// here we use the var keyword to specify we are creating a variable, we then specify the name and type
	var myNumber int
	// after we have created a variable we can assign it values
	myNumber = 5
	println(myNumber)
	// we can reassign values the same way
	myNumber = 10
	println(myNumber)

	// if we already know the initial value for a variable then we can use the := shorthand to create and assign it
	myString := "Hello World!"
	println(myString)
	// we can reassign values the same way as before
	myString = "!dlroW olleH"
	println(myString)
}

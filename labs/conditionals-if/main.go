package main

func main() {
	// This will run as the condition is true
	if true {
		println("I will print!")
	}

	// This will not run as the condition is false
	if false {
		println("I will not print!")
	}

	// we can show this as a boolean variable
	shouldIRun := true
	if shouldIRun {
		println("I will print!")
	}
	shouldIRun = false
	if shouldIRun {
		println("I will not print!")
	}

	// we can also compare variables to values or other variables
	colour := "red"

	// in this if statement we are using a boolean comparator == which checks
	// if the left side is the same as the right side
	if colour == "red" {
		println("red!")
	}

	// as colour is red, it does not equal to blue so the following print
	// will not run
	if colour == "blue" {
		println("blue!")
	}

	colour = "blue"
	// as colour is now blue, it does equal to blue so the following print
	// will run
	if colour == "blue" {
		println("blue!")
	}

	// we can also check that things are not equal using the != comparator
	// in this example remember that colour is currently set to "blue"
	if colour != "red" {
		println("colour is not red!")
		print("colour is ")
		print(colour)
		println()
	}

	// We have access to the else statement, which means that when the `if`
	// fails then the else runs instead
	vehicle := "car"
	if vehicle == "bike" {
		println("I am riding a bike!")
	} else {
		println("I am not riding a bike!")
		print("I am riding a ")
		print(vehicle)
		println()
	}

}

package main

// this struct has been incorrectly defined, fix this
type Person struct {
	FirstName string
	Age
}

// do not change any code past this point, but do review it to see how to fix the struct
func main() {

	myPerson := Person{
		FirstName: "John",
		LastName:  "Doe",
	}
	myPerson.Age = 30

	println("Hello I am", myPerson.FirstName, myPerson.LastName, "and I am", myPerson.Age, "years old")

}

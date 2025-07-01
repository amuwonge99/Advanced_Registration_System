package greet

import "fmt"

// lowercase so cannot be referenced outside the package
var defaultPerson = "steve"

// public function
func Hello() {
	fmt.Println("Hello, World!")
}

// public function, can call a private function within itself
func Goodbye(person string) {
	if person != "" {
		fmt.Println("So long, farewell! We will see you soon", person)
	} else {
		goodbyeDefaultPerson()
	}

}

// private function
func goodbyeDefaultPerson() {
	fmt.Println("Goodbye sweet sweet", defaultPerson)
}

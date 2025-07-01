package main

// we import the fmt package so we can use any public functions within it
import (
	"fmt"
	"github.com/google/uuid"

	"gitlab.platform-engineering.com/golang-academy/lab-imports/greet"
)

func main() {

	fmt.Println("good morning")

	id := uuid.New().String()

	fmt.Println(id)

	greet.Hello()

	greet.Goodbye("sandra")
	greet.Goodbye("")

}

package main

// we import the fmt package so we can use any public functions within it
import (
	"fmt"

	"gitlab.platform-engineering.com/golang-academy/lab-methods/vehicle"
)

func main() {
	// NewCar() is public so we can call it from our package
	myCar := vehicle.NewCar(4)

	myBike := vehicle.NewBike()

	fmt.Println("my car has", myCar.GetWheels(), "wheels")
	fmt.Println("my bike has", myBike.GetWheels(), "wheels")

	myCar.AddWheels(2)

	fmt.Println("my car has", myCar.GetWheels(), "wheels")
}

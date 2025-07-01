package main

// we import the fmt package so we can use any public functions within it
import (
	"fmt"

	"gitlab.platform-engineering.com/golang-academy/lab-interfaces/vehicle"
)

func main() {
	// we can create our different vehicle types
	myCar := vehicle.NewCar(4, "red", "ferrari", 100000)

	myBike := vehicle.NewBike("matte black")

	vehiclesIOwn := []vehicle.Vehicle{&myCar, &myBike}

	for _, v := range vehiclesIOwn {
		fmt.Println(vehicle.TalkAbout(v))
	}

	fmt.Println(vehicle.Advertise(&myCar))
}

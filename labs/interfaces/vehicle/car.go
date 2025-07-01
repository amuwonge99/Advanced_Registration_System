package vehicle

import "fmt"

// Car is public so can be called by other packages
type Car struct {
	// this is lower case so it is private
	// anyone that imports this package cannot directly
	// set the wheel field
	wheels int
	colour string
	make   string
	value  int
}

// NewCar is public so we can use this to create a Car variable
// Cars could have a different number of wheels so lets take that in as an arguement
func NewCar(wheels int, colour string, make string, value int) Car {
	return Car{
		wheels: wheels,
		colour: colour,
		make:   make,
		value:  value,
	}
}

// GetWheels takes your existing object and returns the number of wheels it has
func (c *Car) GetWheels() int {
	return c.wheels
}

func (c *Car) Description() string {
	return fmt.Sprintf("a %v %v", c.colour, c.make)
}

func (c *Car) Price() int {
	return c.value
}

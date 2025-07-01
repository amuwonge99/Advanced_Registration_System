package vehicle

// Bike is also public so can be called by other packages
type Bike struct {
	// this is upper case so it is public
	// anyone that imports this package can directly
	// set the wheel field
	Wheels int
}

// NewBike is public but calls to the private
func NewBike() Bike {
	// we know a bike can only have 2 wheels so this is used as a constructor to ensure all the values are set correctly
	return newBike(2)
}

func newBike(wheels int) Bike {
	return Bike{
		Wheels: wheels,
	}
}

// GetWheels takes your existing object and returns the number of wheels it has
func (b *Bike) GetWheels() int {
	return b.Wheels
}

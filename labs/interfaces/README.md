INTERFACES
==============

Interfaces are named collections of methods. To use them we define a type as an interface and populate it with methods.
The interface then recognises any type with a method matching the one defined in the interface as part of that interface.
For example:

```go
type Vehicle interface {
	GetWheels() int
	Description() string
}
```

This `Vehicle` interface will be satisfied by any defined types that have methods `GetWheels` and `Description` that return an int value and a string value respectively.

We can use this to easily group together types that share these methods

```go
	myCar := vehicle.NewCar(4, "red", "ferrari", 100000)

	myBike := vehicle.NewBike("matte black")

	vehiclesIOwn := []vehicle.Vehicle{&myCar, &myBike}

    for _, v := range vehiclesIOwn {
		fmt.Println(vehicle.TalkAbout(v))
	}

```

In this example, we are using the 2 structs from earlier, Car and Bike, defining additional methods so that the satisfy the vehicle interface, this allows us to append both of the populated struct objects myCar and myBike to the slice of vehiclesIOwn, we can then easily call the TalkAbout function on all the vehicles I own by simply ranging through the slice. If a variable passed to a function is of type interface, then we can call all the methods that are in the named interface within the function. This allows the TalkAbout function to easily access the description and number of wheels of our vehicles.

LAB TASK
--------

Write a type that satisfies the `Sale` interface but not the `Vehicle` interface. You should be able to create a list of items for sale including myCar and an object of your new type, and range through that list to call the Advertise function on each object.

Stretch Goal:
Implement your own interface satisfied by your new type alone. (hint: you should start by defining another Method on your type).
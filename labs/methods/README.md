METHODS
---
A method is a function with a special receiver argument. The receiver binds the function to a type and allows the method to access the properties of the receiver type. This allows Go to utilise OOP or object orientated programming, something that usually relies on the concept of classes and objects. Using structs and methods we can create an object and then have specific functions that are tied to that instance of the object. Receivers do NOT have to be structs. They could be other types such as int, string, etc.
```go
myCar := NewCar(4)       // create a car with 4 wheels
mySecondCar := NewCar(5) // create a car with 5 wheels
fmt.Println("my first car has %d wheels", myCar.GetWheels())
fmt.Println("my second car has %d wheels", mySecondCar.GetWheels())
```

In the example above we created two cars each with a different number of wheels. We can then call a function on the object itself rather than calling a function with the car object as an argument. To do this we add in the object almost like an argument in the function definition but before we provide the name, like so:

```go
func (c *Car) GetWheels() int {
	return c.wheels
}
```

Because we have called it this way the function has access to the private fields within the object (as this function was defined in the same package as the struct).

You can see we also have a method called `AddWheels` defined against our car struct. But when the code is adding wheels to our car, the second println does not reflect the change, why do we think this is?

LAB TASK
--------

So far we have been able to create cars and bikes and see what their numbers of wheels are using the `GetWheels()` function. Extend the code to include a `SetWheels` function that takes in an integer and sets the amount of wheels that car or bike has to the provided amount. Then extend the main code to print out the new value.
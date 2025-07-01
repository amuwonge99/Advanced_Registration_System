ERRORS
======

We don't really want to panic when we find an issue, so instead we should use an [error](https://go.dev/doc/tutorial/handle-errors) to tell us something is wrong.

In Go, errors are a type like anything else, so when we call a function that could fail for some reason, we generally return the thing that we want as well as an error. For example, we saw we panicked when we tried to divide by zero in a previous lab.

To get around this, if we know there is a potential for something to go wrong, then we should also return an error. For example:

```go
func Divide(numerator, denominator int) (result int, err error) {
	if denominator == 0 {
		return 0, errors.New("can't divide by zero")
	}
	return numerator / denominator, nil
}
```

We now always get back two values from when we call the `Divide` function: a `result` of type `int` and an `err` of type `error`. So to call it, we could do the following to capture both of them (be aware we have to let Go know that we are getting all of the returned values):

SIDE-NOTE: due to how the error type has been defined, it can hold a couple of different values, with one of these values being nil (nothing) and the other being an error, of which the error holds the message string.

```go
numerator := 10
denominator := 5
result, err := Divide(numerator, denominator)
```

Okay, so now we have a result and an error, but how do we know if there was a problem? We do that by checking if the `err` is `nil` (Go's version of [null](https://en.wikipedia.org/wiki/Null_pointer)). If it is nil then no error was thrown and we are good! If it is not nil then there is a problem and we can't carry on. We can print out errors as a string to see what the problem was.

```go
if err != nil {
    fmt.Printf("unable to divide %d by %d, err: %s\n", numerator, denominator, err)
}
```

LAB TASK
--------

The `DoSomeMaths()` function first subtracts `value1` from `value2`, then divides `value3` from the result. The `Divide` function in the Lab appears to be throwing an error, but the error isn't returned to the main function.

Update the `DoSomeMaths()` function to return the error received from the `Divide()` function.

STRETCH TASK
------------

Update the main function to return the result only if `DoSomeMaths()` was successful.

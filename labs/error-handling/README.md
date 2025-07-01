CUSTOM ERROR HANDLING
=====================

We saw with our divide operation in a previous lab that we could send back an error if something went wrong. But what about if multiple things could go wrong and some of those we could handle? In that case, what if we checked what the error was in more detail and then saw if we could fix this going forward.

So, let’s extend our `Divide()` function with a few more errors. We know it’s bad to divide by 0, but what’s the point of dividing by 1 as you will get the same number back? What about if the number doesn’t neatly divide? Let’s start making custom errors for each of these. Let’s create ourselves some named errors globally:

```go
var (
	ErrDivByZero      = errors.New("attempt to divide by zero")
	ErrDivByOne       = errors.New("are you sure you want to divide by one")
	ErrRemainderFound = errors.New("there was a remainder")
)
```

We now have three errors: divide by zero, divide by one, and remainder found.

We can pass them around by calling them directly, for example to return an error that says you tried to divide by zero, you can do:

```go
return ErrDivByZero
```

We can check in detail the type of our error by comparing it to our existing ones. We can do this via if or a switch statement:

```go
switch err {
case ErrDivByZero:
	...
case ErrDivByOne:
	...
case ErrRemainderFound:
	...
default:
	doSomething()
}
```

Switch statements can make things easier to read but they are essentially the same as if-else chains. The above is exactly the same as:

```go
if err == ErrDivByZero {
	...
} else if err == ErrDivByOne {
	...
} else if err == ErrRemainderFound {
	...
} else {
	doSomething()
}
```

LAB TASK
--------

Other developers have been using the Divide function but they have asked for improvements to the error messages. You have been asked to improve the returned error messages in the Divide function so that it provides extra **validation** and returns helpful error messages. The following is a list of improvements:

### Task 1: Check for negative numbers
Return an error if either the dividend or divisor is negative.  
Example error message: "negative numbers are not allowed"

### Task 2: Check for remainders
Return an error if the division leaves a remainder.  
Example error message: "there was a remainder"

### Task 3: Check for large inputs
Return an error if either the dividend or divisor is greater than or equal to 1,000,000.  
Example error message: "input values are too large"

### Task 4: Check for divide by 1
Return a warning (as an error) if the divisor is 1.  
Example error message: "are you sure you want to divide by one"

The output should look something like the following:

	10 divided by 5 is 2
	unable to divide 4 by 0, error: attempt to divide by zero
	Negative input detected: -8 or 2. Only positive numbers are allowed, error: negative numbers are not allowed
	23 is not completely divisible by 7, got 3.
	Input too large: 10000000 or 2 exceeds the allowed limit, error: input values are too large
	able to divide 37 by 1 and got 37, but should we have?, error: are you sure you want to divide by one

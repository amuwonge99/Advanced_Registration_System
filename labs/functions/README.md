FUNCTIONS
=========

[Functions](https://go.dev/tour/basics/4) solve two problems in coding. The first is readability: we can put like-code together and give it a logical name to make it clearer what our code is doing while reviewing it. The other is [avoiding duplication](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself), if we have to do something several times in the same way we can put that in a function and call the function to save us space and ensuring if that process needs to change, we can change it once.

SYNTAX OF A FUNCTION
--------------------

A function will always start the same way, with the `func` keyword. We have seen this already with the main function. This is the most basic of functions as it takes no arguments and returns nothing. When we create a function we have the option of adding in arguments and returns. For example if we had the function with the following arguments:

```go
func reverseString(myString string) string
```

This tells us we have a function called `reverseString`, it takes one argument that is a `string` and is called `myString` within the function. It will also return a string at the end of the function.

We can have multiple arguments in a function, for example a function that adds two numbers together and then returns the result could look something like this:

```go
func add(a int, b int) int
```

We can also have multiple returns in Go. One example is if we wanted to divide two integers together and get back the result and the remainder, we can do this as the following function:

```go
func divide(a int, b int) (result int, remainder int)
```

If we have multiple returns, when we call the function we can capture all of them using a comma separated list. For example:

```go
result, remainder := divide(10, 3)
println(result)    // 3
println(remainder) // 1
```

LAB TASK
--------

In the code for this lab we have created the functions for add and divide but we haven't filled out any of the logic for this. Update the `add` and `divide` functions to ensure they work as we expect.

Hint: for the divide function you will need the [`%` operator](https://Gobyexample.com/remainder-modulus-go-Go/).

<!-- markdownlint-disable-file MD013 -->

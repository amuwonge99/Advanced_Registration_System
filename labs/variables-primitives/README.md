VARIABLES - Primitives
=========

In this lab, we will be exploring variables and how to use them in a [statically compiled language](https://stackoverflow.com/a/12600515).

CREATING VARIABLES
------------------

There are two ways of creating variables within Go. We have the basic form using the [`var` keyword](https://go.dev/tour/basics/8).

```go
var age int
var name string
```

As Go is statically typed, when we create a variable, we provide it with a type that remains permanent.

In the example above, we create age as an integer. Because we used the `var` keyword, it now has the default value for an `int` (integer) which is `0`. There is no way we can change the type of `age` after it has been created, it will always be of type `int`. 

We can also create and assign values to variables at the same time using [type inference](https://go.dev/tour/basics/14). Go determines the type based on the value provided.

```go
age := 27
name := "Damon"
```

We have typed a variable name that doesn't exist and we have used the `:=` assigner, so Go will take the value on the right-hand side (RHS) and determine the type of that value and create a variable with that type on the left (LHS) and assign the value.

In this case `age` is still an `int` and `name` has been set as a `string`. 

NAMING VARIABLES
----------------

The way we name variables and functions in Go is a functional part of Go's syntax and conventions. Using a capital letter at the beginning of a variable name indicates that it is a globally accessible variable, but we will delve into more detail later. For now, variables should be named using [lowerCamelCase](https://en.wikipedia.org/wiki/Camel_case). This means that if a variable has multiple words in its name, the first word should be lowercase, and each subsequent word should start with an uppercase letter.

```go
var firstName string
var lastName string

var ageInYears int
var ageInMonths int

var aVeryLongVariableNameWithMultipleWordsInIt bool
```

LAB TASK
--------

Below you can see an example program. 

```go
package main

func main() {
	var myString string
	println(myString)
}
```

Extend this program so that it outputs the following output when ran

```
Hello World!
1234
Welcome To Go!
5678
```

The solution **should** use multiple `println(myString)` statements and **should not** create any other variables. 

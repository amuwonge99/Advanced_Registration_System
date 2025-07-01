
VARIABLES - BASIC MATH OPERATIONS
---------------------------------

Variables can also be used to store the results of calculations, for example you can add the values stored in two or more variables using the + operator.
```go
result = firstVariable + secondVariable
```
this approach can be used with other mathematical operators:
 + for addition
 - for subtraction
 * for multiplication
 / for division
 ( ) for grouping expressions

You can also mix variables with literal numbers, the following examples show some possibilites.

```go
result = firstVariable - 5
result = firstVariable + 5 + secondVariable + 2 
result = firstVariable * 15 + 60 / 2
```
But remember just like in your school maths class, order of operations is important (see order of operators https://en.wikipedia.org/wiki/Order_of_operations)


INTEGER DIVISION
----------------
In the Go programming language dividing two integers will always result in an interger result with the remainder discarded, in other words you can't mix and match variables of different types (eg integers and floats) without some extra steps, this also means we can't add an integer variable to a floating point number variable.

```go
first := 5
second := 2
result := first / second 
println(result) // result = 2
```

If you want the decimals you have two options you can cast the interger to a more appropriate type, a float discussed in seperate lession is a good choice or you can cheat and turn the decimal to a float when assigning the variable like in the below example:
```go
first := 5.0
second := 2.0
result := first / second
println(result)
```
Be warned while this will only work for simple calculations.


VARIABLES - STRING OPERATORS
-----------------------------
As Go is a statically typed language the string data type only works with the + operator this is used to concatenate or join strings together:
```go
username := "coder123"
greeting := "Welcome back, " + username
println(greeting)
```
While you can join strings as shown in the above examples, its important to note:
 * you cant use any other math operators -, * and / will cause compile time errors
 * to join a string and a number together you must first convert the number to a string

for example, this will not compile
```go
usersOnline := 5
greeting := "there are " + usersOnline + " online" // compile error
println(greeting)
```

so instead we convert the variable usersOnline to a string using the strconv package:
```go
import "strconv"

usersOnline := 5
greeting := "there are " + strconv.Itoa(usersOnline) + " online"
println(greeting)
```

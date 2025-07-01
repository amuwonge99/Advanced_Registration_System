CONDITIONALS
============

[Conditionals](https://www.codecademy.com/learn/learn-go/modules/learn-go-conditionals/cheatsheet) are the way in which decisions are made in code. Up to this point, we haven't had to make any decisions in our code, we have just gone step by step. The following lab will cover the `if` conditional and it's usage.

IF STATEMENT
------------

The most famous conditional is the [`if` statement](https://go.dev/tour/flowcontrol/6), which always boils down to a [Boolean](https://en.wikipedia.org/wiki/Boolean_expression) statement, e.g. is something true or false. If the outcome of the condition is true, then the code within the `if` will be executed. If it is false, then the code will be skipped, and processing will move on to the next section of code.

In general, we use the `if` keyword to compare if something is the same as something else, this is true in a lot of cases but not always. As such, there are a set of comparison operators in Go, you can see the most common ones in the table below:

| Operator | Name                     | Example  | Result                                      |
| :--      | :--                      | :--      | :--                                         |
| `==`     | Equal                    | `x == y` | True if `x` is equal to `y`                 |
| `!=`     | Not equal                | `x != y` | True if `x` is not equal to `y`             |
| `<`      | Less than                | `x < y`  | True if `x` is less than `y`                |
| `<=`     | Less than or equal to    | `x <= y` | True if `x` is less than or equal to `y`    |
| `>`      | Greater than             | `x > y`  | True if `x` is greater than `y`             |
| `>=`     | Greater than or equal to | `x >= y` | True if `x` is greater than or equal to `y` |

There is another conditional known as the [switch/case](https://gobyexample.com/switch) statement, but we will not be using that in this lab. 

LAB TASKS
--------

Write a program which:

1. Asks for the user's age as input.
2. Stores the user input into a variable called `myAge`.
3. Determines the user's birth year.
4. Then outputs one of the following messages:

    * If they were born **before** 2000, print out: "You were born in the 20th Century"
    * If they were born **after** 2000, print out: "You were born in the 21st Century"
    * If they were born **in** the year 2000, print out: "You were born in the millennium!"

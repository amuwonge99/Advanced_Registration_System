PANIC!
======
Before we talk about panics, I want to cover a bit about division by zero. In basic maths, dividing by zero is undefined – you can't divide a cake that has already been eaten. Some computers solved this by always returning zero, while others threw an error. As a result, some older programming languages left this undefined, which could cause programs to stop working unexpectedly. As languages and computers became more advanced, they started to catch and define these errors better, leading to the divide-by-zero error becoming somewhat famous in tech circles.

Go, being a modern language, handles this in a 'fail fast' approach. To do this, Go detects division by zero at compile time and will refuse to compile the code. Then, if the code manages to compile cleanly and division by zero happens during execution, Go will produce a panic. [Panic](https://gobyexample.com/panic) is where our code hits a situation where it just cannot continue. It’s an “I don’t know how to carry on, so instead I’ll just tap out” action. Here is some code to demonstrate this (you may notice we had to put the 0 in a variable because Go would refuse to compile if we explicitly put a divide-by-zero operation in our code):

```go
package main

func main() {
	divisor := 0
	total := 5 / divisor
	println(total)
}
```

When I try to run it, my code panics and I get this as a response:

```sh
$ go run .
panic: runtime error: integer divide by zero

goroutine 1 [running]:
main.main()
        /Users/damonwright/go/src/gitlab.platform-engineering.com/golang-academy/lab-panic/main.go:5 +0x11
exit status 2
```

What does this tell us? Let's break it down into sections:

`panic: runtime error: integer divide by zero` We are told our code panicked due to a runtime error that was caused by an integer divide by zero.

`goroutine 1 [running]:` This told us it ran in "goroutine 1", which means the main logic and not a separate goroutine (like a thread, but we’ll get to those later).

`main.main()` The panic happened in the `main()` function in the `main` package.

`/Users/damonwright/go/src/gitlab.platform-engineering.com/golang-academy/lab-panic/main.go:5 +0x11` The panic happened on line 5 of our code in the file called `main.go`; the `+0x11` shows us the location in memory when the error occurred and is only used when performing low-level debugging.

`exit status 2` A problem occurred with our program and it sent back an exit number `2` to inform the user that something went wrong.

WHAT’S THE POINT?
----------------

Panicking is there if we know there is a situation where we cannot continue, so we may as well exit now rather than continuing.

Best practice says our code should never really crash and we should handle errors, but that’s something else we’ll get onto shortly. We can also make our code panic whenever we like by calling the built-in `panic()` function.

SIDE-NOTE: Most uses in the real world tend to revolve around informing the programmer that they have done something incorrect, whether that’s an incorrect parameter passed to a function or unexpectedly hitting the end of an array. It is also valid to use a panic with microservices, especially during startup, as you will often want this to fail early and fail hard.

RECOVERY
--------

Up until this point, you have been led to believe that panics are final, and I have to confess you’ve been misled slightly. The truth of the matter is that panics can be captured and also ignored.

"How do we do this?", I hear you ask.

Now that’s out of the way, we can go back to the `recover()` function. To use `recover()`, you will need to defer it so that it runs after the panic has happened, that way `recover()` has something to recover from.

The following code sample shows how to use `recover()` to return the captured panic message to the user.

```go
func somefunction() {
	defer recoverFromPanic()
	println("start panic")
	panic("panic here")

}

func recoverFromPanic() {
	err := recover()
	if err != nil {
		fmt.Println(err)
	}
}
```

Important:
*   `recover` won’t work in the `main` function.
*   `defer` makes the function run after everything else.

LAB TASK
--------

Run the code that is part of this lab and have a look at the output.

Make sure you understand all the information that it is telling you and how you can trace through the output of the panic to see the route the code used to get to the panic.

STRETCH TASK
------------

Can you adjust the lab code to handle the panic so it returns the panic message, then continues to print the line `I will never get printed`?

HINT: Use the recovery function block above to recover from where the panic happens.


# Debugging

## Learning Objectives
By the end of this session, students will be able to:

- Understand the concept and importance of debugging.
- Identify common bugs in Go programs.
- Use fmt.Println() for basic debugging.
- Use the delve debugger (dlv) to step through code and inspect variables.

## What is Debugging?
Debugging is the process of identifying, isolating, and fixing bugs or errors in a program's code. It’s a crucial step in software development because it ensures that the code behaves as expected. Effective debugging improves the correctness, reliability, and maintainability of software, making it easier to work with and less prone to failure.

## Common Bugs in Go
When working with the Go programming language, developers often encounter a few common types of bugs:

- Nil pointer dereference: This happens when the code tries to access a memory location through a pointer that hasn’t been initialized.
- Off-by-one errors: These occur when loops or indexing go one step too far or not far enough, often due to incorrect boundary conditions.
- Infinite loops: These are loops that never terminate, usually because the exit condition is never met.
- Incorrect use of slices or maps: Mismanaging slices or maps can lead to unexpected behavior or runtime errors.
- Misuse of goroutines or channels: Improper handling of concurrency features like goroutines and channels can cause deadlocks, race conditions, or other synchronization issues.

## Debugging with fmt.Println() 
One of the simplest ways to debug Go code is by using the fmt.Println() function to print out variable values and program flow.

```go
func add(a int, b int) int {
    result := a + b
    fmt.Println("Debug: a =", a, "b =", b, "result =", result) // Debugging line
    return result
}

func main() {
    sum := add(3, 5)
    fmt.Println("Sum is:", sum)
}
```

Pros: It’s very easy to use and doesn’t require any setup or special tools.
Cons: While helpful for small or simple issues, this method doesn’t scale well for complex debugging tasks, especially in large codebases or concurrent programs.

## Introduction to Delve (dlv) 
Delve is a powerful debugger specifically designed for Go. It allows developers to inspect and control the execution of their programs in a more structured way than print statements.

You can find Delve in your IDE already under the 'Run and Debug' menu.

## Lab - Debugging a Faulty Go Program

### Objective
Students will debug a Go program with a logic error using both fmt.Println() and the delve debugger.

### Instructions

1. Identify the Bug
2. Run the program and observe the panic.
3. Use fmt.Println() to print the index and value inside the loop.
4. Fix the Bug
5. Correct the loop condition.
6. Use the Debugger to Debug
7. Set a breakpoint at the start of the sum function.
8. Step through the loop and inspect i and nums[i].

### Expected Outcome
Students understand how to trace and fix an off-by-one error.
They gain hands-on experience with both basic and advanced debugging tools.
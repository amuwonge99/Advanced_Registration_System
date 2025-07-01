# Unit Tests

## Learning Objectives
By the end of this session, students will be able to:

- Understand the purpose and benefits of unit testing.
- Use Go’s built-in testing package.
- Write and run basic unit tests.
- Interpret test results and debug failing tests.

## What is Unit Testing?
Let’s start with the basics. Unit testing is a software testing method where individual units or components of a program are tested in isolation. A "unit" typically refers to the smallest testable part of an application—often a single function or method.

The goal of unit testing is to validate that each unit of the software performs as expected. Think of it as checking each brick in a wall before you build the whole structure. If each brick is solid, the wall is more likely to be strong.

## Why is unit testing important?

It helps catch bugs early in development.
It makes code easier to refactor and maintain.
It serves as documentation for how your code is supposed to behave.
It builds confidence when deploying changes.

## Go’s Built-in Testing Package
Go makes unit testing straightforward with its built-in testing package. You don’t need to install anything extra—just write your tests in a file that ends with _test.go, and Go will recognize it as a test file.

Here’s how a basic test function looks:

```go
import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}
```

Let’s break this down:

- We import thest `testing` package
- The function name must start with Test and take a single parameter of type `*testing.T`.
- Inside the function, you call the code you want to test.
- If the result isn’t what you expect, you use `t.Errorf` to report a failure.

The `testing.T` object provides methods like:

- `t.Error` and `t.Errorf` to report failures.
- `t.FailNow()` to stop the test immediately.
- `t.Log` and `t.Logf` to print helpful debugging information.

## Running Tests with go test
Once you’ve written your test functions, running them is simple. Just open your terminal, navigate to the directory containing your Go files, and run:

Go will automatically find all files ending in _test.go, run the test functions, and report the results.

You’ll see output like:

```txt
PASS
ok  	example.com/myproject	0.002s
```

If a test fails, Go will show you which test failed and why. This makes it easy to pinpoint issues and fix them quickly.

You can also run tests with additional flags:

```zsh
go test -v # For verbose output (shows each test name)
go test -cover # To see test coverage
go test -run # TestAdd to run a specific test
```

## Lab

### Task Instructions
1. Create a new file utils_test.go in the same package.
2. Import the testing package.
3. Write unit tests for Add and Subtract.
4. Run the test suite using `go test`.
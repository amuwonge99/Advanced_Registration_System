# Table-Driven Tests and Error Handling in Go

## What Are Table-Driven Tests?
In Go, a table-driven test is a way to write multiple test cases in a compact and maintainable format. Instead of writing separate test functions or blocks for each case, we define a "table" (usually a slice of structs) where each entry represents a test case. We then loop over the table and run the same logic for each case.

This approach is idiomatic in Go and helps reduce repetition, especially when testing functions with many input/output combinations.

## Benefits of Table-Driven Tests
* Clarity: All test cases are visible in one place.
* Scalability: Easy to add new test cases.
* Maintainability: Less duplicated code.
* Consistency: Encourages a uniform testing style.

## Example: Testing a String Reversal Function
Let’s say we have a function `Reverse(s string)` string that reverses a string. Here’s how we might write a table-driven test for it.

```go
func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes    }
    return string(runes)
}
```

You might write a series of tests to check different functionality has been implemented:

```go
// Test a known working case
func TestReverseHello(t *testing.T) {
	input := "hello"
	expected := "olleh"
	result := Reverse(input)
	if result != expected {
		t.Errorf("Reverse(%q) = %q; want %q", input, result, expected)
	}
}

// Check an empty string
func TestReverseEmpty(t *testing.T) {
	input := ""
	expected := ""
	result := Reverse(input)
	if result != expected {
		t.Errorf("Reverse(%q) = %q; want %q", input, result, expected)
	}
}
```

But you can condense these to just a single test, which is easy to update if you need to add more cases or extend functionality:

```go
func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"Go!", "!oG"},
		{"12321", "12321"},
	}

	for _, test := range tests {
		result := Reverse(test.input)
		if result != test.expected {
			t.Errorf("Reverse(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
```

## Adding Error Handling
If your function can return an error (e.g., parsing input, dividing numbers), you can extend the table to include expected errors.

Example, using our Divide function:

```go 
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

We can add a new column to our tests cases, stating if we expect to receive and error from them:

```go
func TestDivide(t *testing.T) {
    tests := []struct {
        a, b     float64
        expected float64
        wantErr  bool
    }{
        {10, 2, 5, false},
        {5, 0, 0, true},
        {9, 3, 3, false},
    }

    for _, test := range tests {
        result, err := Divide(test.a, test.b)
        // If err doesn't match what we expected
        if (err != nil) != test.wantErr {
            t.Errorf("Divide(%v, %v) error = %v, wantErr %v", test.a, test.b, err, test.wantErr)
        }
        // If we are not expecting and error but the results didn't not match we were not expecting
        if !test.wantErr && result != test.expected {
            t.Errorf("Divide(%v, %v) = %v; want %v", test.a, test.b, result, test.expected)
        }
    }
}
```


## Hands-On Lab
### Objective
Students will write and convert tests into a table-driven format, and add edge cases and error scenarios.

### Instructions
Choose or define **one** function to test. You can use one of the following or your own:

`IsPalindrome(s string) bool`
`Factorial(n int) (int, error)`
`ToUpperCase(s string) string`

- Write individual test cases for the function using testing.T.

- Convert the tests into a table-driven format.

- Add edge cases:

    - Empty strings
    - Very large or small inputs
    - Invalid inputs (e.g., negative numbers for factorial)
    - Include error handling if applicable.

```go
func Factorial(n int) (int, error) {
    if n < 0 {
        return 0, fmt.Errorf("negative input")
    }
    result := 1
    for i := 2; i <= n; i++ {
        result *= i
    }
    return result, nil
}
```
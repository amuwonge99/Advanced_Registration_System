package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	// We can assign the output of `flag.[Type](...)` to a variable but it will come back as a
	//  pointer type which may not always be ideal, so instead we can also do the inverse...
	myBoolAsPointer := flag.Bool("bool-ptr", false, "a flag stored by reference")

	// With the `flag.[Type]Var(...)` function (note the additional word 'Var' at the end of the
	//  function name), we can assign to a non-pointer variable by reference
	var myBoolAsValue bool
	flag.BoolVar(&myBoolAsValue, "bool-value", false, "a flag stored by value")

	// It common that long flags have a shorthand counter part, e.g.
	//     cut -d,          ...
	//     cut --delimiter, ...
	// The in-built flag package doesn't really provide a method of doing this (though several
	//  external packages do). Instead, we can achieve this by binding `myFlagAsValue` to multiple
	//  flags (`--value` and `-b`)
	flag.BoolVar(&myBoolAsValue, "b", false, "a flag stored by value")

	// The above are just Boolean flags that take no argument, if instead we wish for our flag to
	//  take a string we can use the `flag.StringVar()`
	// This will expect a value, e.g.
	//     --string "my string value"
	var myStringFlag string
	flag.StringVar(&myStringFlag, "string", "default value", "a string")

	// Flag can do some type checking/conversion for us, for example if we want an integer
	var myIntFlag int
	flag.IntVar(&myIntFlag, "int", 0, "an integer value")

	// We can even specify custom logic for a particular flag with `flag.Func`
	// Could also do some type checking/conversion in here too, whatever you fancy
	var myFileThatExists string
	flag.Func("file", "a file which must exist", func(s string) error {
		_, err := os.Stat(s) // Naive check that file exists
		if err != nil {
			return err
		}
		myFileThatExists = s
		return nil
	})

	// To get a flag one may specify multiple times, we need a custom type (see below) which
	//  satisfies the `Value` interface
	// We often use the `flag.Var` function for custom types
	var repeatFlag flagMultipleTimes
	flag.Var(&repeatFlag, "repeat", "a flag you can give multiple times")

	// Positional arguments (those without hyphenated options) are not obvious in their setup with Go
	//  unfortunately, one cannot bind these to variables or provide the same type logic
	// Best we can do is override the `flag.Usage` function _implying_ the existence of positionals
	flag.Usage = func() {
		// Here, we are implying with the addition of '[POSITIONALS]' that one may provide positional
		//  arguments, but that's all we can do, imply
		fmt.Printf("Usage: %s [OPTIONS] [POSITIONALS] ...\n", os.Args[0])
		// Then we print the other flags we've set and their defaults
		flag.PrintDefaults()
	}

	// Parse the CLI arguments to populate those variables
	flag.Parse()

	// Print them all out
	fmt.Printf("Bool as value   : %t\n", myBoolAsValue)
	fmt.Printf("Bool as pointer : %t\n", *myBoolAsPointer)
	fmt.Printf("String flag     : %s\n", myStringFlag)
	fmt.Printf("Int flag        : %d\n", myIntFlag)
	fmt.Printf("File that exists: %s\n", myFileThatExists)
	fmt.Printf("Repeated flag   : %s\n", repeatFlag)
	// We access the positionals (remaining arguments) with the `flag.Args()` method
	fmt.Printf("Positionals     : %s\n", strings.Join(flag.Args(), ", "))
}

// We'll be using this to be able to specify a flag multiple times and storing each of the values
type flagMultipleTimes []string

// The String method is required to satisfy the `Value` interface in the `flag` package
func (arr flagMultipleTimes) String() string {
	return strings.Join(arr, ", ")
}

// We specify how the `flag` package is to "set" our variable with each use of the option
func (arr *flagMultipleTimes) Set(value string) error {
	*arr = append(*arr, value)
	return nil
}

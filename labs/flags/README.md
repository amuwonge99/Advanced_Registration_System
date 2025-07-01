# Flags

- [How to Specify a Flag](#how-to-specify-a-flag)
- [Other Types of Flags](#other-types-of-flags)
  - [Strings](#strings)
  - [Integers](#integers)
  - [Floats](#floats)
- [Positional Arguments](#positional-arguments)
- [Custom Validation](#custom-validation)
- [Flag Aliases](#flag-aliases)
- [Task](#task)

Often times you want to provide your script with a means of receiving user input in a scriptable fashion.

For this, it is common to add options to your program, a known example of this would be with the `cut` command:

```sh
cut -d, -f2 my-file.txt
```

We're telling `cut` to use a **d**elimiter of `,`, to select the second (`2`) **f**ield, and to parse the file: `my-file.txt`.

## How to Specify a Flag

We use the [`flag`](https://pkg.go.dev/flag) package to do some of the heavy lifting for us. There are a number of external libraries available ([Cobra](https://github.com/spf13/cobra) being a community favourite), but this one is in the standard library and good enough for smaller projects.

The `flag` package exposes a number of functions related to setting up flags, one for each primitive type, these take the form:

```go
func <Type>(flagName string, defaultValue <type>, usage string) *<type>
```

So, for a Boolean flag:

```go
func Bool(flagName string, defaultValue bool, usage string) *bool
```

The return value of these functions are a pointer to a variable of that type, for example:

```go
var myBoolPtr *bool
myBoolPtr = flag.Bool("my-bool", false, "a flag stored by reference")
```

**An alternative** to handling pointers in your code (directly) is to use a variant of the above function which takes *a reference* to an existing variable (of non-pointer type usually) - these are the named the same but with `Var` at the end, for example:

```go
func <Type>Var(ptr *<type>, flagName string, defaultValue <type>, usage string)
```

So, for a Boolean flag:

```go
func BoolVar(ptr *bool, flagName string, defaultValue bool, usage string)
```

There is no return value but you pass in an existing variable to be populated as mentioned:

```go
var myBool bool
flag.BoolVar(&myBool, "my-bool", false, "a flag stored by value")
```

Either approach would provide us with the command-line flag `--my-bool` with the default value of `false`, stored in the relevant variable.

*Note*: One oddity of Go is that it prefers single-hyphen long options (`-my-long-option`) over the better double-hyphen options (`--my-long-option`). Luckily, the `flag` package covers both, unfortunately this means that combining single character options (`-dl` being the same as `-d -l`) is no longer possible.

Once we've set up our flags, we can call [`flag.Parse()`](https://pkg.go.dev/flag#Parse) to populate the variables.

One point of note is that the `flag` package *automatically* provides a `--help` (referred to as "usage") flag for us, though we may wish to override this as we'll see later.

## Other Types of Flags

It doesn't make sense to go through all of these, but there are a couple that are common enough to go into.

### Strings

```go
var myStringFlag string
flag.StringVar(&myStringFlag, "string", "", "an string value")
```

The above provides the `--string` flag which takes a string parameter, so `--string hello` would set `myStringFlag` to `"hello"`.

### Integers

```go
var myIntFlag int
flag.IntVar(&myIntFlag, "int", 0, "a integer value")
```

The above provides the `--int` flag which takes an integer parameter, so `--int 42` would set `myIntFlag` to `42`.

It should be noted that the type conversion from the command-line string to the integer can fail if the string is invalid for the type:

```sh
$ go run main.go --int hello
invalid value "hello" for flag -int: parse error
```

### Floats

Floating-point number can also be parsed for you.

```go
var myFloatFlag float64
flag.Float64Var(&myFloatFlag, "float", 0., "a float value")
```

The above provides the `--float` flag which takes a float parameter, so `--float 3.2` would set `myFloatFlag` to `3.2`.

## Positional Arguments

The last piece of the puzzle is positional arguments, consider `cut` again:

```sh
cut -d, -f2 my-file.txt
```

Here, the delimiter and the field are flags with values, but `my-file.txt` is simply at the end of the command line with no flag attached to it - this is a positional argument.

The `flag` package unfortunately **does not** provide a proper means of labelling and type-checking positional arguments.

The closest we can get to this is to override the [`flag.Usage`](https://pkg.go.dev/flag#pkg-variables) function and *imply* to the user that a positional argument is expected:

```go
flag.Usage = func() {
  //                                this bit
  fmt.Printf("Usage: %s [OPTIONS] [POSITIONALS] ...\n", os.Args[0])
  flag.PrintDefaults()
}
```

The remaining arguments after having called `flag.Parse()` are available in [`flag.Args() []string`](https://pkg.go.dev/flag#Args).

Validating their presence, their count, and their types is left to the developer to implement.

## Custom Validation

You may wish to provide a function which is able to validate a particular argument, a classic example is that a `--file` option should take an argument which is a file that exists, for this we can use [`flag.Func()`](https://pkg.go.dev/flag#Func).

This has no return value but you can use the function implementation to set any relevant variables:

```go
var myFileThatExists string
flag.Func("file", "a file which must exist", func(s string) error {
  _, err := os.Stat(s)
  if err != nil {
    return err
  }
  myFileThatExists = s
  return nil
})
```

The function takes the string parameter to the option and returns an error - use of the value and setting of the error is down to implementation.

## Flag Aliases

The `flag` package unfortunately **does not** provide a proper means of aliasing flags, commonly for providing shorthand flags. It common that long flags have a shorthand counter part, e.g.

```sh
cut -d,          ...
cut --delimiter, ...
```

with `-d` being the shorthand of `--delimiter`.

One way around this is to bind a variable to two separate flags, whichever is *used on the command line* last will be the found value of the variable, e.g.

```go
var myBool bool
flag.BoolVar(&myBool, "my-bool", false, "a flag stored by value")
flag.BoolVar(&myBool, "b", false, "a flag stored by value")
```

Here, one may set `myBool` with either `--my-bool` or just `-b`.

## Task

Create a CLI that satisfies this `--help` output:

```txt
  -verbose
        Run the program in verbose mode
  -v
        A shorthand for -verbose
  -days int
        The number of days
  -surname value
        A string representing a surname
  -country-code value
        A string which is exactly two characters in length
```

And print the values out afterwards.

<!-- markdownlint-disable-file MD013 -->

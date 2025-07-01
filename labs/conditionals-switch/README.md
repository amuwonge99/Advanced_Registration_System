# What is Switch Logic?
In Go, a `switch` statement is used to select one of many blocks of code to be executed. It's an alternative to writing many if-else statements and is often more readable and concise.

# Basic Syntax
```go
package main

import "fmt"

func main() {
    day := "Tuesday"

    switch day {
    case "Monday":
        fmt.Println("Start of the work week.")
    case "Tuesday":
        fmt.Println("Second day of the work week.")
    case "Friday":
        fmt.Println("Almost weekend!")
    default:
        fmt.Println("Just another day.")
    }
}
```

## Explanation:
`switch` day evaluates the value of day.
Each `case` checks for a match.
`default` runs if no case matches.

# Key Features of Go's Switch
## 1. No Fallthrough by Default
Unlike C/C++, Go does not fall through to the next case unless explicitly told to.
```go
switch x := 2; x {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
    fallthrough
case 3:
    fmt.Println("Three") // This will also run because of fallthrough
}
```

## 2. Multiple Values in a Case
```go
switch day {
case "Saturday", "Sunday":
    fmt.Println("Weekend!")
default:
    fmt.Println("Weekday.")
}
```

## 3. Switch Without an Expression
You can use switch like a cleaner if-else chain:

```go
score := 85

switch {
case score >= 90:
    fmt.Println("Grade: A")
case score >= 80:
    fmt.Println("Grade: B")
default:
    fmt.Println("Keep trying!")
}
```

# Lab: Build a Simple Command Interpreter Using switch
## Objective:
You will write a Go program that simulates a basic command-line interface using a switch statement. The program should respond to user commands like "help", "status", "restart", and "exit".

## Instructions:
Create a new Go file called command_interpreter.go.

Prompt the user to enter a command.

Use a switch statement to handle the following commands:

* "help" → Print a list of available commands.
* "status" → Print a fake system status (e.g., "All systems operational").
* "restart" → Print a message like "System restarting...".
* "exit" → Exit the program.
* Any other input → Print "Unknown command".
* The program should keep running until the user types "exit".

## Hints:
Use a for loop to keep the program running.
Use `fmt.Scanln()` to read user input.
Consider using `strings.ToLower()` to make input case-insensitive.

## Example Output:
```txt
Enter command: help
Available commands: help, status, restart, exit

Enter command: status
System status: All systems operational.

Enter command: foo
Unknown command.

Enter command: exit
Goodbye!
```
## Challenge Extension (Optional):
Add a "version" command that prints the version of your program (e.g., "v1.0.0").
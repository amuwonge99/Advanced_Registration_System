package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	err := openingAndEditingAnExistingFile()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println()

	err = creatingAndWritingANonExistantFile()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println()

	err = appendingToAFile()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println()

	fmt.Println(
		"Have a look around your current directory, one file should have changed and " +
			"another created\nRun ./cleanup.sh to reset to run again",
	)
}

func openingAndEditingAnExistingFile() error {
	fmt.Println("----------------- OPENING AND EDITING AN EXISTING FILE -----------------")

	// Because we are referring to this file as `./...`, the program will attempt to read the file
	//  from your current directory, i.e. whichever directory you were 'cd'ed in when you ran
	//  `go run main.go`
	existingFile := "./my-file.txt"

	// We use a library function to read our file, in its entirety, into memory - there are issues
	//  with this for larger which is addressed later in the lab
	data, err := os.ReadFile(existingFile)
	if err != nil {
		return fmt.Errorf("Error reading file (%s): %w\n", existingFile, err)
	}

	// `data` is a `[]byte`, let's work with something a little more known, a string
	contents := string(data)

	// And print out the contents
	fmt.Printf("os.ReadFile contents:\n%s", contents)

	fmt.Println()

	// Let's make a couple changes to our file contents (now as a string)
	newContent := strings.Replace(contents, "a file", "Ben!", 1)

	// And write our string to the file
	// `WriteFile` works with byte-arrays, so we cast to that type
	err = os.WriteFile(existingFile, []byte(newContent), 0644)
	if err != nil {
		return fmt.Errorf("Error writing file (%s): %w\n", existingFile, err)
	}

	// Just for verification, we re-read the file and output its contents
	data, err = os.ReadFile(existingFile)
	if err != nil {
		return fmt.Errorf("Error reading file (%s): %w\n", existingFile, err)
	}
	fmt.Printf("os.ReadFile contents (changed):\n%s", string(data))

	// No error, so a nil return value
	return nil
}

func creatingAndWritingANonExistantFile() error {
	fmt.Println("----------------- CREATING AND EDITING A NON-EXISTANT FILE -----------------")

	// Now let's do the same thing with a file that doesn't exist
	nonExistantFile := "./this-file-does-not-exist.txt"

	// Now, we open a file that we know doesn't exst...
	data, err := os.ReadFile(nonExistantFile)

	// ... so, we're expecting an error and return with an error code (1) if no error is found
	if err == nil {
		return fmt.Errorf("File (%s) was read but shouldn't exist\n", nonExistantFile)
	}

	fmt.Printf("Expected and received error for: %s\n(%s)\n", nonExistantFile, err)

	fmt.Println()

	// Here, we simply create the file and receive what's known as a "handle" to it
	// Under the hood this this is pretty much just a number referring to the file but in Go, we get
	//  a nice `File` object to play with
	file, err := os.Create(nonExistantFile)
	if err != nil {
		return fmt.Errorf("Error creating file (%s): %w\n", nonExistantFile, err)
	}

	// Since we now have an open file handle... we'll need to close it at some point before we end the
	//  function. But our function may have several exit points (`return` statements), so we can either:
	//    - Put a `file.close()` before _every_ exit point, or
	//    - Defer (`defer`) the closing of the file until we close the function off
	//  this way, whenever we exit from the function, our file handle is closed
	defer file.Close()

	fmt.Printf("File (%s) opened with descriptor: %d\n", nonExistantFile, file.Fd())

	_, err = file.WriteString("Hello, I didn't exist before now\n")
	if err != nil {
		return fmt.Errorf("Error writing file (%s): %w\n", nonExistantFile, err)
	}

	// Let's copy what we did before and print out the file contents for verification
	data, err = os.ReadFile(nonExistantFile)
	if err != nil {
		return fmt.Errorf("Error reading file (%s): %w\n", nonExistantFile, err)
	}
	fmt.Printf("os.ReadFile contents (changed):\n%s", string(data))

	// No error, so a nil return value
	return nil
}

func appendingToAFile() error {
	fmt.Println("----------------- APPENDING TO A FILE -----------------")

	// This file exists and already has content, let's not wipe out the file and instead append to it
	existingFile := "./my-file.txt"

	// Here, we simply open the file and receive another handle to it
	// Those arguments at the end are open flags, these form one single int, but Go provides us with
	//  more readable versions of these number that we can "stack" on top of eachother:
	//    - O_APPEND: Append to the file, don't overwrite
	//    - O_RDWR:   Open permitting both read and write
	file, err := os.OpenFile(existingFile, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("File (%s) was unable to be opened: %w\n", existingFile, err)
	}

	// We have a handle again, so we'll need to delay its closing again
	defer file.Close()

	file.WriteString("Hello, I am a log line\n")

	// Let's copy what we did before and print out the file contents for verification
	data, err := os.ReadFile(existingFile)
	if err != nil {
		return fmt.Errorf("Error reading file: %s (%w)\n", existingFile, err)
	}
	fmt.Printf("os.ReadFile contents (appended to):\n%s", string(data))

	// No error, so a nil return value
	return nil
}

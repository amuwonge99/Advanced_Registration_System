package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"syscall"
)

func main() {
	readLineByLine()
	fmt.Println()

	theCishWay()
	fmt.Println()

	problemsWithUmask()
}

func readLineByLine() {
	fmt.Println("----------------- READ FILE LINE BY LINE -----------------")

	fileWithLines := "./file-with-lines.txt"

	// Open the file with the default "read-only" permission
	file, err := os.Open(fileWithLines)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// Defer the closing of the file until the end of the function
	defer file.Close()

	// `bufio.NewScanner` takes an `io.Reader` which is something that `os.Open`'s return value,
	//  `File`, satisfies
	scanner := bufio.NewScanner(file)

	// `bufio.Scanner.Split` takes a function that indicates how to split the reader
	// A handy function (`bufio.ScanLines`) is already in the standard library to deal with
	//  line-parsing - **this is the default** so the addition of this line is more for demonstration
	// Similarly, there is a `ScanWords`, `ScanBytes` (think ASCII), and `ScanRunes` (think UTF-8)
	scanner.Split(bufio.ScanLines)

	// This loop will run until `scanner.Scan` returns false
	// `scanner.Scan` will return true until there are no more newlines and EOF (end-of-file) is
	//  reached
	// (The index (`idx`) is just for demonstration purposes, do with this loop as you wish)
	var idx int
	for scanner.Scan() {
		fmt.Printf("#%d : %s\n", idx, scanner.Text())
		idx++
	}
}

func theCishWay() {
	// The purpose of this function is to demonstrate the benefit of standard library functions when
	//	writing your Go code - this the ol' skool way
	// Though one can learn a lot from implementing everything yourself, there is undoubtedly added
	//	complexity, more room for human error, and ultimately redundancy in reimplementing something
	//	that already exists...
	fmt.Println("----------------- READING FILES THE C-ISH WAY -----------------")

	existingFile := "../my-file.txt"

	if contents, err := readFileTheCishWay(existingFile); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read file (%s): %s\n", existingFile, err)
	} else {
		fmt.Printf("Contents:\n%s\n", contents)
	}

	nonExistantFile := "./this-file-does-not-exist.txt"
	if contents, err := readFileTheCishWay(nonExistantFile); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read: %s:\n%s\n", nonExistantFile, err)
	} else {
		fmt.Printf("Contents:\n%s\n", contents)
	}
}

func readFileTheCishWay(filename string) (string, error) {
	fmt.Printf("Reading file (%s) in the C-ish way\n", filename)

	// Here, we simply "open" the file and get what's known as a "handle" to it. Under the hood this
	//  this is pretty much just a number referring to the file
	//  But in Go, we get a nice `File` object to play with
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	// Since we now have an open file handle... we'll need to close it at some point before we end the
	//  function. But our function has several exit points (`return` statements), so we can either:
	//    - Put a `file.close()` before _every_ return statement
	//    - Defer (`defer`) the closing of the file until we close the function off
	//  this way, whenever we `return` from the function, our file handle is closed
	defer file.Close()

	// We move our file pointer to the end of the file and the byte-offset in the file is returned to
	//  us - as the last byte, this is also how large the file is (in bytes)
	size, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		return "", err
	}

	// Our pointer is at the end of the file, but we need to read from the start of the file,
	//  so we "seek" the "start" of the file
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}

	// Create an empty buffer in which we can store the data
	data := make([]byte, size)

	// Here, we read a certain number of bytes from the file, this is constrained by two factors:
	//    1. The size of our buffer (`size`)
	//    2. The maximum number of bytes that `os.Read` will read
	// This is where some **issues** with the "do it yourself" approach comes in, `os.Read` can only
	//  read an int32's worth of bytes at a time (2,147,483,647). Fine for our `my-file.txt` but not
	//  if my file is above 2GB.
	// The "proper" solution would be to iterate these reads, reading `math.MaxInt32` chunks, adding
	//  to the slice (`data`) from where you left off... bit of a pain. This is why `os.ReadFile`
	//  exists, abstracts this more complex implementation away from the user
	nBytesRead, err := file.Read(data)
	if err != nil {
		return "", err
	}

	// If we have read fewer bytes (`nBytesRead`) than the size of our file (`size`), i.e. 2GB, then
	//  in this simple implementation I just return and error
	// This is basically a "file too big for this implementation" error
	if int64(nBytesRead) < size {
		return "", fmt.Errorf("Number of bytes read (%d) does not match number of bytes expected: %d", nBytesRead, size)
	}

	return string(data), nil
}

func problemsWithUmask() {
	fmt.Println("----------------- PROBLEMS WITH UMASK ----------------")

	filename := "/tmp/some-file"
	// This will fail if the file doesn't exist... but that doesn't matter here
	os.Remove(filename)

	fmt.Println("Writing file with 777 mode: -rwxrwxrwx")
	if err := os.WriteFile(filename, []byte{}, 0777); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file (%s): %s\n", filename, err)
	}

	stat, err := os.Stat(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error stating file (%s): %s\n", filename, err)
	}

	fmt.Printf("Actual file mode from stat: %s\n", stat.Mode())

	// It should be noted that this mechanism is deprecated, there is an external `unix` package now
	//  or just use `os.Chmod`
	newUmask := 0
	defaultUmask := syscall.Umask(newUmask)

	// We defer the setting of the original umask to the end of the function
	defer syscall.Umask(defaultUmask)

	fmt.Printf("Default umask (0%o)       : -%09b\n", defaultUmask, defaultUmask)
	fmt.Println("                          :      1  1")

	fmt.Println()

	fmt.Printf("New umask (0%o)            : -%09b\n", newUmask, newUmask)

	os.Remove(filename)

	fmt.Println("Writing again with 777    : -rwxrwxrwx")
	if err := os.WriteFile(filename, []byte{}, 0777); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file (%s): %s\n", filename, err)
	}

	stat, err = os.Stat(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error stating file (%s): %s\n", filename, err)
	}

	fmt.Printf("New file mode from stat   : %s\n", stat.Mode())
}

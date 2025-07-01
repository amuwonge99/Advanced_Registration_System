# File IO

This lab's broken down into a few sections:

- [Main.go](#maingo)
  - [Opening and Editing an Existing File](#opening-and-editing-an-existing-file)
  - [Creating and Writing a New File](#creating-and-writing-a-new-file)
  - [Appending to an Existing File](#appending-to-an-existing-file)
- [A Task](#a-task)
- [Bonus/main.go](#bonusmaingo)
  - [Reading a File Line-by-line](#reading-a-file-line-by-line)
  - [Reading a File How you Would in C](#reading-a-file-how-you-would-in-c)
  - [Problems with Umask](#problems-with-umask)

File IO is an expansive topic and not all the information provided here is strictly "necessary", I've noted superfluous information with "***Extra***:".

## Main.go

Because file IO is such a fairly hefty topic, it made sense to split it into common operations one may wish to perform with files.

*Note*: We haven't covered [command-line arguments as yet](https://pkg.go.dev/flag), so all the paths to files we use are [relative paths](https://desktop.arcgis.com/en/arcmap/latest/tools/supplement/pathnames-explained-absolute-relative-unc-and-url.htm). In short, to run the lab, you need to be in the right directory:

```sh
cd ./lab-file-io
go run main.go
```

*Note*: This lab edits files, a script ([`cleanup.sh`](./cleanup.sh)) is provided to reset the lab back to its former self. One could do this with `git` commands, but a script seemed easier:

```sh
# To make it easier to just hit "up and enter" to re-run the lab, you could
#  combine them into one command: one to setup the files, and one to run the Go
./cleanup.sh && go run main.go
```

### Opening and Editing an Existing File

Opening an existing file and "editing" it (i.e. changing the contents in memory and writing back to disk) seemed a good place to start.

Our first task is to read the contents of the given file (`existingFile`) into memory. The standard [`os`](https://pkg.go.dev/os) package has what we need for this, [`os.ReadFile(string)`](https://pkg.go.dev/os#ReadFile).

What is returned from this is a pair of values:

1. An array of bytes (`[]byte`) containing the byte-values of our file's contents
2. An error which is set if the read operation encountered an error

Next up we want to edit this content. One might choose to edit the byte-array directly but as we're more familiar with the string (`string`) type at this point it makes sense to convert the array to a string in the form:

```go
myString := string([]byte{0x01, 0x02, 0x03})
```

We are now free to edit the string.

***Extra***: It should be *noted* that [Go strings are immutable](https://www.educative.io/answers/strings-in-golang) and so each "edit" is secretly a recreation of the whole string, but this can be ignored for our purposes.

We use the [`strings`](https://pkg.go.dev/strings) package's `strings.Replace(string, source, dest, count)` function to make our edit:

```go
newContent := strings.Replace(contents, "a file", "Ben!", 1)
```

Finally, we need to write our string back to the file. We used `os.ReadFile(string)` to read the file, so as you might guess there's a complementary [`os.WriteFile(string, []byte, FileMode)`](https://pkg.go.dev/os#WriteFile) we can use.

The [`FileMode`](https://pkg.go.dev/os#FileMode) (`uint32` under the hood) represents the file permissions of the written file, we use a fairly standard `0644` for this:

Character | Meaning
:--       | :--
`0`       | This number is to be interpreted as octal
`6`       | Owning user can read and write
`4`       | Owning group can read
`4`       | Anyone else can read

***Extra***: There are [a number](https://cs.opensource.google/go/go/+/refs/tags/go1.22.2:src/os/types.go;l=35) of other file modes available that we won't be getting into - for *some* further reading, [read here](https://www.cbtnuggets.com/blog/technology/system-admin/linux-file-permissions-understanding-setuid-setgid-and-the-sticky-bit).

As before we converted the byte-array to a string, we now need to do the same in reverse in the form:

```go
myByteArray := []byte(myString)
```

`os.WriteFile` has only a single return value, an error if an error is encountered so we should check that (perhaps you have the permission to read the file, but not write to it).

Just to make things obvious, we also read the file again and print its contents to see that the change has been made.

### Creating and Writing a New File

Here, the file we're targeting (`nonExistantFile`) doesn't exist. This is demonstrated using, again, the `os.ReadFile` function which we expect to fail (no file to read).

So, seeing that `os.ReadFile` has failed, we use another function in the `os` package: [`os.Create(string)`](https://pkg.go.dev/os#Create).

This does as you might expect, creates a file that doesn't exist.

***Extra***: One **important note** for this is that if you run `os.Create` on a file **that already exists**, the file will be truncated, i.e. the contents wiped.

What is returned from this is a pair again:

1. A [`File`](https://pkg.go.dev/os#File) pointer referring to our file
2. An error which is set if the create operation (technically an [open](https://man7.org/linux/man-pages/man2/open.2.html) operation on UNIX systems) encountered an error

The `File` object represents a [file handle](https://www.lenovo.com/us/en/glossary/file-handling/), under the hood this is just an integer representing the opened file. But Go gives us a more user-friendly object to use.

To demonstrate this, a line referring to the underlying number (file descriptor) is printed to the console.

One important aspect of file handles is that open handled is closed before the program exits to avoid resource leaks and access issues later down the line. To do this, `File` provides a [`Close`](https://pkg.go.dev/os#File.Close) method for just this purpose.

***Extra***: The OS commonly cleans up for you in the event that you do not close the handle, but this isn't always guaranteed (often the case with embedded systems) and so good practice should be followed, close your handles!

Common practice is to [defer](https://go.dev/tour/flowcontrol/12) the close operation. This means that whenever the function returns, the handle will be closed for you:

```go
defer file.Close()
```

***Extra***: If you use `os.Exit` to exit from your program with a status code, `defer`ed calls are not honoured, only `return`s from a function. There is [a hacky way](https://stackoverflow.com/a/24601700) around this, but it's better to use a separate function.

We now want to write some content to our file via the handle and use the [`File.WriteString(string)`](https://pkg.go.dev/os#File.WriteString) method on our `File` object:

```go
nBytes, error := file.WriteString("My string\n")
```

A pair of values is returned:

1. The number of bytes written (which should match the number of bytes in our string)
2. An error in the event that the write operation fails

***Extra***: There is an [`os.Write([]byte)`](https://pkg.go.dev/os#File.Write) method instead, but as we're more used to strings it made sense to use this convenience method.

Once again, we read the file using `os.ReadFile` and print the contents to verify the operation.

***Extra***: `os.ReadFile` here will open another handle to the file before we close the first, this is perfectly acceptable - a file can have multiple open handles to it.

***Note***: We *could* have used the `os.WriteFile` function again from before from the get go and avoided the `os.Create` call completely: `os.WriteFile` opens the file with the [`O_CREATE`](https://pkg.go.dev/os#pkg-constants) flag, however this seemed like a good opportunity to go into file handles for finer-grained control of the file.

### Appending to an Existing File

This uses a lot of information from the previous lab, in short we need to open a given file (`existingFile`) in append mode, and write content to it.

Instead of creating a file this time, we go direct to the [`os.Open(string, int, FileMode)`](https://pkg.go.dev/os#Open) function which takes three parameters:

1. A string containing the path of the file in question
2. File opening **flags** (more on this below)
3. The file permissions as per [Opening and Editing an Existing File](#opening-and-editing-an-existing-file)

The **flags** mentioned refer to the IO mode (i.e. read, write, or read/write) and behaviours (e.g. append, create, truncate, etc.). These are in the `os` package as [constants](https://pkg.go.dev/os#pkg-constants). For our purposes, we want read/write mode and the append behaviour when opening:

```go
os.OpenFile(existingFile, os.O_APPEND|os.O_RDWR, 0644)
```

It should be noted that if we're not sure if the file exists already, we could add `os.O_CREATE` into our mix of file behaviours:

```go
os.OpenFile(filename, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
```

, and this will create the file if it doesn't exist.

***Extra***: These IO modes and behaviours are combined together into a single integer (`int`) using a [bitwise OR](https://yourbasic.org/golang/bitwise-operator-cheat-sheet/).

Once again, we have an open file handle and so defer its closing.

And we use `os.WriteString` as per [Appending to an existing file](#appending-to-an-existing-file) to append our new content.

## A Task

- Take the file [`my-task.txt`](./my-task.txt) and open it
  - We won't need to read the content, so write-only
  - This operation should truncate the existing content (emptying the file of contents)
- Write new content to the file of your choosing
  - Can be anything you fancy, but give it some actual content to play with
- Close the file (can either be manual or through `defer`)

- Read the contents of the file (your method of choice)
- If the file contains an `a` character, append a line saying `I contain an 'a'`
- **Otherwise**, append a line saying `I do not contain an 'a'`

- Finally, print the **new** contents of the file out

## Bonus/main.go

If people want to read further into file IO, these are a couple extra mini-labs demonstrating some techniques and operations.

*Note*: Once again, because we haven't covered command-line arguments, all the paths to files we use are relative, so to run the lab, you need to be in the right directory:

```sh
cd ./lab-file-io/bonus
go run main.go
```

### Reading a File Line-by-line

Thus far, every time we've read a file, we've read the entire contents of the file into memory. This is fine for our small file, but consider opening a 12GB file... that's 12GB of memory you need spare.

You often want instead to read a file, chunk by chunk, into a buffer of a fixed size only ever using a single chunk of memory of that buffer's size.

Go provides the [`bufio`](https://pkg.go.dev/bufio) package for just this purpose.

First off, we use `os.Open` to get a handle to the file. But, instead of reading the whole thing, we instantiate a [`bufio.Scanner`](https://pkg.go.dev/bufio#Scanner), this is what we'll use to split by line:

```go
scanner := bufio.NewScanner(file)
```

We set the scanning function we wish to use the convenient [`bufio.ScanLines`](https://pkg.go.dev/bufio#ScanLines) function which will scan for newlines and/or end of file (EOF) characters in our file:

```go
scanner.Split(bufio.ScanLines)
```

*Extra*: `ScanLines` is the default function but it made sense to include for demonstration purposes.

Then, we iterate the [`Scanner.Scan()`](https://pkg.go.dev/bufio#Scanner.Scan) method to "load" the next chunk (referred to as a "token") and access the last-scanned token with [`Scanner.Text()`](https://pkg.go.dev/bufio#Scanner.Text):

```go
for scanner.Scan() {
  myToken := scanner.Text()
}
```

*Note*: The maximum token size if 64KiB (as shown [here](https://pkg.go.dev/bufio#Scanner.Buffer) and again [in the source code](https://go.dev/src/bufio/scan.go) (line 86 at the time of writing)), so if your tokens are longer than 64KiB, you need another implementation.

### Reading a File How you Would in C

This exists solely as a history-lesson in how certain operations would commonly take place in older languages like [C](https://en.wikipedia.org/wiki/C_(programming_language)).

Largely, much is the same, there's just a lack of convenience - it's this convenience which makes languages like Go "high-level" languages and demonstrates the power of Go's standard library.

So, we open a handle to the file as we have been doing and wish to read the entire contents into memory.

But how big do we make the buffer to store the contents? We need to know how large the file is. There's a couple ways around this but the "standard" C-ish way is to move our file handle to the end of the file using the [`File.Seek(int64)`](https://pkg.go.dev/os#File.Seek) method:

```go
size, err := file.Seek(0, io.SeekEnd)
```

In C, you would then use the [`ftell`](https://en.cppreference.com/w/c/io/ftell) function to be informed of the current byte-offset from the start of the file, as we're at the end of the file, this is the number of bytes in the file. Go's `File.Seek` does this for us in its first return value (`size`).

We now know how big to make our buffer, but our pointer is at the end of the file, we need to seek again, seeking the start:

```go
size, err := file.Seek(0, io.SeekStart)
```

Then we make our buffer (which in C would have to be freed when done but again Go handles this for us):

```go
data := make([]byte, size)
```

And finally we can read into our buffer:

```go
nBytesRead, err := file.Read(data)
```

However, this could lead to some errors, [`File.Read`](https://pkg.go.dev/os#File.Read) only reads a certain number of bytes (an `int32`'s worth), but if our file is larger than 2GiB we'll need to iterate these reads 2GiB at a time...

This is, in fact, what the [`os.ReadFile` implementation](https://go.dev/src/io/fs/readfile.go) does for us (line 52 at the time of writing). It also uses [`os.Stat`](https://pkg.go.dev/os#File.Stat) instead of the seek-and-tell method.

### Problems with Umask

One topic that's a little out of scope for this session is the UNIX [umask](https://askubuntu.com/a/44548). This is a security-ish feature of UNIX-based/styled systems (BSD (Mac) and Linux being the big ones) which enables a system administrator to set file permissions that a user of the system cannot set on files by default.

The value for this can be found with the [`umask` command](https://www.geeksforgeeks.org/umask-command-in-linux-with-examples/).

Though easily set again with the `umask` command (e.g. `umask 0`) or ignored with the [`chmod` command](https://www.nexcess.net/help/what-is-chmod/).

We can see the effect of this mask in the `problemsWithUmask` function. Note how the bits of the umask when printed align with the missing permissions despite the `0777` permissions we requested.

We override this with [`syscall.Umask`](https://pkg.go.dev/syscall#Umask) function though **it should be noted** that [the `syscall` module has been deprecated](https://github.com/dominikh/go-tools/issues/1412) since Go 1.11 - though its use *is* still permitted as "*it is impossible to avoid it entirely*".

The "correct" way to do this now would be through the [`sys/unix` package](https://github.com/golang/sys/):

```go
import "golang.org/x/sys/unix"

func main() {
    unix.Umask(0)
    // ...
}
```

<!-- markdownlint-disable-file MD013 -->

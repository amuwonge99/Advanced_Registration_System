IMPORTS
=======

There is a very large amount of code out there that can make common tasks a lot easier. To use this code we will need to *import* this into our code, a very common example is the `fmt` or the [format package](https://pkg.go.dev/fmt).

IMPORTING A PACKAGE
-------------------

We can import dependent code by using the `import` keyword. Our imports always come at the top of our file immediately after we define our package.

```go
package main

import "fmt"
...
```

After we have imported a package we can use any function within it.

```go
fmt.Println("Hello World!")
```

There are two important parts here:

1. We have to prefix the function with the package it comes from, in this case `fmt`
2. The second part is that the function starts with a capital and is in CamelCase. This is [how Go differs](https://go.dev/tour/basics/3) between private functions, which can only be used inside the code where it is defined, and public functions which can be used after being imported.

STANDARD LIBRARY VS THIRD-PARTY PACKAGES
----------------------------------------

The format package is an example of a package from the Go [standard library](https://pkg.go.dev/std). A built-in library of reusable code that you can use for common tasks.

Third-party packages are packages outside of the standard library, and are imported by referencing their fully-qualified names, usually the 
site that hosts the code, the user or organization that maintains it, and the base name of the package. When referencing the package in your code, you will only use the base name.


```go
package main

import "github.com/google/uuid"
```

```go
	id := uuid.New().String()

	fmt.Println(id)
```

In the above example, the package is hosted on github.com, maintained by Google, and called uuid. To be able to import this package, we must get it as a dependency in our go mod file, this can be done by running

```
go get github.com/google/uuid
```

We can also import local packages like we do third-party packages, if we look at our go.mod file we can see the name of our module

```
gitlab.platform-engineering.com/Go-academy/lab-07
```

This allows us to import our local greet package using the fully qualified name. By joining our module name and package name:

```go
package main

import "gitlab.platform-engineering.com/Go-academy/lab-07/greet"
```

PUBLIC VS PRIVATE FUNCTIONS
---------------------------

In a given package, there will be public and private functions. Public functions start with a capital letter, and can be accessed from other packages. Private functions start with a lowercase letter, and can only be accessed from within the same package.


```go
package greet

func Hello()   // this can be used outside of the package as it starts with a capital
func goodbye() // this cannot be used outside the package as it starts with a lower case character
```

LAB TASK
--------

Create your own local package, write both a public and private function in the package, then import and use the public function in your main.go.
Stretch goal: Import a package from the Go standard library and call a function from it in the main.go.
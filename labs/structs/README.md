STRUCTS
=======

[Structs](https://go.dev/tour/moretypes/2) are a way where we can create a group of fields together to define an object.

Structs are Go's implementation of data-structures, grouped data, but their semantics permit a form of Object Orientated programming.

CREATING A NEW STRUCT
---------------------

We create all structs the same way:

- We start with the `type` keyword meaning we are creating a new variable type
- This is immediately followed by the name of the type we are creating
- Then we end the statement with `struct` to show it is a collection of other types
- Between the curly braces we can then define the fields that make up this `struct`

```go
type MyStruct struct {
  ...
}
```

For example if we wanted to create a shopping service we might create a user with the following details:

```go
type User struct {
  Email string

  FirstName string
  LastName  string

}
```

We can then create a variable of that type in the same way we work with built-in variables:

```go
// using the var keyword to create the variable
var myUser User

// we can then use dot notation to set fields within the variable or print them out
myUser.Email = "foo@bar.com"
println(myUser.Email) // foo@bar.com

// we can also use the := notation for creation and assignment like other variables
myUserTwo := User{
  Email:     "john.doe@bar.com",
  FirstName: "John",
  LastName:  "Doe",
}

println(myUserTwo.Email) // john.doe@bar.com
```

The fields of a struct can be any valid type: string, int, slice, etc... This includes another struct that we have defined. The field of a struct that is also a struct is called a nested struct.

```go

type User struct {
	Email string

	FirstName string
	LastName  string

	Address Location
}

type Location struct {
	HouseNumber int
	Street      string
	City        string
	Postcode    string
}

```
LAB TASK
--------

In this directory you can find a file called `lab.go` which has some broken code inside of it.

Fix the definition of the `Person` struct so that when you run `go run lab.go` you get the following output:

```sh
$ go run lab.go
Hello I am John Doe and I am 30 years old
```

Stretch goal: Define your own struct in lab.go, and create a field under the `Person` struct that is of that type, update the information in `myPerson` and fetch it again in the println.
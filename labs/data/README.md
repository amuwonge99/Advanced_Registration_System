CONSUMING DATA
==============

A lot of the data we like our programs to consume comes in some well known types like [JSON](https://www.json.org/json-en.html) or [YAML](https://yaml.org/). Wouldn't it be great if Go understood YAML or JSON and allowed us to converted directly to or from data to a struct?

TAGS
----

When we create a structure we can add tags if we know we are going to be converting to/from a given data type. For example:

```go
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}
```

This `Person` struct has JSON tags defined for each of its fields which means that if we get sent a JSON message with those fields, we can immediately convert that into a `Person` struct that we can then manipulate.

To do that we can do the following:

```go
jsonString := `{"first_name":"John","last_name":"Doe","age":32}`
var person Person
json.Unmarshal([]byte(jsonString), &person)
```

In this example, we have a JSON string that contains all our fields, and we create ourselves a new `Person` using the `var` keyword that is currently only using the default values.

We then call `json.Unmarshal` with our `jsonString` (cast to a slice of bytes) and then the result of that is put into the `person` variable we just created.

We can then use that object like normal and all the fields will be set:

```go
fmt.Println(person.FirstName)
// John
```

LAB TASK
--------

Utilising what we learned in the file I/O task, and the given example of interacting with json data, write some code that will read in the file data.json and unmarshal the json into our `animal` struct.

<!-- markdownlint-disable-file MD013 MD010 -->

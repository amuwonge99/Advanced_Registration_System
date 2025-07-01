package main

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

// we can easily convert to/from json
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

// we can have multiple tag types if we are going to be using multiple data types
type Animal struct {
	Name    string `json:"name" yaml:"name"`
	Species string `json:"species" yaml:"species"`
	Age     int    `json:"age" yaml:"age"`
}

func main() {

	p := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       32,
	}

	bytes, err := json.Marshal(p)
	if err != nil {
		// lets just panic as we can't carry on
		panic(err)
	}

	// we will take the bytes that json.Marshal and cast them to a string
	fmt.Println(string(bytes))

	// lets try taking in some json data and converting it to yaml
	animalString := `{"name":"Oreo","species":"cat","age":2}`

	// first lets unmarshal it from json
	var animal Animal
	json.Unmarshal([]byte(animalString), &animal)

	// then convert it into yaml
	bytes, err = yaml.Marshal(animal)
	if err != nil {
		// lets just panic as we can't carry on
		panic(err)
	}

	// Lets prove that we consumed it correctly
	fmt.Println("animal name is", animal.Name)

	// we will take the bytes that yaml.Marshal and cast them to a string
	fmt.Println(string(bytes))
}

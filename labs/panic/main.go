package main

import "fmt"

func main() {
	num := 5
	fmt.Println("about to goDoSomething()")
	goDoSomething(num)
	// this line of code will not get run as we are panicking in the function call above
	fmt.Println("I will never get printed")
}

func goDoSomething(num int) {
	num = num * num
	goDoSomethingElse(num)
}

func goDoSomethingElse(num int) {
	num = num + 10
	goPanic(num)
}

func goPanic(num int) {
	if num == 35 {
		panic("ah!!! I can't carry on if num is 35")
	}
}

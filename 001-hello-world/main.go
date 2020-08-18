package main

import (
	"fmt"
)

func main() {
	var age int64
	age = 27

	var isEmployed bool
	isEmployed = true
	fmt.Printf("Are you employed ? %t", isEmployed)

	fmt.Print(age)
	var message = "Hello, Playground"
	var greeting, name = "Welcome", "Stanley"
	fmt.Println(message)
	fmt.Println(greeting, name)
}

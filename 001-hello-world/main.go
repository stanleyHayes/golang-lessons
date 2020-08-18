package main

import (
	"fmt"
)

func main() {
	//////////////////////////////////////////////
	//short declaration
	occupation := "Teaching"
	fmt.Println(occupation)
	fmt.Printf("Type of %s is %T\n", occupation, occupation)

	//use short declaration in loops e.g. for loops
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// two ways to get the type of variable
	variable := 123
	fmt.Printf("%T\n", variable) // int

	fmt.Print(reflect.TypeOf(variable).String()) // int
}
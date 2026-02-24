package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")

	// Variables from variable.go
	var student1 string = "John" //type is string
	var student2 = "Alvi"        // type is inferred
	x := 2                       // type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	fmt.Printf("T\n", student2)
}

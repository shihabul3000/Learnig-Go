package main

import "fmt"

// for -> only construct in go for looping
func main() {

	// while loop
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1

	// }

	// infinite loop

	// for {
	// 	fmt.Println("1")
	// }

	// classing for loop

	// for i := 0; i <= 3; i++ {
	// 	fmt.Println(i)
	// }

	for i := range 10 {
		fmt.Println(i)
	}

}

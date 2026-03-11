package main

import "fmt"

// Slices  slices is dynamic arrays
// most used construct in go
// +useful methods

func main() {

	// uninitialized slice is nill=null in go we use nill
	// var nums []int

	// fmt.Println(len(nums))

	//var nums = make([]int, 2, 5)

	//capacity -> maximum numbers of elements can fit
	// fmt.Println(cap(nums))

	// fmt.Println(nums == nil)

	a := [6]int{5, 4, 3, 2, 1}

	var b []int = a[1:6]
	fmt.Println(b, len(b))

}

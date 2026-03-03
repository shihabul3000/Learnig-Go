package main

import "fmt"

const (
	port = 5000
	host = "localhost"
)
const age = 30

var name string = "golang"

func main() {

	//const name string = "golang"

	//const age = 30
	//name := "javascript"

	fmt.Println(age, name)

	fmt.Println(port, host)
}

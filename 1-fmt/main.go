package main

import "fmt"

func main() {
	var firstName string = "Alfian"
	var lastName string = "Taka"

	fmt.Println("Hello, '", firstName, lastName, "'!")
	fmt.Printf("Hello, '%s %s'!\n", firstName, lastName)
}
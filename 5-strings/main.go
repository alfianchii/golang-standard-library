package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string = "Alfian Taka"

	fmt.Println(strings.Contains(name, "Alfian"))
	fmt.Println(strings.Split(name, " "))
	fmt.Println(strings.ToLower(name))
	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.Trim("!!" + name + "!!", "!"))
	fmt.Println(strings.ReplaceAll(name, "Alfian", "Cool"))
}
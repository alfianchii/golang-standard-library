package main

import (
	"fmt"
	"os"
)

func main() {
	var args []string = os.Args

	for _, arg := range args {
		fmt.Println(arg)
	}	
}
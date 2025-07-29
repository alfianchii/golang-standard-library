package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var name string = "Alfian Taka"
	var encodedName = base64.StdEncoding.EncodeToString([]byte(name))

	fmt.Println(encodedName)

	var decodedName, err = base64.StdEncoding.DecodeString(encodedName)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(decodedName)
	}

	fmt.Println([]byte(name))
	fmt.Println(string(name))
	fmt.Println("================================================================================================================")
}
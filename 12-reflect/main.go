package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string
}

type User struct {
	Name string
	Address string
	Age int
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)

	fmt.Println("Type name:", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)
		fmt.Println(valueField.Name, "w/", valueField.Type, "type.")
	}
}

func main() {
	var alfian User = User{
		Name: "Alfian",
		Address: "Indonesia",
		Age: 69,
	}
	var alfianType reflect.Type = reflect.TypeOf(alfian)
	var alfianFields reflect.StructField = alfianType.Field(0)

	fmt.Println(alfianType, alfianFields)
	fmt.Println("================================================================================================================")

	var taka User = User{
		Name: "Taka",
		Address: "Indonesia",
		Age: 69,
	}
	readField(taka)

	fmt.Println("================================================================================================================")
	readField(Sample{ Name: "Moe" })
}
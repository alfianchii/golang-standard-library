package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type User struct {
	Name string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Age int `required:"true" max:"10"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)

	fmt.Println("Type name:", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var field reflect.StructField = valueType.Field(i)
		var fieldRequiredTagValue string = field.Tag.Get("required")
		var fieldMaxTagValue string = field.Tag.Get("max")
		
		fmt.Println(field.Name, "w/", field.Type, "type and its tags:", fieldRequiredTagValue, fieldMaxTagValue)
	}
}

func isValid(value interface{}) bool {
	var t reflect.Type = reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			if data == "" {
				return false
			}
		}
	}

	return true
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
	fmt.Println("----------------")

	var taka User = User{
		Name: "Taka",
		Address: "Indonesia",
		Age: 69,
	}
	readField(taka)

	fmt.Println("----------------")
	readField(Sample{ Name: "Moe" })
	fmt.Println("================================================================================================================")

	var galih User = User{
		Name: "Galih",
		Address: "",
		Age: 9999,
	}

	fmt.Println(isValid(galih))
}
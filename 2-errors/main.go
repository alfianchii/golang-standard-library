package main

import (
	"errors"
	"fmt"
)

var (
	ErrValidation = errors.New("validation error")
	ErrNotFound = errors.New("not found error")
)

func getById(id string) error {
	if id == "" {
		return ErrValidation
	}

	if id != "alfian" {
		return ErrNotFound
	}

	return nil
}

func main() {
	var username string = "taka"

	var err error = getById(username)
	if err != nil {
		if errors.Is(err, ErrValidation) {
			fmt.Println("VALIDATION ERROR!")
		} else if errors.Is(err, ErrNotFound) {
			fmt.Println("NOT FOUND ERROR!")
			} else {
			fmt.Println("UNKNOWN ERROR!")
		}
	}
}
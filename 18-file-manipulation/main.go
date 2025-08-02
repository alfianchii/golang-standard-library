package main

import (
	"os"
)

func createLogFile(name string, message string) error {
	const fileExt string = ".log"
	var fileName string = name + fileExt

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(message)
	return nil
}

func main() {
	createLogFile("OK", "The program runs smoothly, actually.")
}
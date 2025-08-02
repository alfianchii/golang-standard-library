package main

import (
	"bufio"
	"fmt"
	"io"
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

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}

	defer file.Close()

	var buffMessage string
	var buffReader *bufio.Reader = bufio.NewReader(file)

	for {
		line, _, err := buffReader.ReadLine()

		if err == io.EOF {
			break
		}

		buffMessage += string(line) + "\n"
	}

	return buffMessage, nil
}

func addTextToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(message)

	return nil
}

func main() {
	createLogFile("OK", "The program runs smoothly, actually.")

	result, _ := readFile("OK.log")
	fmt.Println(result)

	addTextToFile("OK.log", "\nUpdated!")
}
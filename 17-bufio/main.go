package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	var stringReader *strings.Reader = strings.NewReader("This is long string\nwith new line\n")
	var buffReader *bufio.Reader = bufio.NewReader(stringReader)

	for {
		line, _, err := buffReader.ReadLine()

		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}

	fmt.Println(string([]byte{65, 97}))
	fmt.Println("================================================================================================================")
}
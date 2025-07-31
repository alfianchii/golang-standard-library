package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

	var buffWriter = bufio.NewWriter(os.Stdout) // My terminal
	buffWriter.WriteString("Hello!\n")
	buffWriter.WriteString("Happy learning~\n")
	buffWriter.Flush()
}
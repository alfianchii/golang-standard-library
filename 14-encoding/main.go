package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
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

	var csvString string = "alfian,taka\n" +
		"galih,gacor\n" +
		"moe,poi"
	var reader *csv.Reader = csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
	fmt.Println("----------------")

	var writer *csv.Writer = csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"alfian", "taka"})
	_ = writer.Write([]string{"galih", "gacor"})
	_ = writer.Write([]string{"moe", "poi"})
	writer.Flush()

	_ = writer.WriteAll([][]string{
		{"alfian", "taka"},
		{"galih", "gacor"},
		{"moe", "poi"},
	})
	writer.Flush()
}
package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`t([a-z])([a-z])a`)

	fmt.Println(regex.MatchString("taka"))
	fmt.Println(regex.MatchString("teka"))
	fmt.Println(regex.MatchString("toki"))

	fmt.Println(regex.FindAllString("taka toki toko tiki teka toKa teko toke tana tema", 2))
}
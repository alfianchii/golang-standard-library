package main

import "fmt"

func main() {
	// Bit overflow example
	var num int32 = 50;
	fmt.Println(num)

	var num64 int64 = int64(num)
	fmt.Println(num64)

	var num8 int8 = int8(num)
	fmt.Println(num8)

	var largeNum int32 = 300
	var largeNum8 int8 = int8(largeNum)
	fmt.Println(largeNum, largeNum8)
}
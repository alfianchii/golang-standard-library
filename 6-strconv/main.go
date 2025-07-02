package main

import (
	"fmt"
	"strconv"
)

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
	fmt.Println("================================================================================================================")

	// Strconv
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error", err.Error())
	}

	parseInt, err := strconv.ParseInt("120", 10, 8)
	if err == nil {
		fmt.Println(parseInt)
	} else {
		fmt.Println("Error", err.Error())
	}

	atoi, err := strconv.Atoi("1000")
	if err == nil {
		fmt.Println(atoi)
	} else {
		fmt.Println("Error", err.Error())
	}

	var fmtInt string = strconv.FormatInt(999, 2)
	fmt.Println(fmtInt)

	var fmtIntItoa string = strconv.Itoa(199031)
	fmt.Println(fmtIntItoa)
}
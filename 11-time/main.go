package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Println(now)

	var utc time.Time = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(utc.Local())

	var formatter string = "2006-01-02 15:04:05"
	var value string = "2020-10-10 10:10:10"
	// var value string = "Hutao"
	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(valueTime)
	}

	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())
	fmt.Println(valueTime.Hour())
	fmt.Println("================================================================================================================")

	var dur1 time.Duration = time.Second * 100
	var dur2 time.Duration = time.Millisecond * 1000
	var dur3 time.Duration = dur1 - dur2
	
	fmt.Printf("Duration 3: %d\n", dur3)
	fmt.Println("Duration 3:", dur3)
}
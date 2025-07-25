package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	var data *ring.Ring = ring.New(5)
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}
	
	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
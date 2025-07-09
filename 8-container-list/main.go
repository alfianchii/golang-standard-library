package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()
	data.PushBack("Alfian")
	data.PushBack("Taka")

	var dataHead *list.Element = data.Front()
	fmt.Println(dataHead.Value)
	
	var next *list.Element = dataHead.Next()
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
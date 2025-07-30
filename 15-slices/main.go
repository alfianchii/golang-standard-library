package main

import (
	"fmt"
	"slices"
)

func main() {
	var names []string = []string{"Taka", "Moe", "Galih"}
	var nums []int = []int{100, 90, 80, 10}
	
	fmt.Println(slices.Min(nums))
	fmt.Println(slices.Max(nums))
	fmt.Println(slices.Contains(names, "Taka"))
	fmt.Println(slices.Index(names, "Moe"))
}
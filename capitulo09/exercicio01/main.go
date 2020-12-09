package main

import (
	"fmt"
)

func main() {

	array := [5]int{10, 20, 30, 40, 50}
	fmt.Println(array)
	for _, valor := range array {
		fmt.Println(valor)
	}
	fmt.Printf("%T", array)
}

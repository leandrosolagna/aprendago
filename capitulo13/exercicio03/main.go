package main

import (
	"fmt"
)

func main() {
	x := 10
	y := 20
	total := x + y
	defer fmt.Println(total)
	fmt.Println("Hello, playground")
}

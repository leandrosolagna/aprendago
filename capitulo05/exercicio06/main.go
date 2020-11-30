package main

import (
	"fmt"
)

const (
	_ = iota + 2020
	x1
	x2
	x3
	x4
)

func main() {
	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)
	fmt.Println(x4)
}

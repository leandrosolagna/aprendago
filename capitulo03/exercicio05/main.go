package main

import (
	"fmt"
)

type solagna int

var x solagna
var y int

func main() {
	fmt.Printf("value: %v\n", x)
	fmt.Printf("type: %T\n", x)
	x = 42
	fmt.Printf("%v\n", x)

	y = int(x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%T", y)
}

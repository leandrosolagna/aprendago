package main

import (
	"fmt"
)

type solagna int

var x solagna

func main() {
	fmt.Printf("value: %v\n", x)
	fmt.Printf("type: %T\n", x)
	x = 42
	fmt.Printf("%v", x)
}

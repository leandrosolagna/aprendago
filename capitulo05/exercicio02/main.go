package main

import (
	"fmt"
)

var x bool

func main() {
	x = (100 == 100)
	fmt.Println(x)
	x = (100 != 100)
	fmt.Println(x)
	x = (200 <= 10)
	fmt.Println(x)
	x = (235 < 1235)
	fmt.Println(x)
	x = (100 >= 10)
	fmt.Println(x)
	x = (123214 > 12343)
	fmt.Println(x)
}

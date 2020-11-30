package main

import (
	"fmt"
)

func main() {

	var x int = 50
	fmt.Printf("Valor em decimal: %d\n", x)
	fmt.Printf("Valor em binario: %b\n", x)
	fmt.Printf("Valor em hexadecimal: %#x\n\n", x)

	y := x << 1
	fmt.Printf("Valor em decimal: %d\n", y)
	fmt.Printf("Valor em binario: %b\n", y)
	fmt.Printf("Valor em hexadecimal: %#x", y)
}

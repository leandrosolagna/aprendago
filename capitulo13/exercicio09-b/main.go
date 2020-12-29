package main

import (
	"fmt"
)

func main() {
	y := retornafuncao(soma, 5)
	fmt.Println(y)
}

func soma(x int) int {
	outro := 10 + x
	return outro
}

func retornafuncao(f func(int) int, z int) int {
	total := f(z)
	return total
}

package main

import (
	"fmt"
)

func main() {
	funcaoa := retornafuncao()

	y := funcaoa(2020)

	fmt.Println(y)
}

func retornafuncao() func(int) int {
	return func(i int) int {
		return i + 80
	}
}

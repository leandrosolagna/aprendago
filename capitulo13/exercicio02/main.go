package main

import (
	"fmt"
)

func main() {
	x := funcao1(1, 2, 3, 4, 5, 6)

	y := []int{1, 2, 3, 4, 5, 6}

	totaly := funcao2(y)

	fmt.Println(x)
	fmt.Println(totaly)
}

func funcao1(x ...int) int {
	soma1 := 0
	for _, funcao1 := range x {
		soma1 += funcao1
	}
	return soma1
}

func funcao2(y []int) int {
	soma2 := 0
	for _, funcao2 := range y {
		soma2 += funcao2
	}
	return soma2
}

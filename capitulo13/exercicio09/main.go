package main

import (
	"fmt"
)

func main() {
	y := retornafuncao(soma, "Solagna")
	fmt.Println(y)
}

func soma(x string) string {
	outro := "Leandro " + x
	return outro
}

func retornafuncao(f func(string) string, z string) string {
	total := f(z)
	return total
}

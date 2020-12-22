package main

import (
	"fmt"
)

func main() {
	fmt.Println(funcao1())
	fmt.Println(funcao2())
}

func funcao1() int {
	return 10
}

func funcao2() (int, string) {
	return 13, "Misery Signals"
}

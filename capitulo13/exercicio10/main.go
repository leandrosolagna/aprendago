package main

import (
	"fmt"
)

func main() {
	s := umafunc("Leandro")
	fmt.Println(s())
}

func umafunc(x string) func() string {
	sobrenome := " Solagna"
	return func() string {
		completo := x + sobrenome
		return completo
	}
}

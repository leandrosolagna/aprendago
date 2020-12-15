package main

import (
	"fmt"
)

type pessoa struct {
	Nome      string
	Sobrenome string
	Sabores   []string
}

func main() {
	pessoa1 := pessoa{
		Nome:      "Leandro",
		Sobrenome: "Solagna",
		Sabores:   []string{"tutti frutti", "hortela"},
	}
	fmt.Println(pessoa1.Nome, pessoa1.Sobrenome)
	for _, v := range pessoa1.Sabores {
		fmt.Println(v)
	}
	pessoa2 := pessoa{
		Nome:      "Maria",
		Sobrenome: "Kall",
		Sabores:   []string{"chocolate", "morango"},
	}
	fmt.Println(pessoa2.Nome, pessoa2.Sobrenome)
	for _, val := range pessoa2.Sabores {
		fmt.Println(val)
	}
}

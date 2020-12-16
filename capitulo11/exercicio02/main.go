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

	mapa := map[string]pessoa{
		"Solagna": pessoa{
			Nome:      "Leandro",
			Sobrenome: "Solagna",
			Sabores:   []string{"tutti frutti", "hortela"},
		},
		"Kall": pessoa{
			Nome:      "Maria",
			Sobrenome: "Kall",
			Sabores:   []string{"chocolate", "morango"},
		},
	}

	for _, value := range mapa {
		fmt.Println(value.Nome)
		for _, value := range value.Sabores {
			fmt.Println(value)
		}
	}
}

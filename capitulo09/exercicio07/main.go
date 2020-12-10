package main

import (
	"fmt"
)

func main() {

	slice := [][]string{
		[]string{
			"Leandro",
			"Solagna",
			"Video game"},
		[]string{
			"Tania",
			"Solagna",
			"Limpar"},
		[]string{
			"Rodrigo",
			"Solagna",
			"Cozinhar"},
	}

	for _, indice := range slice {
		fmt.Println(indice[0])
		for _, valor := range indice{
			fmt.Println("\t", valor)
		}
	}
}

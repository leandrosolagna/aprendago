package main

import (
	"fmt"
)

func main() {
	mapa := map[string][]string{
		"Solagna_leandro": []string{"Video game"},
		"Solagna_rodrigo": []string{"Cozinhar"},
		"Solagna_tania":   []string{"Limpar"},
	}

	mapa["Solagna_elaine"] = []string{"Trabalhar"}

	for key, value := range mapa {
		fmt.Println(key)
		for i, hobby := range value {
			fmt.Println("\t", i, hobby)
		}
	}
}

package main

import (
	"fmt"
)

func main() {

	anonimo := struct {
		mapa  map[string]string
		slice []int
	}{
		mapa: map[string]string{
			"umacor":   "azul",
			"outracor": "vermelhor",
		},
		slice: []int{1, 2},
	}
	fmt.Println(anonimo)
}

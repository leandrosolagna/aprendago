package main

import (
	"fmt"
)

func main() {

	var esporteFavorito string = "Hockey"

	switch esporteFavorito {
	case "Hockey":
		fmt.Println("O esporte favorito eh Hockey.")
	case "Futebol":
		fmt.Println("O esporte favorito eh Futebol.")
	case "Curling":
		fmt.Println("O esporte favorito eh Curling.")
	}

}

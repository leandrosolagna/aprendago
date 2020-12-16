package main

import (
	"fmt"
)

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	tracaoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {
	truck := caminhonete{
		veiculo: veiculo{
			portas: 4,
			cor:    "preto",
		},
		tracaoNasQuatro: true,
	}
	fmt.Println(truck)

	carro := sedan{
		veiculo: veiculo{
			portas: 2,
			cor:    "azul",
		},
		modeloLuxo: false,
	}
	fmt.Println(carro)

	fmt.Printf("A caminhonete possui %v de portas\n", truck.portas)
	fmt.Printf("O sedan tem a cor %v.", carro.cor)
}

package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func main() {
	x := pessoa{
		nome:      "Leandro",
		sobrenome: "Solagna",
	}
	fmt.Println(x)
	mudeMe(&x)
	fmt.Println(x)
}

func mudeMe(p *pessoa) {
	(*p).nome = "Tania"
}

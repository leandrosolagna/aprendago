package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) dados() {
	fmt.Println("Nome completo:", p.nome, p.sobrenome, " Idade:", p.idade, "anos")
}

func main() {

	pessoa1 := pessoa{"Leandro", "Solagna", 30}
	pessoa1.dados()
}

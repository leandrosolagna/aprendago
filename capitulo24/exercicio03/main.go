package main

import (
	"fmt"
)

type erroEspecial struct {
	umaMensagem string
}

func (e erroEspecial) Error() string {
	return "Erro #12345214."
}

func main() {
	umErro := erroEspecial{"Um erro"}
	funcaoParaErro(umErro)
}

func funcaoParaErro(e error) {
	fmt.Println("Seguinte erro: ", e)
}

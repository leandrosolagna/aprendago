package main

import "fmt"

type pessoa struct {
	frase string
}

func (p *pessoa) falar() {
	fmt.Println(p.frase)
}

type humanos interface {
	falar()
}

func main() {
	fala := pessoa{"Eu nao sei o que dizer."}
	dizerAlgumaCoisa(&fala)
//	dizerAlgumaCoisa(fala) //nao funciona
}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Comeco do programa")

	wg.Add(2)

	go funcao1()
	go funcao2()

	wg.Wait()
	fmt.Println("Fim do programa")
}

func funcao1() {
	fmt.Println("Funcao 1")
	wg.Done()
}

func funcao2() {
	fmt.Println("Funcao 2")
	wg.Done()
}

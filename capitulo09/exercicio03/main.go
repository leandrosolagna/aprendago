package main

import (
	"fmt"
)

func main() {

	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println(slice)

	primeiroslice := slice[0:3]
	fmt.Println(primeiroslice)

	segundoslice := slice[4:]
	fmt.Println(segundoslice)

	terceiroslice := slice[1:7]
	fmt.Println(terceiroslice)

	quartoslice := slice[2:9]
	fmt.Println(quartoslice)

	desafio := slice[2:len(slice)-1]
	fmt.Println(desafio)

	fmt.Printf("%T", slice)
}

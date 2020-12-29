package main

import (
	"fmt"
)

func main() {
	funcaoa := 2020

	func(funcaoa int) {
		fmt.Println(funcaoa + 90)
	}(funcaoa)
}

package main

import (
	"fmt"
)

func main() {

	x := 5

	switch {
	case x == 10:
		fmt.Println("X eh igual a 10.")
	case x < 10:
		fmt.Println("X eh menor que 10.")
	case x > 10:
		fmt.Println("X eh maior que 10.")
	}

}

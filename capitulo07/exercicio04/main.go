package main

import (
	"fmt"
)

func main() {

	x := 1990

	for {
		if (x <= 2020) {
			fmt.Println(x)
			x++
		} else {
			break
		}
	}

}

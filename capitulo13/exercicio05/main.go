package main

import (
	"fmt"
)

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

func (q quadrado) area() {
	areaquadrado := q.lado * q.lado
	fmt.Println("Area do quadrado:", areaquadrado)
}

func (c circulo) area() {
	areacirculo := 3.14 * c.raio * c.raio
	fmt.Println("Area do circulo:", areacirculo)
}

type figura interface {
	area()
}

func info(f figura) {
	f.area()
}

func main() {
	x := quadrado{
		lado: 20,
	}

	y := circulo{
		raio: 30,
	}

	info(x)
	info(y)
}

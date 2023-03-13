package main

import "fmt"

type persona struct {
	nombre string
}

func (p *persona) hablar() {
	fmt.Println("Hola soy", p.nombre)
}

type humano interface {
	hablar()
}

func diAlgo(h humano) {
	h.hablar()
}

func main() {
	p1 := persona{
		nombre: "Eduar",
	}

	//diAlgo(p1)
	//diAlgo(&p1)
	p1.hablar()
}

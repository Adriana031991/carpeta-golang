package main

import (
	"fmt"

	"gitlab.com/eduar/go-programming/010-Ninja-nivel-trece/01/codigo-inicial/perrito"
)

type canino struct {
	nombre string
	edad   int
}

func main() {
	fido := canino{
		nombre: "Fido",
		edad:   perrito.Edad(10),
	}
	fmt.Println(fido)
	fmt.Println(perrito.EdadDos(20))
}

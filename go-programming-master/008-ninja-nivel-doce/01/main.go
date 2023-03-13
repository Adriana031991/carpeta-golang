package main

import (
	"fmt"

	"gitlab.com/eduar/go-programming/008-ninja-nivel-doce/01/perro"
)

type canino struct {
	nombre string
	edad   int
}

func main() {
	c1 := canino{
		nombre: "Chester",
		edad:   perro.Edad(2),
	}
	fmt.Println(c1)
}

package main

import (
	"fmt"

	"gitlab.com/eduar/go-programming/010-Ninja-nivel-trece/02/01-codigo-inicial/cita"
	"gitlab.com/eduar/go-programming/010-Ninja-nivel-trece/02/01-codigo-inicial/palabra"
)

func main() {
	fmt.Println(palabra.Conteo(cita.Cuando))

	for k, v := range palabra.ConteoUso(cita.Cuando) {
		fmt.Println(v, k)
	}
}

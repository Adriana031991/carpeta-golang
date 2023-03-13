package main

import (
	"fmt"

	"gitlab.com/eduar/go-programming/010-Ninja-nivel-trece/02/02-codigo-finalizado/cita"
	"gitlab.com/eduar/go-programming/010-Ninja-nivel-trece/02/02-codigo-finalizado/palabra"
)

func main() {
	fmt.Println(palabra.Conteo(cita.Cuando))

	for k, v := range palabra.ConteoUso(cita.Cuando) {
		fmt.Println(v, k)
	}
}

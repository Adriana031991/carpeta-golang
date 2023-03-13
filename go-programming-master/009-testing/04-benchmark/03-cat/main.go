package main

import (
	"fmt"
	"strings"

	"gitlab.com/eduar/go-programming/009-testing/04-benchmark/03-cat/mystr"
)

const s = "Nos preguntamos: ¿Quién soy yo para ser brillante, hermosa, talentosa, fabulosa? En realidad, ¿quién eres tú para no ser? Tu pequeñez no le sirve al mundo. No hay nada iluminado sobre ser menos de modo que otras personas no se sientan inseguras a tu alrededor. Todos estamos destinados a brillar, como lo hacen los niños. Nacimos para manifestar la gloria que está dentro de nosotros. No es solo en algunos de nosotros; está en todos y a medida que dejamos que nuestra propia luz brille, inconscientemente damos permiso a otras personas para hacer lo mismo. Al liberarnos de nuestro propio miedo, nuestra presencia libera automáticamente a los demás. - Marianne Williamson"

func main() {

	xs := strings.Split(s, " ")

	for _, v := range xs {
		fmt.Println(v)
	}

	fmt.Printf("\n%s\n", mystr.Cat(xs))
	fmt.Printf("\n%s\n\n", mystr.Join(xs))
}

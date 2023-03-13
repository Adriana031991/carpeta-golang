package ejercicios

import (
	"fmt"
)

func Eje10() {

	// loops

	for i := 0; i < 10; i++ {
		fmt.Printf("Exercise 10 ciclo for -> %v\n ", i)
	}

	word := `words`
	for i := 0; i < len(word); i++ {
		fmt.Printf("Exercise 10 ciclo anidado externo-> %v\n ", i)
		for j := 0; j < 5; j++ {
			fmt.Printf("\t Exercise 10 ciclo anidado interno-> \t%v\n ", j)

		}
	}


	for _, v := range word {
		fmt.Printf("Exercise 10 range -> \t %v \n", string(v))
		
	}

	// for infinito

	ind := 0

	for {
		if ind > 9 {
			break  //salir del bloque infinito
		}
		fmt.Println("Exercise 10 infinite ->", ind)
		ind ++
	}
	

}

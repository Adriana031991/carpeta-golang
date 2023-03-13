package ejercicios

import (
	"fmt"
)

func Eje11() {
	
	// loops break and continuos
		
	ind := 0

	for {
		ind ++
		if ind > 100 {
			break  //salir del bloque infinito
		}
		if ind % 2 != 0 {
			continue // cuando se hace continuo , las lineas de abajo no se leen
		}
		fmt.Println("Exercise 11 break and continue ->", ind )
	}
		
		

	// imprimiendo caracter

	for i := 0; i < 122; i++ {
		fmt.Printf("valor del caracter decimal %d\t, hexadecimal %#x\t, unicode codePoint %#U\n", i,i,i)
		
	}


}

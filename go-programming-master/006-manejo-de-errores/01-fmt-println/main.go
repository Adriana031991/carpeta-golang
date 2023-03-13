package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("sin-archivo.txt")
	if err != nil {
		fmt.Println("ocurrió un error", err)
		//		log.Println("err happened", err)
		//		log.Fatalln(err)
		//		panic(err)
	}
}

// Println formatea utilizando los formatos predeterminados para sus operandos y escribe a la salida estándar.

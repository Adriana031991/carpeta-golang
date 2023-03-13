package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("sin-archivo.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		//		log.Println("err happened", err)
		log.Fatalln(err)
		//		panic(err)
	}
}

func foo() {
	fmt.Println("Cuando os.Exit() es llamada, las funciones diferidas no corren")
}

/*
... las funciones de Fatal llaman a os.Exit(1) después de escribir el mensaje del log ...
*/

// Fatalln es quivalente a Println() seguido por una llamada a os.Exit(1).

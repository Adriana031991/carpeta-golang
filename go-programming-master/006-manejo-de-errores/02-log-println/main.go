package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("sin-archivo.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		log.Println("un error ocurrió", err)
		//		log.Fatalln(err)
		//		panic(err)
	}
}

/*
El paquete log implementa un paquete simple de logging ... escribe a standard error e imprime la fecha y hora de cada mensaje logueado...
*/

// log.Println llama a Output a imprimir a el logger estándar. Los argumentos son manejados de la misma manera que en fmt.Println.

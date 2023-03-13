package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("sin-archivo.txt")
	if err != nil {
		//		fmt.Println("un error ocurrió", err)
		log.Println("un error ocurrió", err)
		//		log.Fatalln(err)
		//		panic(err)
	}
	defer f2.Close()

	fmt.Println("Revisa el archivo log.txt en el directorio.")
}

// Println llama Output para imprimir al logger estándar. Los argumentos son manejados de igual manera que en fmt.Println.

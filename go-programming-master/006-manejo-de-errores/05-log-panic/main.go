package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("sin-archivo.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		//		log.Println("err happened", err)
		//		log.Fatalln(err)
		log.Panicln(err)
		//		panic(err)
	}
}

/*
Panicln es quivalente a Println() seguido por una llamada a panic().
*/

// Fatalln es equivalente a Println() seguido por una llamada a os.Exit(1).

package main

import (
	"os"
)

func main() {
	_, err := os.Open("sin-archivo.txt")
	if err != nil {
		//		fmt.Println("un error ocurrió", err)
		//		log.Println("un error ocurrió", err)
		//		log.Fatalln(err)
		//		log.Panicln(err)
		panic(err)
	}
}

// http://godoc.org/builtin#panic

package main

import "fmt"

func main() {
	saludar("rrr")
	r := sum(7, 9)
	fmt.Println("la suma es", r)
}

func saludar(nombre string) {
	fmt.Println("hola!, ", nombre)
}

func sum(dat1, dat2 int) int {
	return dat1 + dat2
}

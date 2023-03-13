package main

import "fmt"

func main() {
	nombre := "nombre"
	edad := 55

	fmt.Printf("hey hey %s tiene %d dias \n", nombre, edad)
	// si no se el tipo es %v
	fmt.Printf("hey hey %v tiene %v dias \n", nombre, edad)

	mensaje := fmt.Sprintf("hola, %s", nombre)
	fmt.Printf(mensaje)

	// para tomar un dato en la consola

	fmt.Printf("ingrese nombre: ")

	fmt.Scan(&nombre)

	fmt.Println("otro nombre: ", nombre)
}

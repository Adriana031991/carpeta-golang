package main

import "fmt"

func main() {
	dias := make(map[int]string)
	//agregar,,, ojo que ya nao es indice sino clave

	dias[0] = "lunes"
	dias[1] = "martes"
	dias[2] = "miercoles"
	dias[3] = "juevs"
	dias[4] = "viernes"

	//modificar

	dias[4] = "VierNes"

	fmt.Println(dias)

	//delete
	delete(dias, 1)

	fmt.Println(dias, len(dias))

	//nuevo mapas

	estudiantes := make(map[string][]int)

	estudiantes["Alex"] = []int{1, 2, 3, 4}
	estudiantes["Ro"] = []int{13, 22, 3, 34}
	estudiantes["Ra"] = []int{21, 2, 23, 43}

	fmt.Println(estudiantes)
	fmt.Println("dato especifico del map", estudiantes["Ro"][2])

}

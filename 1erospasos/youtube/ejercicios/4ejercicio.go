package ejercicios

import (
	"fmt"
)

type x int

var num x
var y int

func Eje4() {

	fmt.Println("ejercicio 4 ", num)
	num = 44
	fmt.Printf(" ejercicio 4 -  %T", num)

	fmt.Println("ejercicio 4 ", y)
	y = int(66)
	fmt.Println("ejercicio 4 ", y)

}

package ejercicios

import (
	"fmt"
)

// funciones callback


func Eje25() {

	a := []int{1, 2, 3, 4, 5}

	g := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi)-1] 
	}

	x := i(g, a)

	fmt.Printf("%T\n", x)
	fmt.Println("ejercicio 25", x)

}

func i (function func(xi []int) int, sliceInt []int) int {
	n := function(sliceInt)
	n++
	return n
}
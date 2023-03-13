package ejercicios

import (
	"fmt"
)

// funciones

func Eje16() {

	ii := []int { 1, 2, 3, 4, 5}

	a := c(ii...)
	b := d(ii)

	fmt.Println(a)
	fmt.Println(b)

}

func c(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func d(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total 
}

package main

import (
	"fmt"

	"gitlab.com/eduar/go-programming/010-Ninja-nivel-trece/03/01-codigo-inicial/matematicas"
)

func main() {
	xxi := gen()
	for _, v := range xxi {
		fmt.Println(matematicas.PromedioCentrado(v))
	}
}

func gen() [][]int {
	a := []int{1, 4, 6, 8, 100}
	b := []int{0, 8, 10, 1000}
	c := []int{9000, 4, 10, 8, 6, 12}
	d := []int{123, 744, 140, 200}
	e := [][]int{a, b, c, d}
	return e
}

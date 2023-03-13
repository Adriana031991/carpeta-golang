package main

import "fmt"

func main() {
	fmt.Println("suma de 5 + 7 es :", miSuma(5, 7))
	fmt.Println("suma de 1 + 3 es :", miSuma(1, 3))
	fmt.Println("suma de 4 + 6 es :", miSuma(4, 6))
}

func miSuma(xi ...int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}

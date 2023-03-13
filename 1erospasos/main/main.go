package main

import (
	"fmt"
	"reflect"
)

func main() {

	// variables
	var enter int = 2
	var decimal float64 = 2.5
	var text string = "jj"
	var boolean bool = true

	// enter = 2
	// decimal = 2.5
	// text = "hi"
	// boolean = false

	fmt.Println("variables->", enter, " - ", decimal, " - ", text, " - ", boolean)

	num := 99.22
	fmt.Println(num)

	// operadores-matematicos

	sum := 10 + 5
	res := 10 - 5
	multi := 10 * 5
	div := 10 / 5
	restDiv := 10 % 5

	fmt.Println("operadores-matematicos->", sum, " - ", res, " - ", multi, " - ", div, " - ", restDiv)

	// operadores incremento y decremento

	sum++
	fmt.Println("operadores incremento ->", sum)
	sum--
	fmt.Println("operadores decremento ->", sum)

	sum += 10
	fmt.Println("operadores incremento 2 ->", sum)

	// operadores relacionales

	fmt.Println("expresion mayor verdadera->", 10 > 5)
	fmt.Println("expresion menor falsa->", 10 < 5)
	fmt.Println("expresion igualdad->", 10 == 5)
	fmt.Println("expresion mayor o igual->", 10 >= 5)
	fmt.Println("expresion menor o igual->", 10 <= 5)

	// operadores logicos

	fmt.Println("operador logico && ->", 10 <= 5 && 10 >= 4)
	fmt.Println("operador logico || ->", 10 <= 5 || 10 >= 4)

	// condicional y control de flujo

	if num == 5 {
		fmt.Println("es igual")
	} else if num == 33 {
		fmt.Println("el numero es igual a 33")
	} else {
		fmt.Println("No es igual, el num es: ", num)

	}

	for i := 0; i <= 10; i++ {
		// fmt.Println("ciclo for ", i+1)
		if i <= 10 {
			fmt.Println("ciclo for ", i)

		}

	}

	nombre := "hhheeee"
	for i := 0; i < 7; i++ {
		fmt.Println("ciclo for 2", string(nombre[i]))

	}

	variables("pp")

	fmt.Println(sumar(5, 6))

	ejer()

	arr()

	struc()

}

func variables(nam string) {
	fmt.Println("hi", nam)
}

func sumar(a int, b int) int {

	c := a + b
	return c

}

func ejer() {

	name := []string{"nombre"}
	for i := 0; i < len(name); i++ {
		fmt.Println("caracter ", i)
	}

	for index, v := range name {
		fmt.Println("caracter 2", index, "-", v)

	}

	name2 := "nombre"
	for i := 0; i < len(name2); i++ {
		fmt.Println("letter ", string(name2[i]))
	}

	for index, v := range name2 {
		fmt.Println("letter 2", index, "-", string(v))

	}

	for i := 0; i < 20; i++ {

		if i%2 == 0 {
			fmt.Println("es par ", i)

		}
	}
}

func arr() {
	books := map[string]int{
		"maths":     5,
		"biology":   9,
		"chemistry": 6,
		"physics":   3,
	}
	for key, val := range books {
		fmt.Println("array", key, val)
	}

	// just key
	for key := range books {
		fmt.Println("key", key)
	}
	for _, val := range books {
		fmt.Println("val", val)
	}

}

func struc() {
	type Person struct {
		Name   string
		Age    int
		Gender string
		Single bool
	}

	ubay := Person{
		Name:   "John",
		Gender: "Female",
		Age:    17,
		Single: false,
	}
	values := reflect.ValueOf(ubay)
	types := values.Type()
	for i := 0; i < values.NumField(); i++ {
		fmt.Println(types.Field(i).Index[0], types.Field(i).Name, values.Field(i))
	}
}

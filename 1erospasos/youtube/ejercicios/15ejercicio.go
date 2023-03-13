package ejercicios

import (
	"fmt"
)

// funciones

func Eje15() {

	a := a()
	b, c := b()

	fmt.Println(a)
	fmt.Println(b, " y", c)

}

func a() int {
	return 42
}

func b() (int, string) {
	return 23, "hola"
}

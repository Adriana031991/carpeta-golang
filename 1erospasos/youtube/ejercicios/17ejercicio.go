package ejercicios

import (
	"fmt"
)

// funciones

func Eje17() {

	defer e()
	fmt.Println("primer espacio")
	// fmt.Println(b)

}

func e() {
	//defer ayuda a que una func se reaice cuando todo termine
	defer func() {
		fmt.Println("func diferida")
	}()
	fmt.Println("function")
}



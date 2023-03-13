package ejercicios

import (
	"fmt"
)

// funciones

// ciudadanos de primer orden

var z int
var g func() = func() {
	fmt.Println("g desde afuera")
}

func Eje23() {

	f := func() {
		for i := 0; i < 2; i++ {
			fmt.Println(i)
		}
	}

	g()
	f()
	fmt.Printf("%T\n", f)
	fmt.Println("list")
	fmt.Println(z)

}

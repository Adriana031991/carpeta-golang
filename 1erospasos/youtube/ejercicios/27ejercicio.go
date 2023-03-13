package ejercicios

import (
	"fmt"
)

// funciones
// clousures

func Eje27() {

	f := k()
	h := k()
	fmt.Printf("%T\n", f)
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())

}

func k() func() int {
	x:= 0
	return func() int {
		x++
		return x
	}
}




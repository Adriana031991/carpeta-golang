package ejercicios

import (
	"fmt"
)

// funciones
// clousures

func Eje26() {

	f := j()
	h := j()
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

func j() func() int {
	x:= 0
	return func() int {
		x++
		return x
	}
}
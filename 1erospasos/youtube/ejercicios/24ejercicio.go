package ejercicios

import (
	"fmt"
)

// funciones


func Eje24() {

	f := h()
	fmt.Printf("%T\n", f)
	fmt.Println(f())

}

func h() func() int {
	return func() int {
		return 45
	}
}
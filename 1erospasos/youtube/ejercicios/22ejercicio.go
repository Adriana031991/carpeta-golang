package ejercicios

import (
	"fmt"
)

// funciones




func Eje22() {

	func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	}()
	fmt.Println("list")
}

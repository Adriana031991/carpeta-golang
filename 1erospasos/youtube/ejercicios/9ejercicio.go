package ejercicios

import (
	"fmt"
)

func Eje9() {
	
	// iota
		const (
			a = 2022 + iota
			b = 2022 + iota
			c = 2022 + iota
			d = 2022 + iota
		)
		fmt.Printf("Exercise 9 -> %v\t %v\t %v\t %v\t ", a,b,c,d)
		
		
}

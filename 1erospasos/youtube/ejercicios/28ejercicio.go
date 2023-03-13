package ejercicios

import (
	"fmt"
)

// goroutines

type Human struct {
	name string
	surName string
	age int
}


func (h *Human) talk()  {
	fmt.Printf("Hi! I'm %s %s, and have %v years old\n", h.name, h.surName, h.age)
}

type Activities interface {
	talk1() 
}

func talk1(h *Human) {
	h.talk()
}




func Eje28() {


	h1 := &Human{
		name: "Ana",
		surName: "Franklin",
		age: 28,
	}

	talk1(h1)
	h1.talk()


}

// func l() func() int {
// 	x:= 0
// 	return func() int {
// 		x++
// 		return x
// 	}
// }




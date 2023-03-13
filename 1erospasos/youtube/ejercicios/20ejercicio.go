package ejercicios

import (
	"fmt"
)

// funciones



func (p Person1) h() {
	fmt.Println("hello i'm:", p.name, p.surName, "and have", p.age)
	// return `hello i'm: p.name, and I have p.age`
}

type share1 interface {
	h() 
}



func Eje20() {

	p2 := Person1{
		name:    "Maria",
		surName: "vallejo",
		age:     22,
		address: "address",
	}

	

	p2.h()
}

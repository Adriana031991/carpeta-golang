package ejercicios

import (
	"fmt"
)

// funciones

type Person1 struct {
	name      string
	surName string  
	age       int
	address   string
	iceCreams []string
}

func (p Person1) f() {
	fmt.Println("hello i'm:", p.name, "and have", p.age)

}

func Eje18() {

	p2 := Person1{
		name:    "Maria",
		age:     22,
		address: "address",
	}

	p2.f()
}

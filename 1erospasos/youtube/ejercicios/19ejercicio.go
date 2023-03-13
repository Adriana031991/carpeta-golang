package ejercicios

import (
	"fmt"
)

// funciones



func (p Person1) g() string{
	// fmt.Println("hello i'm:", p.name, "and have", p.age)
	return `hello i'm: p.name, and I have p.age`
}

type share interface {
	g() string
}

func info(s share) {
	fmt.Println(s.g())
}

func Eje19() {

	p2 := Person1{
		name:    "Maria",
		age:     22,
		address: "address",
	}

	

	info(p2)
}

package ejercicios

import (
	"fmt"
	"math"
)

// funciones


type Square struct {
	longitude float64

}

type Circle struct {
	radio float64
}


func (p Square) h() float64 {
	return p.longitude* p.longitude
}

func (p Circle) h() float64 {
	return math.Pi * (p.radio * p.radio)
}

type Form interface {
	h() float64
}

func info1(f Form){
	fmt.Println("el area es:", f.h())

}

func Eje21() {

	p2 := Circle {
		radio: 99,
	}

	p3 := Square {
		longitude: 0030211,
	}
	

	info1(p2)
	info1(p3)
}

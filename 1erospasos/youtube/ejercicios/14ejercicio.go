package ejercicios

import (
	"fmt"
)

type vehicle struct {
	doors       int
	color   string
	
}

type car struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxuries bool
}

func Eje14() {

	c := car {
		vehicle: vehicle {doors: 2, color: "white",},
		fourWheel: true,
	}
	
	s := sedan {
		vehicle: vehicle {doors: 4, color: "black",},
		luxuries: true,
	}

	fmt.Println(c)
	fmt.Println(s)

	
}

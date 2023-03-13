package ejercicios

import (
	"fmt"
)

type Person struct {
	name      string
	age       int
	address   string
	iceCreams []string
}

func Eje13() {

	p := Person{
		name:      "aa",
		age:       9,
		address:   "address",
		iceCreams: []string{"vanilla", "cafe", "coffee", "strawberry"},
	}

	p2 := Person{
		name:      "bb",
		age:       10,
		address:   "address",
		iceCreams: []string{"vanilla", "strawberry"},
	}

	m := map[string]Person{
		p.name:  p,
		p2.name: p2,
	}

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v)

	}

	fmt.Println(p.name)
	fmt.Println(p.age)
	for _, v := range p.iceCreams {
		fmt.Println(v)

	}

	fmt.Println(p2.name)
	fmt.Println(p2.age)
	for _, v := range p2.iceCreams {
		fmt.Println(v)

	}


	// struct anonimo ejem
	// ex := struct {
	// 	name      string
	// 	age       int
	// 	address   string
	// 	iceCreams []string
	// } {
	// 	name      :"string",
	// 	age       :2,
	// 	address   :"string",
	// 	iceCreams :[]string {"1", "2"},
	// }

}

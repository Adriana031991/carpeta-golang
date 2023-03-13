package ejercicios

import (
	"fmt"
)

// goroutines


func Eje33() {

	salir := make(chan int)
	c := gen(salir)

	recibir(c, salir)

	fmt.Println("A punto de finalizar.")
}

func recibir(cr <-chan int, salir chan int) {
	for {
		select {
		case v := <-cr:
			fmt.Println(v)
		case v := <-salir:
			fmt.Println(v)
			return
		}
	}
}

func gen(salir chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		salir <- 1
	}()

	return c
}

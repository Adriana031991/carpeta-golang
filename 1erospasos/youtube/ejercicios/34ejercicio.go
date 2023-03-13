package ejercicios

import (
	"fmt"
	"runtime"
)

// goroutines


func Eje34() {

		a := 10
		b := 10
		g := gen1(a, b)
	
		recibir1(g)
	
		fmt.Println("A punto de finalizar.")
	}
	
	func recibir1(c <-chan int) {
		//for v := range c {
		//	fmt.Println(v + 1)
		//}
		for j := 0; j < 100; j++ {
			fmt.Println(j, <-c)
		}
		fmt.Println("GoRutinas", runtime.NumGoroutine())
	}
	
	func gen1(j, i int) <-chan int {
		c := make(chan int)
	
		for j := 0; j < 10; j++ {
			go func() {
				for i := 0; i < 10; i++ {
					c <- i
				}
	
			}()
			fmt.Println("GoRutinas", runtime.NumGoroutine())
	
		}
	
		return c
	
	}
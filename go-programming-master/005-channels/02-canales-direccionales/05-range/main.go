package main

import "fmt"

func main() {
	c := make(chan int)

	//enviar
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	//recibir
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Finalizando.")
}

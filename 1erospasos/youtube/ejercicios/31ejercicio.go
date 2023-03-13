package ejercicios

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// goroutines

//atomic

func Eje31() {

	var wg sync.WaitGroup //establezco el wait group para que espere la goroutin
	var count int64

	gs := 100
	wg.Add(gs) //establezco la cantidad de goroutines

	for i := 0; i < gs; i++ {
		go func() {

			atomic.AddInt64(&count, 1) 
			/* 
			2do parametro se llama delta
			delta es el numero de veces que se cambia la variable del primer parametro*/

			fmt.Println( atomic.LoadInt64(&count))
			
			wg.Done() // termina la goroutine
		}()
	}
	wg.Wait()
	fmt.Println("the valor final is:", count)
	
}

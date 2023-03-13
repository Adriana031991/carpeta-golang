package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incremento := 0
	gs := 100
	wg.Add(gs)
	var m sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			v := incremento
			v++
			incremento = v
			fmt.Println(incremento)
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("El valor final de incremento es:", incremento)
}

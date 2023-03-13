package ejercicios

import (
	"fmt"
	"runtime"
	"sync"
)

// goroutines

// mutex

func Eje30() {

fmt.Println("Number of CPUs:", runtime.NumCPU())
fmt.Println("Number of Goroutines:", runtime.NumGoroutine())

count := 0

const gs = 100

var wg sync.WaitGroup //establezco el wait group para que espere la goroutin

wg.Add(gs) //establezco la cantidad de goroutines

var m sync.Mutex 

for i := 0; i < gs; i++ {
	go func() {
		m.Lock()
		v := count
		v++
		count = v
		m.Unlock()
		fmt.Println("increment:", count)
		wg.Done() // termina la goroutine
	}()
}
wg.Wait()
fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}






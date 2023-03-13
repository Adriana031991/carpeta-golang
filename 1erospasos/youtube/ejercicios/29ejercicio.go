package ejercicios

import (
	"fmt"
	"runtime"
	"sync"
)

// goroutines

//race condition

func Eje29() {

fmt.Println("Number of CPUs:", runtime.NumCPU())
fmt.Println("Number of Goroutines:", runtime.NumGoroutine())

count := 0

const gs = 100

var wg sync.WaitGroup //establezco el wait group para que espere la goroutin

wg.Add(gs) //establezco la cantidad de goroutines

for i := 0; i < gs; i++ {
	go func() {
		v := count
		v++
		runtime.Gosched() //cede el procesador
		count = v
		wg.Done() // termina la goroutine
	}()
}
wg.Wait()
fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}

//usar go run -race main.go -race



// func l() func() int {
// 	x:= 0
// 	return func() int {
// 		x++
// 		return x
// 	}
// }




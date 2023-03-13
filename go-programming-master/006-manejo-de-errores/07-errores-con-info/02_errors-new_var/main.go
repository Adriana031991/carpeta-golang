package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrMat = errors.New("de matemática elemental: no hay raíz cuadrada real de un número negativo")

func main() {
	fmt.Printf("%T\n", ErrMat)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrMat
	}
	return 42, nil
}

// lee sobre el uso de errors.New en la biblioteca estándar:
// http://golang.org/src/pkg/bufio/bufio.go
// http://golang.org/src/pkg/io/io.go

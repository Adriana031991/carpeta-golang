package main

import (
	"fmt"
	"log"
)

type ubicacionError struct {
	lat  string
	long string
	err  error
}

func (n ubicacionError) Error() string {
	return fmt.Sprintf("un error de concepto matemático ha ocurrido en: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("Error mat: raíz cuadrada de número negativo: %v", f)
		return 0, ubicacionError{"50.2289 N", "99.4656 W", nme}
	}
	return 42, nil
}

// lerr sobre el uso structs con tipo error en la biblioteca estándard:
//
// http://golang.org/pkg/net/#OpError
// http://golang.org/src/pkg/net/dial.go
// http://golang.org/src/pkg/net/net.go
//
// http://golang.org/src/pkg/encoding/json/decode.go

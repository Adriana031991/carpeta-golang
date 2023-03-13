package main

import (
	"fmt"

	"gitlab.com/eduar/go-programming/009-testing/03-examples/01/mate"
)

func main() {
	fmt.Println(mate.Sum(2, 4, 5))
	fmt.Println(mate.Sum(4, 7, 8, 0))
}

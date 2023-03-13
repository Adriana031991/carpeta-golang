package ejercicios

import (
	"fmt"
)

func Eje12() {

	x := 1991

	for {
		if x > 2000 {
			break
		}
		fmt.Println("x", x)
		x++
	}

	for i := 0; i <= 10; i++ {
		fmt.Printf("Si dividimos %v entre 4, el resto (modulo) es %v\n", i, i%4)
	}

	// declaraciones if

	y := ""
	if y == "orale" {
		fmt.Println("muuuu")
	} else if y == "potro" {
		fmt.Println("relincho")
	} else {
		fmt.Println("ninguno")
	}

	//switch

	switch  {
	case false:
		fmt.Println("ninguno")
	case true:
		fmt.Println("uno")

	default:
		fmt.Println("epa")
	}


	deporteFav := "futbol"
	
	switch deporteFav {
	case "futbol":
		fmt.Println("futbol")
	case "baseball":
		fmt.Println("baseball")

	default:
		fmt.Println("cross")
	}

}

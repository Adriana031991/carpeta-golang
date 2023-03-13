package matematicas

import (
	"fmt"
	"testing"
)

func TestPromedioCentrado(t *testing.T) {
	type prueba struct {
		datos     []int
		respuesta float64
	}

	pruebas := []prueba{
		prueba{[]int{1, 2, 3, 4, 5}, 3},
		prueba{[]int{3, 5, 7, 8}, 6},
		prueba{[]int{10, 20, 30, 40, 50}, 30},
		prueba{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5},
	}

	for _, v := range pruebas {
		x := PromedioCentrado(v.datos)
		if x != v.respuesta {
			t.Error("Esperaba", v.respuesta, "Obtuvo", x)
		}
	}
}

func ExamplePromedioCentrado() {
	fmt.Println(PromedioCentrado([]int{1, 2, 3, 4, 5}))
	//Output:
	//3
}

func BenchmarkPromedioCentrado(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PromedioCentrado([]int{1, 2, 3, 4, 5, 6, 7, 1000000})
	}
}

package main

import "testing"

func TestMiSuma(t *testing.T) {

	type prueba struct {
		datos     []int
		respuesta int
	}

	pruebas := []prueba{
		prueba{[]int{1, 3, 5}, 9},
		prueba{[]int{0, 3, 7}, 10},
		prueba{[]int{1, 4, -5}, 0},
		prueba{[]int{0, 3, -3}, 0},
	}

	for _, x := range pruebas {
		v := miSuma(x.datos...)
		if v != x.respuesta {
			t.Error("Expected", x.respuesta, "Got", v)
		}
	}
}

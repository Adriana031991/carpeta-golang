package decir

import (
	"fmt"
	"testing"
)

func TestSaludar(t *testing.T) {
	s := Saludar("Eduar")
	if s != "Bienvenido querido Eduar" {
		t.Error("Expected", "Bienvenido querido Eduar", "Got", s)
	}
}

func ExampleSaludar() {
	fmt.Println(Saludar("Eduar"))
	//Output:
	//Bienvenido querido Eduar
}

func BenchmarkSaludar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Saludar("Eduar")
	}
}

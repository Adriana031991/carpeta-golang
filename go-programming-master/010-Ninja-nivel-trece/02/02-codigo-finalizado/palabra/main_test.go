package palabra

import (
	"fmt"
	"testing"

	"gitlab.com/eduar/go-programming/010-Ninja-nivel-trece/02/02-codigo-finalizado/cita"
)

func TestConteUso(t *testing.T) {
	m := ConteoUso("uno dos tres tres tres")
	for k, v := range m {
		switch k {
		case "uno":
			if v != 1 {
				t.Error("Esperaba", 1, "Obtuvo", v)
			}
		case "dos":
			if v != 1 {
				t.Error("Esperaba", 1, "Obtuvo", v)
			}
		case "tres":
			if v != 3 {
				t.Error("Esperaba", 3, "Obtuvo", v)
			}
		}
	}
}

func TestConteo(t *testing.T) {
	n := Conteo(cita.Cuando)
	if n != 355 {
		t.Error("Esperaba", 355, "Obtuvo", n)
	}
}

func ExampleConteo() {
	fmt.Println(Conteo("uno dos tres"))
	//Output:
	//3
}

func BenchmarkConteo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Conteo(cita.Cuando)
	}
}

func BenchmarkConteoUso(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConteoUso(cita.Cuando)
	}
}

// Package palabra nos ayuda a trabajar con strings
package palabra

import "strings"

// ConteoUso nos retorna un mapa con las palabras y las cantidades de veces que aparecen.
// No necesitas escribir un ejemplo para esta función
// escribir una prueba para esta es un bono retador; un poco más difícil
func ConteoUso(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Conteo nos da la cantidad total de palabras
func Conteo(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}

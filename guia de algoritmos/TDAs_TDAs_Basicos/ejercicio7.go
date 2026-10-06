package main

type colaDinamica[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

type nodoCola[T any] struct {
	dato      T
	siguiente *nodoCola[T]
}

func (c *colaDinamica[T]) Multiprimeros(k int) []T {
	arreglo := make([]T, 0)
	actual := c.primero
	for range k {
		if actual == nil {
			break
		}
		arreglo = append(arreglo, actual.dato)
		actual = actual.siguiente
	}
	return arreglo
}

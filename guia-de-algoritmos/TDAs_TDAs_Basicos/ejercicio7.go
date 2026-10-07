package main

type colaDinamicaa[T any] struct {
	primero *nodoColaa[T]
	ultimo  *nodoColaa[T]
}

type nodoColaa[T any] struct {
	dato      T
	siguiente *nodoColaa[T]
}

func (c *colaDinamicaa[T]) Multiprimeros(k int) []T {
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

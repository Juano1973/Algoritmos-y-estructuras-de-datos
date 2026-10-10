package main

import TDAcola "entrenamiento/tdas/cola"

func Multiprimeros[T any](cola TDAcola.Cola[T], k int) []T {

	arr := []T{}
	for range k {
		if cola.EstaVacia() {
			break
		}
		arr = append(arr, cola.Desencolar())
	}

	return arr
}

package main

import (
	TDApila "entrenamiento/tdas/pila"
)

func CambiarOrden[T any](Arreglo []T) []T {
	pila := TDApila.CrearPilaDinamica[T]()
	for i := 0; i < len(Arreglo); i++ {
		pila.Apilar(Arreglo[i])
	}
	nuevo_arreglo := make([]T, len(Arreglo))

	for j := 0; j < len(Arreglo); j++ {
		nuevo_arreglo[j] = pila.Desapilar()
	}
	return nuevo_arreglo
}

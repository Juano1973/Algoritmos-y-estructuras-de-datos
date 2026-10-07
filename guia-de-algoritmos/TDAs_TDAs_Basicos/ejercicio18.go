package main

import (
	TDApila "entrenamiento/tdas/pila"
	"fmt"
)

func cant_pila[T any](pila TDApila.Pila[T]) int {

	if pila.EstaVacia() {
		return 0
	}
	elem := pila.Desapilar()
	cantidad := cant_pila(pila)
	pila.Apilar(elem)
	return cantidad + 1
}

func main() {

	pila := TDApila.CrearPilaDinamica[int]()
	for i := range 100 {
		pila.Apilar(i)
	}
	pila.Desapilar()
	pila.Desapilar()
	pila.Desapilar()

	fmt.Println(cant_pila(pila))

}

package main

import (
	TDApila "entrenamiento/tdas/pila"
	"fmt"
)

// O(n²)
//Este ejercicio lo saque de otro repo de la misma materia, es jodido T_T

func main() {
	pila := TDApila.CrearPilaDinamica[int]()
	pila.Apilar(4)
	pila.Apilar(1)
	pila.Apilar(5)
	pila.Apilar(2)
	pila.Apilar(3)
	Ordenar(pila)
	for !pila.EstaVacia() {
		fmt.Println(pila.Desapilar())
	}
}

func Ordenar(pila TDApila.Pila[int]) {
	pilaAuxiliar := TDApila.CrearPilaDinamica[int]()
	for !pila.EstaVacia() {
		desapilado := pila.Desapilar()
		for !pilaAuxiliar.EstaVacia() && pilaAuxiliar.VerTope() < desapilado {
			pila.Apilar(pilaAuxiliar.Desapilar())
		}
		pilaAuxiliar.Apilar(desapilado)
	}
	for !pilaAuxiliar.EstaVacia() {
		pila.Apilar(pilaAuxiliar.Desapilar())
	}
}

package main

import Cola "entrenamiento/tdas/cola"

func FiltrarCola[K any](cola *Cola.Cola[K], filtro func(K) bool) {

	cola_aux := Cola.CrearColaEnlazada[K]()
	for !(*cola).EstaVacia() {
		v := (*cola).Desencolar()
		if filtro(v) {
			cola_aux.Encolar(v)
		}
	}

	cola = &cola_aux
}

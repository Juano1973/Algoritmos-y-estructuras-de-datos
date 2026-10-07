package main

import (
	TDApila "entrenamiento/tdas/pila"
)

func Esta_ordenado(p TDApila.Pila[int]) bool {
	actual := p.VerTope()
	for !p.EstaVacia() {
		act := p.Desapilar()
		if act < actual {
			return false
		}
		actual = act
	}

	return true
}

package main

import (
	TDAhash "entrenamiento/tdas/diccionario"
	TDApila "entrenamiento/tdas/pila"
)

func mergePilas(p_uno, p_dos TDApila.Pila[int]) []int {
	res := make([]int, 0)
	dicc := TDAhash.CrearHash[int, bool]()
	for !p_uno.EstaVacia() && !p_dos.EstaVacia() {
		if p_uno.VerTope() > p_dos.VerTope() {
			ingresante := p_dos.Desapilar()
			if dicc.Pertenece(ingresante) {
				continue
			}
			dicc.Guardar(ingresante, true)
			res = append(res, ingresante)
		} else {
			ingresante := p_uno.Desapilar()
			if dicc.Pertenece(ingresante) {
				continue
			}
			dicc.Guardar(ingresante, true)
			res = append(res, ingresante)
		}
	}
	for !p_uno.EstaVacia() {
		res = append(res, p_uno.Desapilar())
	}

	for !p_dos.EstaVacia() {
		res = append(res, p_dos.Desapilar())
	}

	return res
}

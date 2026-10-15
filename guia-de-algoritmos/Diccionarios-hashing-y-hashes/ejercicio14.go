package main

import TDAhash "entrenamiento/tdas/diccionario"

func sumanEsto(arreglo []int, largo int, k int) bool {

	if largo == 1 {
		return false
	}

	h := TDAhash.CrearHash[int, int]()
	for i := 0; i < largo; i++ {
		h.Guardar(k-arreglo[i], i)
	}

	for j := 0; j < largo; j++ {
		if h.Pertenece(arreglo[j]) && h.Obtener(arreglo[j]) != j {
			return true
		}
	}

	return false
}

package main

import (
	TDAheap "entrenamiento/tdas/cola_prioridad"
	TDAlista "entrenamiento/tdas/lista"
)

func comparar_cadenas(cadena_uno, cadena_dos string) int {

	if len(cadena_uno) > len(cadena_dos) {
		return 1
	}

	if len(cadena_uno) < len(cadena_dos) {
		return -1
	}

	return 0
}

func top_cadenas(arr []string, k int) TDAlista.Lista[string] {
	res := TDAlista.CrearListaEnlazada[string]()
	h := TDAheap.CrearHeapArr(arr, comparar_cadenas)
	for range k {
		res.InsertarUltimo(h.Desencolar())
	}
	return res
}

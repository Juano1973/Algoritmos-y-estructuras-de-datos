package main

import TDAlista "entrenamiento/tdas/lista"

func arbol_balanceado[K comparable](arr []K) TDAlista.Lista[K] {

	list := TDAlista.CrearListaEnlazada[K]()
	if len(arr) == 0 {
		return list
	}

	_balancear(arr, list, 0, len(arr))
	return list
}

func _balancear[K comparable](arr []K, l TDAlista.Lista[K], inicio, fin int) {

	medio := (inicio + fin) / 2
	l.InsertarUltimo(arr[medio])
	_balancear(arr, l, inicio, medio)
	_balancear(arr, l, medio+1, fin)
}

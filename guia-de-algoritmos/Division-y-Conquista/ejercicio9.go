package main

func estaOrdenado(arr []int, largo int) bool {
	return _estaOrdenado(arr, 0, largo)
}

func _estaOrdenado(arr []int, inicio, fin int) bool {
	if inicio >= fin {
		return true
	}

	medio := (inicio + fin) / 2
	if arr[medio] > arr[medio+1] {
		return false
	}

	izq := _estaOrdenado(arr, inicio, medio)
	der := _estaOrdenado(arr, medio+1, fin)

	return izq && der
}

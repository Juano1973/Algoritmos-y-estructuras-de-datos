package main

import (
	TDAhash "entrenamiento/tdas/diccionario"
	"fmt"
)

func masDeLaMitad(arr []int) bool {
	if len(arr) == 1 {
		return true
	}
	dict := TDAhash.CrearHash[int, int]()
	return _masDeLaMitad(arr, 0, len(arr), dict)
}

func _masDeLaMitad(arr []int, ini, fin int, h TDAhash.Diccionario[int, int]) bool {

	if ini >= fin {
		return false
	}

	medio := (ini + fin) / 2

	if !h.Pertenece(arr[medio]) {
		h.Guardar(arr[medio], 1)
	} else {
		cant := h.Obtener(arr[medio]) + 1
		h.Guardar(arr[medio], cant)
		if cant > len(arr)/2 {
			return true
		}
	}

	izq := _masDeLaMitad(arr, ini, medio, h)
	der := _masDeLaMitad(arr, medio+1, fin, h)
	return izq || der
}

func main() {
	arr := []int{1, 2, 2}
	fmt.Println(masDeLaMitad(arr))
}

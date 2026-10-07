package main

import "math"

func elemFueraDeLugar(arr []int) int {
	return _fuera_de_lugar(arr, 0, len(arr))
}

func _fuera_de_lugar(arr []int, inicio, fin int) int {
	if inicio >= fin {
		return int(math.Inf(-1))
	}

	medio := (inicio + fin) / 2

	if arr[medio] > arr[medio+1] {
		return arr[medio]
	}
	izq := _fuera_de_lugar(arr, inicio, medio)
	der := _fuera_de_lugar(arr, medio+1, fin)

	return max(izq, der)

}

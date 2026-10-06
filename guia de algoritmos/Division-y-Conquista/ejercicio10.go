package main

func elemFueraDeLugar(arr []int) int {
	return _fuera_de_lugar(arr, 0, len(arr))
}

func _fuera_de_lugar(arr []int, inicio, fin int) {
	if inicio >= fin {
		return arr[inicio]
	}
}

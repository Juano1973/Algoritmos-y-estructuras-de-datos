package main

func unos_ceros(arr []int) int {

	if arr[0] == 0 {
		return 0
	}
	return _unos_ceros(arr, 0, len(arr))
}

func _unos_ceros(arr []int, inicio, fin int) int {
	if inicio >= fin {
		return -1
	}

	medio := (inicio + fin) / 2

	if arr[medio] == 1 && arr[medio+1] == 0 {
		return medio + 1
	}

	if arr[medio] == 0 {
		return _unos_ceros(arr, inicio, medio)
	}
	return _unos_ceros(arr, medio+1, fin)
}

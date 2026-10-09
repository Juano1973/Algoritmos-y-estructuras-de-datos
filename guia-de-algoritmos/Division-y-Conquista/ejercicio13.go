package main

func encontrarPico(arr []int) int {
	return _encontrarPico(arr, 0, len(arr))
}

func _encontrarPico(arr []int, inicio, fin int) int {
	if inicio >= fin {
		return inicio
	}

	medio := (inicio + fin) / 2

	if arr[medio] < arr[medio+1] {
		return _encontrarPico(arr, medio+1, fin)
	}
	return _encontrarPico(arr, inicio, medio)
}

/*
func main() {
	arr := []int{1, 2, 3, 10, 50, 115, 3, 2, 1, 0}
	fmt.Println(encontrarPico(arr))
}
*/

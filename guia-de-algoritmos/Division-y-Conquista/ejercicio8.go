package main

func minimo(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	if len(arr) == 1 {
		return arr[0]
	}

	return _minimo(arr, 0, len(arr), arr[0])
}

func _minimo(arr []int, inicio, fin, minimo int) int {
	if inicio >= fin {
		return minimo
	}

	medio := (inicio + fin) / 2
	if arr[medio] < minimo {
		minimo = arr[medio]
	}

	izq := _minimo(arr, 0, medio, minimo)
	der := _minimo(arr, medio+1, fin, minimo)

	return min(izq, der)
}

/*
func main() {
	arr := []int{50, 88, 3, 765, 4, 334, 54, 5}
	fmt.Println(minimo(arr))
}
*/

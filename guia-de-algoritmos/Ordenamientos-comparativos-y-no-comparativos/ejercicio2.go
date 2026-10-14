package main

func ordenar(arr []int) []int {

	largo := len(arr)
	for indice, _ := range arr {
		arr[indice], arr[largo-indice] = arr[largo-indice], arr[indice]
	}

	for indice, elemnto := range arr {
		if indice == largo-1 {
			break
		}
		if elemnto > arr[indice+1] {
			arr[indice], arr[indice+1] = arr[indice+1], arr[indice]
		}
	}

	return arr
}

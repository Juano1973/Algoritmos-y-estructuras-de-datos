package main

import "strconv"

func countingSortt(arreglo []string, k int, f func(int) int) []string {

	frecuencias := make([]int, 0)

	for _, e := range arreglo {
		num, _ := strconv.Atoi(e)
		frecuencias[f(num)] += 1
	}

	inicios := make([]int, k)

	for i := 1; i < k; i++ {
		inicios[i] = inicios[i-1] + frecuencias[i-1]

	}

	ordenado := make([]string, len(arreglo))

	for j := 0; j < len(arreglo); j++ {
		elem := arreglo[j]
		num, _ := strconv.Atoi(elem)
		funcElem := f(num)
		pos := inicios[funcElem]
		inicios[funcElem] += 1
		ordenado[pos] = elem
	}

	return ordenado

}

//La mejor forma de ordenar lo planteado en el ejercicio 4 es un radixsort, usando un counting sort para cada una de las cosas por
//las cuales hay que ordenar

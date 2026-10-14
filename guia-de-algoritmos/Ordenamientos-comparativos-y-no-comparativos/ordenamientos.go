package main

import "strconv"

/*
La complejidad temporal (orden) del algoritmo Counting Sort es O(n + k) en todos los casos (mejor, promedio y peor caso), donde:
n es el número de elementos que tiene el array de entrada.
k es el rango de los valores posibles
(la diferencia entre el valor máximo y el mínimo, o el valor máximo si empieza en cero).
*/
func countingSort(arr []string, k int, f func(int) int) []string {
	frecuencias := make([]int, k)

	for _, e := range arr {
		num, _ := strconv.Atoi(e)
		frecuencias[f(num)] += 1
	}

	inicios := make([]int, k)

	for i := 1; i < k; i++ {
		inicios[i] = inicios[i-1] + frecuencias[i-1]
	}

	ordenado := make([]string, len(arr))

	for j := 0; j < len(arr); j++ {
		elem := arr[j]
		num, _ := strconv.Atoi(elem)
		funcElem := f(num)
		pos := inicios[funcElem]
		inicios[funcElem] += 1
		ordenado[pos] = elem

	}
	return ordenado
}

/*
Mejor caso: O(n log n)
Ocurre cuando el pivote elegido divide al arreglo en dos mitades aproximadamente iguales en cada paso recursivo.

Caso promedio: O(n log n)
Es el rendimiento habitual cuando los datos están desordenados de forma aleatoria o normal.

Peor caso: O(n^2))
Ocurre con listas ya ordenadas, inversamente ordenadas o cuando se hace una elección de pivote deficiente (por ejemplo, siempre elegir el primer o último elemento).
*/
func quickSort() {}

// La complejidad del algoritmo Merge Sort es O(n log n)
func Mergesort()

// La complejidad temporal del algoritmo Selection Sort (ordenación por selección) es O(n²)
func SelectionSort() {}

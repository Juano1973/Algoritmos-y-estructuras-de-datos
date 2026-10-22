package main

import TDAheap "entrenamiento/tdas/cola_prioridad"

func comparar(a, b elemento_arreglo) int {
	return b.valor - a.valor
}

type elemento_arreglo struct {
	valor   int
	indice  int
	arreglo int
}

func KMerge(arr [][]int) []int {

	if len(arr) == 1 {
		return arr[0]
	}

	res := make([]int, 0)

	if len(arr) == 0 {
		return res
	}

	h := TDAheap.CrearHeap(comparar)

	for i, _ := range arr {
		if len(arr[i]) == 0 {
			continue
		}
		elem := new(elemento_arreglo)
		elem.indice = 0
		elem.arreglo = i
		elem.valor = arr[i][0]
		h.Encolar(*elem)
	}

	for !h.EstaVacia() {
		elem := h.Desencolar()
		valor := elem.valor
		largo := len(arr[elem.arreglo])
		if elem.indice+1 < largo {
			elem.indice++
			elem.valor = arr[elem.arreglo][elem.indice]
			h.Encolar(elem)
		}
		res = append(res, valor)
	}

	return res
}

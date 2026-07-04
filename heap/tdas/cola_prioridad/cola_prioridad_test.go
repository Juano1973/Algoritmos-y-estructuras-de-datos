package cola_prioridad_test

import (
	TDAheap "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

func cmp(a, b int) int {
	return a - b
}

const _VOLUMEN = 100000

func TestEncolar(t *testing.T) {
	heap := TDAheap.CrearHeap(cmp)
	arreglo := []int{10, 1, 30, 50, 10, 20, 60, 90, 5}
	arreglo_ordenado := []int{90, 60, 50, 30, 20, 10, 10, 5, 1}
	for _, e := range arreglo {
		heap.Encolar(e)
	}

	require.Equal(t, 90, heap.VerMax())
	require.Equal(t, 9, heap.Cantidad())
	require.False(t, heap.EstaVacia())

	for _, e := range arreglo_ordenado {
		require.Equal(t, e, heap.Desencolar())
	}
}

func TestHeapify(t *testing.T) {
	arreglo := []int{1, 10, 3, 7, 8, 5, 4, 6, 2, 9}
	arreglo_ordenado := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	heap := TDAheap.CrearHeapArr(arreglo, cmp)

	for _, e := range arreglo_ordenado {
		require.Equal(t, e, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())

}

func TestHeapsort(t *testing.T) {
	arreglo := []int{67, 5, 7, 2, 8, 43, 76, 32, 87, 98}
	TDAheap.HeapSort(arreglo, cmp)
	require.Equal(t, []int{2, 5, 7, 8, 32, 43, 67, 76, 87, 98}, arreglo)
}

func TestVolumen(t *testing.T) {
	heap := TDAheap.CrearHeap(cmp)

	for j := 0; j <= _VOLUMEN; j++ {
		heap.Encolar(j)
	}
	for i := _VOLUMEN; i >= 0; i-- {
		require.Equal(t, heap.Desencolar(), i)
	}

}

func TestHeapVacio(t *testing.T) {
	heap := TDAheap.CrearHeap(cmp)

	require.Panics(t, func() { heap.Desencolar() })
	require.Panics(t, func() { heap.VerMax() })
	require.True(t, heap.EstaVacia())
}

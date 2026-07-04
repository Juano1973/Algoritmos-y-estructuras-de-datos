package cola_prioridad

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func cmp(a, b int) int {
	return a - b
}

func TestEncolar(t *testing.T) {
	heap := CrearHeap(cmp)

	heap.Encolar(10)
	heap.Encolar(1)
	heap.Encolar(30)
	heap.Encolar(50)
	heap.Encolar(10)
	heap.Encolar(20)
	heap.Encolar(60)
	heap.Encolar(90)
	heap.Encolar(5)

	require.Equal(t, 90, heap.VerMax())
	require.Equal(t, 9, heap.Cantidad())
	require.False(t, heap.EstaVacia())
	//
	require.Equal(t, 90, heap.Desencolar())
	require.Equal(t, 60, heap.Desencolar())
	require.Equal(t, 50, heap.Desencolar())
	require.Equal(t, 30, heap.Desencolar())
	require.Equal(t, 20, heap.Desencolar())
	require.Equal(t, 10, heap.Desencolar())
	require.Equal(t, 10, heap.Desencolar())
	require.Equal(t, 5, heap.Desencolar())
	require.Equal(t, 1, heap.Desencolar())
}

func TestHeapify(t *testing.T) {
	arreglo := []int{1, 10, 3, 7, 8, 5, 4, 6, 2, 9}
	heap := CrearHeapArr(arreglo, cmp)

	require.Equal(t, 10, heap.Desencolar())
	require.Equal(t, 9, heap.Desencolar())
	require.Equal(t, 8, heap.Desencolar())
	require.Equal(t, 7, heap.Desencolar())
	require.Equal(t, 6, heap.Desencolar())
	require.Equal(t, 5, heap.Desencolar())
	require.Equal(t, 4, heap.Desencolar())
	require.Equal(t, 3, heap.Desencolar())
	require.Equal(t, 2, heap.Desencolar())
	require.Equal(t, 1, heap.Desencolar())
	require.True(t, heap.EstaVacia())

}

func TestHeapsort(t *testing.T) {
	arreglo := []int{67, 5, 7, 2, 8, 43, 76, 32, 87, 98}
	HeapSort(arreglo, cmp)
	require.Equal(t, []int{2, 5, 7, 8, 32, 43, 67, 76, 87, 98}, arreglo)
}

func TestVolumen(t *testing.T) {
	heap := CrearHeap(cmp)
	var volumen int = 10000000

	for j := 0; j <= volumen; j++ {
		heap.Encolar(j)
	}
	for i := volumen; i >= 0; i-- {
		require.Equal(t, heap.Desencolar(), i)
	}

}

func TestHeapVacio(t *testing.T) {
	heap := CrearHeap(cmp)

	require.Panics(t, func() { heap.Desencolar() })
	require.Panics(t, func() { heap.VerMax() })
	require.True(t, heap.EstaVacia())
}

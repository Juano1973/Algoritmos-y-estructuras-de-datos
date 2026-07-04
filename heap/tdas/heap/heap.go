package cola_prioridad

const _PANICO_HEAP = "La cola esta vacia"
const _CAPACIDAD_INICIAL = 10

type heap[T any] struct {
	comparacion func(T, T) int
	arreglo     []T
	cant        int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{comparacion: funcion_cmp, arreglo: make([]T, _CAPACIDAD_INICIAL)}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {

	largo := len(arreglo)

	if largo < _CAPACIDAD_INICIAL {
		largo = _CAPACIDAD_INICIAL
	}

	arregloH := make([]T, largo)
	copy(arregloH, arreglo)
	heapify(arregloH, funcion_cmp)

	return &heap[T]{comparacion: funcion_cmp, arreglo: arregloH, cant: len(arreglo)}
}

func (heap *heap[T]) redimensionar(nuevoTam int) {
	nuevoArreglo := make([]T, nuevoTam)
	copy(nuevoArreglo, heap.arreglo)
	heap.arreglo = nuevoArreglo
}

func swap[T any](elemento1, elemento2 *T) {
	*elemento1, *elemento2 = *elemento2, *elemento1
}

func heapify[T any](arreglo []T, funcion_cmp func(T, T) int) {
	for i := (len(arreglo) - 1); i >= 0; i-- {
		downheap(arreglo, i, funcion_cmp, len(arreglo))
	}
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	heapify(elementos, funcion_cmp)
	for i := len(elementos) - 1; i > 0; i-- {
		swap(&elementos[0], &elementos[i])
		downheap(elementos[:i], 0, funcion_cmp, i)
	}
}

func (heap *heap[T]) EstaVacia() bool {
	return heap.cant == 0
}

func posicion_padre(indice int) int {
	return (indice - 1) / 2
}

func posicion_hijos(indice int) (int, int) {
	return 2*indice + 1, 2*indice + 2
}

func (heap *heap[T]) Encolar(elemento T) {
	if heap.cant == len(heap.arreglo)/2 {
		heap.redimensionar(len(heap.arreglo) * 2)
	}
	heap.arreglo[heap.cant] = elemento
	heap.cant++
	if heap.cant == 1 {
		return
	}
	var indice int = heap.cant - 1

	heap.upheap(indice)
}
func (heap *heap[T]) upheap(indice int) {
	if indice <= 0 {
		return
	}
	indice_padre := posicion_padre(indice)

	if heap.comparacion(heap.arreglo[indice_padre], heap.arreglo[indice]) < 0 {
		swap(&heap.arreglo[indice_padre], &heap.arreglo[indice])
		heap.upheap(indice_padre)
	}
}

func (heap *heap[T]) VerMax() T {
	if heap.EstaVacia() {
		panic(_PANICO_HEAP)
	}
	return heap.arreglo[0]
}

func (heap *heap[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic(_PANICO_HEAP)
	}
	dato := heap.arreglo[0]

	principio := &heap.arreglo[0]
	fin := &heap.arreglo[heap.cant-1]

	swap(principio, fin)

	if len(heap.arreglo) == 1 {
		var nuevoArreglo []T = make([]T, _CAPACIDAD_INICIAL)
		heap.cant--
		heap.arreglo = nuevoArreglo
		return dato
	}

	if heap.cant == len(heap.arreglo)/4 {
		heap.redimensionar(len(heap.arreglo) / 2)
	}

	heap.arreglo = heap.arreglo[:heap.cant-1]
	heap.cant--
	downheap(heap.arreglo, 0, heap.comparacion, heap.cant)

	return dato
}
func encontrar_maximo[T any](indice, cantidad int, arreglo []T, comparacion func(T, T) int) int {
	IndiceIzq, IndiceDer := posicion_hijos(indice)
	maximo := indice
	if IndiceDer < cantidad && comparacion(arreglo[maximo], arreglo[IndiceDer]) < 0 {
		maximo = IndiceDer
	}
	if IndiceIzq < cantidad && comparacion(arreglo[maximo], arreglo[IndiceIzq]) < 0 {
		maximo = IndiceIzq
	}
	return maximo
}
func downheap[T any](arreglo []T, indice int, comparacion func(T, T) int, cantidad int) {
	hijoMasGrande := encontrar_maximo(indice, cantidad, arreglo, comparacion)

	if hijoMasGrande == indice {
		return
	}

	swap(&arreglo[hijoMasGrande], &arreglo[indice])
	downheap(arreglo, hijoMasGrande, comparacion, cantidad)

}

func (heap *heap[T]) Cantidad() int {
	return heap.cant
}

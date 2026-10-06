package cola_prioridad

const _PANICO_HEAP = "La cola esta vacia"
const _CAPACIDAD_INICIAL = 10
const _DOBLE = 2
const _COTA_SUPERIOR_REDIMENSION = 2
const _COTA_INFERIOR_REDIMENSION = 4
const _MITAD = 2
const _PRIMER_ELEMENTO = 0
const _HEAP_VACIO = 0
const _UN_SOLO_ELEMENTO = 1
const _PRIMERA_CLAVE_MENOR_QUE_SEGUNDA = 0

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
	heapify(arregloH[:len(arreglo)], funcion_cmp)

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
	return heap.cant == _HEAP_VACIO
}

func posicion_padre(indice int) int {
	return (indice - 1) / 2
}

func posicion_hijos(indice int) (int, int) {
	return 2*indice + 1, 2*indice + 2
}

func (heap *heap[T]) Encolar(elemento T) {
	if heap.cant == len(heap.arreglo)/_COTA_SUPERIOR_REDIMENSION {
		heap.redimensionar(len(heap.arreglo) * _DOBLE)
	}
	heap.arreglo[heap.cant] = elemento
	heap.cant++
	if heap.cant == _UN_SOLO_ELEMENTO {
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

	if heap.comparacion(heap.arreglo[indice_padre], heap.arreglo[indice]) < _PRIMERA_CLAVE_MENOR_QUE_SEGUNDA {
		swap(&heap.arreglo[indice_padre], &heap.arreglo[indice])
		heap.upheap(indice_padre)
	}
}

func (heap *heap[T]) VerMax() T {
	if heap.EstaVacia() {
		panic(_PANICO_HEAP)
	}
	return heap.arreglo[_PRIMER_ELEMENTO]
}

func (heap *heap[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic(_PANICO_HEAP)
	}
	dato := heap.arreglo[_PRIMER_ELEMENTO]

	principio := &heap.arreglo[_PRIMER_ELEMENTO]
	fin := &heap.arreglo[heap.cant-1]

	swap(principio, fin)

	heap.cant--

	if heap.cant > _HEAP_VACIO {
		downheap(heap.arreglo, 0, heap.comparacion, heap.cant)
	}
	if heap.cant <= len(heap.arreglo)/_COTA_INFERIOR_REDIMENSION && len(heap.arreglo) > _CAPACIDAD_INICIAL {
		heap.redimensionar(len(heap.arreglo) / _MITAD)
	}

	return dato
}
func encontrar_maximo[T any](indice, cantidad int, arreglo []T, comparacion func(T, T) int) int {
	IndiceIzq, IndiceDer := posicion_hijos(indice)
	maximo := indice
	if IndiceDer < cantidad && comparacion(arreglo[maximo], arreglo[IndiceDer]) < _PRIMERA_CLAVE_MENOR_QUE_SEGUNDA {
		maximo = IndiceDer
	}
	if IndiceIzq < cantidad && comparacion(arreglo[maximo], arreglo[IndiceIzq]) < _PRIMERA_CLAVE_MENOR_QUE_SEGUNDA {
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

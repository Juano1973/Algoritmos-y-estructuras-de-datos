package main

type nodoLista[T any] struct {
	prox *nodoLista[T]
	dato T
}

type ListaEnlazada[T any] struct {
	prim *nodoLista[T]
}

func (l *ListaEnlazada[T]) Ante_k_ultimo(k int) T {
	actual := l.prim

	for i := 0; i < k; i++ {
		actual = actual.prox
	}

	otro := l.prim

	for actual != nil {
		otro = otro.prox
		actual = actual.prox
	}

	return otro.dato
}

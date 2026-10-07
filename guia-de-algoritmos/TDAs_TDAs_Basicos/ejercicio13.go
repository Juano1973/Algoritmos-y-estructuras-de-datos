package main

const _PANICO_LISTA string = "La lista esta vacia"
const _PANICO_ITERADOR string = "El iterador termino de iterar"

type listaEnlazada[T any] struct {
	primero  *nodo[T]
	ultimo   *nodo[T]
	cantidad int
}

type nodo[T any] struct {
	dato      T
	siguiente *nodo[T]
}

type iteradorLista[T any] struct {
	actual   *nodo[T]
	anterior *nodo[T]
	lista    *listaEnlazada[T]
}

func crearNodo[T any](dato T) *nodo[T] {
	return &nodo[T]{dato: dato}
}

func CrearListaEnlazada[T any]() *listaEnlazada[T] {
	//Funcion que crea la lista enlazada, incializa los punteros en nil y la cantidad en 0
	return &listaEnlazada[T]{}
}

func (l *listaEnlazada[T]) InsertarPrimero(dato T) {
	nodo := crearNodo[T](dato)

	if l.EstaVacia() {
		l.ultimo = nodo
	}

	nodo.siguiente = l.primero
	l.primero = nodo
	l.cantidad++
}

func (l *listaEnlazada[T]) InsertarUltimo(dato T) {
	nodo := crearNodo[T](dato)

	if l.EstaVacia() {
		l.primero = nodo
	} else {
		l.ultimo.siguiente = nodo
	}
	l.ultimo = nodo
	l.cantidad++

}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	if l.EstaVacia() {
		panic(_PANICO_LISTA)
	}

	dato := l.primero.dato
	l.primero = l.primero.siguiente
	l.cantidad--
	return dato
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.cantidad == 0
}

func (l *listaEnlazada[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic(_PANICO_LISTA)
	}
	return l.primero.dato
}

func (l *listaEnlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic(_PANICO_LISTA)
	}
	return l.ultimo.dato
}

func (l *listaEnlazada[T]) Largo() int {
	return l.cantidad
}

func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := l.primero
	for actual != nil {
		if !visitar(actual.dato) {
			break
		}
		actual = actual.siguiente
	}

}

//Primitivas del iterador

func (l *listaEnlazada[T]) Iterador() *iteradorLista[T] {
	return &iteradorLista[T]{l.primero, nil, l}
}

func (i *iteradorLista[T]) VerActual() T {
	if !i.HayAlgoMas() {
		panic(_PANICO_ITERADOR)
	}
	return i.actual.dato
}

func (i *iteradorLista[T]) HayAlgoMas() bool {
	return i.actual != nil
}

func (i *iteradorLista[T]) Avanzar() {
	if !i.HayAlgoMas() {
		panic(_PANICO_ITERADOR)
	}
	i.anterior = i.actual
	i.actual = i.actual.siguiente
}

func (i *iteradorLista[T]) Insertar(dato T) {
	nuevo_nodo := crearNodo[T](dato)

	if i.lista.EstaVacia() {
		i.lista.primero = nuevo_nodo
		i.lista.ultimo = nuevo_nodo

	} else if i.actual == i.lista.primero {
		nuevo_nodo.siguiente = i.lista.primero
		i.lista.primero = nuevo_nodo

	} else if i.actual == nil {
		i.lista.ultimo.siguiente = nuevo_nodo
		i.lista.ultimo = nuevo_nodo

	} else {
		i.anterior.siguiente = nuevo_nodo
		nuevo_nodo.siguiente = i.actual
	}
	i.actual = nuevo_nodo
	i.lista.cantidad++
}

func (i *iteradorLista[T]) Borrar() T {

	if !i.HayAlgoMas() {
		panic(_PANICO_ITERADOR)
	}

	dato := i.actual.dato

	if i.actual == i.lista.primero && i.actual == i.lista.ultimo {
		i.lista.primero = nil
		i.lista.ultimo = nil
		i.actual = nil
		i.anterior = nil

	} else if i.actual == i.lista.primero {
		i.lista.primero = i.actual.siguiente
		i.actual = i.lista.primero

	} else if i.actual == i.lista.ultimo {
		i.actual = nil
		i.lista.ultimo = i.anterior
		i.anterior.siguiente = i.actual

	} else {
		i.actual = i.actual.siguiente
		i.anterior.siguiente = i.actual
	}

	i.lista.cantidad--
	return dato
}

func SumaPares(l listaEnlazada[*int]) int {
	sumatoria := 0
	l.Iterar(func(t *int) bool {
		if t != nil && *t%2 == 0 {
			sumatoria += *t
		}

		return true
	})

	return sumatoria
}

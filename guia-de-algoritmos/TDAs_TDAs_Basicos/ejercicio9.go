package main

func (lista *listaEnlazada[T]) Invertir() {

	if lista.EstaVacia() {
		return
	}

	actual := lista.primero
	anterior := lista.primero
	for actual != nil {
		siguiente := actual.siguiente
		if actual == lista.primero {
			actual.siguiente = nil
			actual = siguiente
			continue
		}

		actual = siguiente
		actual.siguiente = anterior
		anterior = actual

	}
	lista.primero, lista.ultimo = lista.ultimo, lista.primero
}

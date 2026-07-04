package cola

const MENSAJE_PANICO string = "La cola esta vacia"

type colaDinamica[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

type nodoCola[T any] struct {
	dato      T
	siguiente *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return new(colaDinamica[T])
}

func crearNodo[T any]() *nodoCola[T] {
	return new(nodoCola[T])
}

func (c *colaDinamica[T]) EstaVacia() bool {
	return c.primero == nil
}

func (c *colaDinamica[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic(MENSAJE_PANICO)
	}
	return c.primero.dato
}

func (c *colaDinamica[T]) Encolar(dato T) {
	nuevo_nodo := crearNodo[T]()
	nuevo_nodo.dato = dato

	if c.EstaVacia() {
		c.primero = nuevo_nodo
	} else {
		c.ultimo.siguiente = nuevo_nodo
	}
	c.ultimo = nuevo_nodo
}

func (c *colaDinamica[T]) Desencolar() T {
	if c.EstaVacia() {
		panic(MENSAJE_PANICO)
	}

	dato := c.primero.dato
	c.primero = c.primero.siguiente

	if c.EstaVacia() {
		c.ultimo = nil
	}
	return dato
}

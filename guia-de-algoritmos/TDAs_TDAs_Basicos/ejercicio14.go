package main

import TDApila "entrenamiento/tdas/pila"

type PilaConMaximo[T comparable] struct {
	arreglo  []T
	cantidad int
	maximo   TDApila.Pila[T]
}

func (p *PilaConMaximo[T]) Apilar(elemento T) {
	p.arreglo[p.cantidad] = elemento
	p.cantidad++
}

func (p *PilaConMaximo[T]) Desapilar() T {
	retorno := p.VerTope()
	p.cantidad--
	return retorno
}

func (p *PilaConMaximo[T]) VerTope() T {
	return p.arreglo[p.cantidad-1]
}

func (p *PilaConMaximo[T]) VerMaximo() T {

	return p.arreglo[0]
}

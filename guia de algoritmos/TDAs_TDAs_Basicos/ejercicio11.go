package main

/* Definición del struct pila proporcionado por la cátedra. */

// Agrego algunas constantes que mejoren la legibilidad

const DOBLE int = 2
const PILA_VACIA int = 0
const VALOR_INICIAL int = 100
const CUATRO_PARTES int = 4
const MENSAJE_PANICO string = "La pila esta vacia"

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func (p *pilaDinamica[T]) redimensionar(nuevoTam int) { //Le paso a redimensionar directamente el tamano para que ejecute la redimension sin importarle si sea achicar o agrandar la capacidad del arreglo
	nuevosDatos := make([]T, nuevoTam)
	copy(nuevosDatos, p.datos)
	p.datos = nuevosDatos
}

func CrearPilaDinamica[T any]() pilaDinamica[T] {
	Nueva_pila := new(pilaDinamica[T])
	Nueva_pila.datos = make([]T, VALOR_INICIAL)
	Nueva_pila.cantidad = PILA_VACIA
	return *Nueva_pila
}

func (p *pilaDinamica[T]) Apilar(dato T) {
	var largo int = len(p.datos)

	if p.cantidad == largo { //Si la cantidad llega al tope del arreglo actual, redimensiono
		p.redimensionar(largo * DOBLE)
	}

	p.datos[p.cantidad] = dato
	p.cantidad += 1
}

func (p *pilaDinamica[T]) Desapilar() T {
	var largo int = len(p.datos)
	if p.EstaVacia() {
		panic(MENSAJE_PANICO)
	}

	dato := p.VerTope()
	p.cantidad -= 1

	//pongo un piso en la cantidad, ya que una vez que entra en 0, no sale
	if p.cantidad > PILA_VACIA && p.cantidad == largo/CUATRO_PARTES && largo/DOBLE >= VALOR_INICIAL { //pongo cuatro partes por que siento que es la manera mas intuitiva que se me ocurre de plasmar ese cuatro en una constante
		p.redimensionar(largo / DOBLE)
	}

	return dato
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == PILA_VACIA
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic(MENSAJE_PANICO)
	}
	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[int]) Ordenar() {
	pila_aux := CrearPilaDinamica[int]()
	if p.EstaVacia() {
		return
	}
	actual := p.Desapilar()

	for !p.EstaVacia() {
		a := p.Desapilar()
		if p.EstaVacia() {
			pila_aux.Apilar(a)
			break
		}

	}

}

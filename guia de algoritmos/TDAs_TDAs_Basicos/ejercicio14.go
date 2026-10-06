package main

/* Definición del struct pila proporcionado por la cátedra. */

// Agrego algunas constantes que mejoren la legibilidad

type pilaConMaximo struct {
	datos    []int
	cantidad int
	maximo   maximo
}

type maximo struct {
	valor    int
	anterior *maximo
}

func (p *pilaConMaximo) redimensionar(nuevoTam int) { //Le paso a redimensionar directamente el tamano para que ejecute la redimension sin importarle si sea achicar o agrandar la capacidad del arreglo
	nuevosDatos := make([]int, nuevoTam)
	copy(nuevosDatos, p.datos)
	p.datos = nuevosDatos
}

func CrearPilaConMaximo() *pilaConMaximo {
	Nueva_pila := new(pilaConMaximo)
	Nueva_pila.datos = make([]int, VALOR_INICIAL)
	Nueva_pila.cantidad = PILA_VACIA
	return Nueva_pila
}

func (p *pilaConMaximo) Apilar(dato int) {
	var largo int = len(p.datos)

	if p.cantidad == largo { //Si la cantidad llega al tope del arreglo actual, redimensiono
		p.redimensionar(largo * DOBLE)
	}
	if p.EstaVacia() {
		max := new(maximo)
		max.valor = dato
		p.maximo = *max
	} else {
		if dato > p.maximo.valor {
			nuevo_max := new(maximo)
			nuevo_max.valor = dato
			nuevo_max.anterior = &p.maximo
			p.maximo = *nuevo_max

		}
	}
	p.datos[p.cantidad] = dato
	p.cantidad += 1
}

func (p *pilaConMaximo) Desapilar() int {
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

func (p *pilaConMaximo) EstaVacia() bool {
	return p.cantidad == PILA_VACIA
}

func (p *pilaConMaximo) VerTope() int {
	if p.EstaVacia() {
		panic(MENSAJE_PANICO)
	}
	return p.datos[p.cantidad-1]
}

func (p *pilaConMaximo) VerMaximo() int {
	if p.EstaVacia() {
		panic(MENSAJE_PANICO)
	}
	return p.maximo.valor
}

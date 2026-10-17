package main

type Arbol struct {
	izq *Arbol
	der *Arbol
}

func (arbol *Arbol) Invertir() {

	if arbol.izq == nil && arbol.der == nil {
		return
	}

	if arbol.izq != nil {
		arbol.izq.Invertir()
	}

	if arbol.der != nil {
		arbol.der.Invertir()
	}

	arbol.izq, arbol.der = arbol.der, arbol.izq
}

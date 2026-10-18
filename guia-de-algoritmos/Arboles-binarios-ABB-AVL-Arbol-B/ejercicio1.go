package main

func (a *nodoAbb[K, V]) altura() int {

	if a.izq == nil && a.der == nil {
		return 1
	}

	var izq int = 0
	var der int = 0

	if a.der != nil {
		izq = a.izq.altura()
	}

	if a.der != nil {
		der = a.der.altura()
	}

	return max(izq, der) + 1

}

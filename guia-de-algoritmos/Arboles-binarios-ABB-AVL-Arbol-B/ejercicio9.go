package main

func (arbol *Arbol) Quiebres() int {

	if arbol.izq == nil && arbol.der == nil {
		return 0
	}

	hay_izq := false
	hay_der := false

	var izq int = 0
	var der int = 0

	if arbol.izq != nil {
		hay_izq = true
		izq = arbol.izq.Quiebres()
	}

	if arbol.der != nil {
		hay_der = true
		der = arbol.der.Quiebres()
	}

	if !hay_der || !hay_izq {
		return 1 + izq + der
	}

	return izq + der

}

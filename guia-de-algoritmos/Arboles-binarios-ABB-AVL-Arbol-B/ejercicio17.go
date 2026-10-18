package main

import (
	TDAlista "entrenamiento/tdas/lista"
)

func (ab *abb[K, V]) cotas_piola(techo, piso K, nivel int) TDAlista.Lista[V] {

	list := TDAlista.CrearListaEnlazada[V]()
	ab.raiz._cotas_piola(techo, piso, nivel, 1, list)
	return list
}

func (nodo *nodoAbb[K, V]) _cotas_piola(techo, piso K, nivel, nivel_actual int, l TDAlista.Lista[V]) {

	if nivel_actual > nivel {
		return
	}

	if nodo.clave > piso && nodo.izq != nil {
		nodo.izq._cotas_piola(techo, piso, nivel, nivel_actual+1, l)
	}

	if nodo.clave >= piso && nodo.clave <= techo {
		l.InsertarUltimo(nodo.dato)
	}

	if nodo.clave < techo && nodo.der != nil {
		nodo.der._cotas_piola(techo, piso, nivel, nivel_actual+1, l)
	}
}

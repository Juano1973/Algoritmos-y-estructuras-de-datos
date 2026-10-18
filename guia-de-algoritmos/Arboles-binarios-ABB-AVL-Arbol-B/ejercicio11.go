package main

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	comparar func(K, K) int
	cantidad int
}

type nodoAbb[K comparable, V any] struct {
	izq   *nodoAbb[K, V]
	der   *nodoAbb[K, V]
	clave K
	dato  V
}

func (ab *abb[K, V]) claves_ordenadas() []K {
	actual := ab.raiz
	arreglo := make([]K, 0)
	actual._claves_ordenadas(arreglo)
	return arreglo
}

func (nodo *nodoAbb[K, V]) _claves_ordenadas(arreglo []K) []K {

	arreglo = append(arreglo, nodo.clave)

	if nodo.izq != nil {
		nodo.izq._claves_ordenadas(arreglo)
	}

	if nodo.der != nil {
		nodo.der._claves_ordenadas(arreglo)
	}

	return arreglo
}

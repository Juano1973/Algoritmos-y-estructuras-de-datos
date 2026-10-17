package main

type abb[K comparable, V any] struct {
	hijoizq *abb[K, V]
	hijoder *abb[K, V]
	clave   K
	dato    int
}

func (a *abb[K, V]) altura() int {

	if a.hijoder == nil && a.hijoizq == nil {
		return 1
	}

	var izq int = 0
	var der int = 0

	if a.hijoizq != nil {
		izq = a.hijoizq.altura()
	}

	if a.hijoder != nil {
		der = a.hijoder.altura()
	}

	return max(izq, der) + 1

}

package diccionario

import (
	TDApila "tdas/pila"
)

const _PANICO_ABB = "La clave no pertenece al diccionario"

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

type iterAbb[K comparable, V any] struct {
	actual   *nodoAbb[K, V]
	pila     TDApila.Pila[*nodoAbb[K, V]]
	desde    *K
	hasta    *K
	comparar func(K, K) int
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{comparar: funcion_cmp}
}

func crearNodo[K comparable, V any](clave K, valor V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{clave: clave, dato: valor}
}

func (abb *abb[K, V]) Guardar(clave K, valor V) {

	nuevoNodo := crearNodo(clave, valor)

	if abb.cantidad != 0 {
		lugar := abb.raiz.encontrarLugar(clave, abb.comparar)

		contenido := *lugar
		if contenido != nil {

			if abb.comparar(contenido.clave, clave) != 0 {
				abb.cantidad++
			}

			contenido.dato = valor

			return
		}

		*lugar = nuevoNodo

	} else {
		abb.raiz = nuevoNodo
	}

	abb.cantidad++

}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	if abb.Cantidad() == 0 {
		return false
	}
	nodo := *abb.raiz.encontrarLugar(clave, abb.comparar)
	if nodo == nil {
		return false
	}

	return abb.comparar(nodo.clave, clave) == 0
}

func (abb *abb[K, V]) Obtener(clave K) V {
	if abb.Cantidad() == 0 {
		panic(_PANICO_ABB)
	}
	nodo := *abb.raiz.encontrarLugar(clave, abb.comparar)
	if nodo == nil || abb.comparar(nodo.clave, clave) != 0 {
		panic(_PANICO_ABB)
	}
	return nodo.dato
}
func (abb *abb[K, V]) Borrar(clave K) V {
	if abb.Cantidad() == 0 {
		panic(_PANICO_ABB)
	}

	puntero := abb.raiz.encontrarLugar(clave, abb.comparar)
	nodo := *puntero
	if nodo == nil || abb.comparar(nodo.clave, clave) != 0 {
		panic(_PANICO_ABB)
	}

	dato := nodo.dato
	var remplazo **nodoAbb[K, V]

	if nodo.izq != nil || nodo.der != nil {
		if nodo.der != nil {
			remplazo, _ = nodo.der.masIzquierdoDerecha(&nodo)
		} else {
			remplazo, _ = nodo.izq.masDerechoIzquierda(&nodo)
		}

		contenido_remplazo := *remplazo

		lugar_remplazo := abb.raiz.encontrarLugar(contenido_remplazo.clave, abb.comparar)

		nodo.clave = contenido_remplazo.clave
		nodo.dato = contenido_remplazo.dato

		puntero = lugar_remplazo
		nodo = *puntero
	}

	var hijo *nodoAbb[K, V]
	if nodo.izq != nil {
		hijo = nodo.izq
	} else {
		hijo = nodo.der
	}

	*puntero = hijo

	if nodo == abb.raiz {
		abb.raiz = hijo
	}

	abb.cantidad--

	return dato
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Iterar(f func(clave K, valor V) bool) {

	if abb.raiz == nil {
		return
	}

	corte := false
	abb.raiz._iterar(f, &corte)
}

func (nodo *nodoAbb[K, V]) _iterar(f func(clave K, valor V) bool, corte *bool) {
	if !*corte {
		if nodo.izq != nil {
			nodo.izq._iterar(f, corte)
		}
	}

	if !*corte {
		if !f(nodo.clave, nodo.dato) {
			*corte = true
			return
		}
	}

	if !*corte {
		if nodo.der != nil {
			nodo.der._iterar(f, corte)
		}
	}

}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

func (nodo *nodoAbb[K, V]) hijos_nulos_menores_condiciones(cmp func(K, K) int, desde *K) bool {
	//esta funcion me dice si estoy en la posicion minima donde se cumple desde
	if nodo.der == nil && nodo.izq == nil {
		return true
	}

	if (nodo.der != nil && cmp(nodo.der.clave, *desde) > 0) || (nodo.izq != nil && cmp(nodo.izq.clave, *desde) > 0) {
		return false
	}

	return true
}

func (iter *iterAbb[K, V]) cumpleRangoo(nodo *nodoAbb[K, V]) bool {
	cota_inferior := iter.desde == nil || iter.comparar(nodo.clave, *iter.desde) >= 0
	cota_superior := iter.hasta == nil || iter.comparar(nodo.clave, *iter.hasta) <= 0
	return cota_inferior && cota_superior
}

func (iter *iterAbb[K, V]) apilarDesde(nodo *nodoAbb[K, V]) {
	for nodo != nil {
		if iter.desde == nil || (iter.hasta == nil && iter.comparar(*iter.desde, nodo.clave) <= 0) || (iter.comparar(*iter.desde, nodo.clave) <= 0 && iter.comparar(*iter.hasta, nodo.clave) >= 0) {
			iter.pila.Apilar(nodo)
			nodo = nodo.izq
		} else if iter.desde != nil && iter.comparar(nodo.clave, *iter.desde) < 0 {
			nodo = nodo.der
		} else {
			nodo = nodo.izq
		}
	}
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	pila := TDApila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iterador := &iterAbb[K, V]{pila: pila, desde: desde, hasta: hasta, comparar: abb.comparar}

	if abb.raiz == nil {
		return iterador
	}

	iterador.apilarDesde(abb.raiz)

	return iterador
}

func (iterador *iterAbb[K, V]) Avanzar() {
	if !iterador.HayAlgoMas() {
		panic(_PANICO_ITERADOR)
	}

	siguiente := iterador.pila.Desapilar()

	if siguiente.der != nil {

		if iterador.hasta != nil {

			siguiente.der.apilarNodosRango(iterador.pila, iterador.comparar, iterador.desde, iterador.hasta)

		} else {

			siguiente.der.apilarNodos(iterador.pila)

		}

	}

	iterador.actual = siguiente
}

func (iterador *iterAbb[K, V]) VerActual() (K, V) {
	if !iterador.HayAlgoMas() {
		panic(_PANICO_ITERADOR)
	}
	return iterador.pila.VerTope().clave, iterador.pila.VerTope().dato
}

func (iterador *iterAbb[K, V]) HayAlgoMas() bool {
	return !iterador.pila.EstaVacia()
}

func (nodo *nodoAbb[K, V]) cumpleRango(desde, hasta *K, comparar func(K, K) int) bool {
	cota_inferior := desde == nil || comparar(nodo.clave, *desde) >= 0
	cota_superior := hasta == nil || comparar(nodo.clave, *hasta) <= 0
	return cota_inferior && cota_superior
}
func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {

	pila := TDApila.CrearPilaDinamica[*nodoAbb[K, V]]()
	copia_desde := desde

	nodo := abb.raiz

	if abb.cantidad == 0 || ((desde != nil && hasta != nil) && abb.comparar(*desde, *hasta) > 0) {

		return
	}

	if desde != nil {
		for !nodo.cumpleRango(desde, hasta, abb.comparar) {
			if abb.comparar(nodo.clave, *desde) <= 0 && nodo.der != nil {
				nodo = nodo.der
			} else if nodo.izq != nil {
				nodo = nodo.izq
			} else {
				break
			}
		}

	} else {
		desde = &nodo.clave
		for nodo.izq != nil {
			nodo = nodo.izq
			desde = &nodo.clave
		}

	}
	if nodo == nil || !nodo.cumpleRango(desde, hasta, abb.comparar) {
		return
	}
	if copia_desde == nil && hasta == nil {
		corte := false
		abb.raiz._iterar(visitar, &corte)
		return
	} else {
		nodo._iterar_rango(nodo, visitar, pila, desde, hasta, abb.comparar)
	}

}

func (nodo *nodoAbb[K, V]) _iterar_rango(ultimo *nodoAbb[K, V], f func(clave K, dato V) bool, pila TDApila.Pila[*nodoAbb[K, V]], desde, hasta *K, cmp func(K, K) int) {

	if nodo.izq != nil && (desde == nil || cmp(nodo.clave, *desde) >= 0) && nodo.izq != ultimo {
		nodo.izq._iterar_rango(nodo, f, pila, desde, hasta, cmp)
	}

	if desde != nil && hasta != nil {
		if (cmp(nodo.clave, *desde) >= 0 && cmp(nodo.clave, *hasta) <= 0) && !f(nodo.clave, nodo.dato) {
			return
		}
	} else if desde != nil && hasta == nil {
		if !(cmp(nodo.clave, *desde) >= 0) || !f(nodo.clave, nodo.dato) {
			return
		}
	} else if hasta != nil && desde == nil {
		if (cmp(nodo.clave, *hasta) <= 0) || (cmp(nodo.clave, *hasta) == 0 && !f(nodo.clave, nodo.dato)) {
			return
		}
	}

	if nodo.der != nil && (hasta == nil || cmp(nodo.clave, *hasta) <= 0) {
		nodo.der._iterar_rango(nodo, f, pila, desde, hasta, cmp)
	}

}

func (nodo *nodoAbb[K, V]) _iterar_rangoo(f func(clave K, dato V) bool, pila TDApila.Pila[*nodoAbb[K, V]], desde, hasta *K, cmp func(K, K) int) {

	if nodo.izq != nil && (desde == nil || cmp(nodo.izq.clave, *desde) >= 0) {
		nodo.izq._iterar_rangoo(f, pila, desde, hasta, cmp)
	}

	if desde != nil && hasta != nil {
		if ((cmp(nodo.clave, *desde) >= 0 && cmp(nodo.clave, *hasta) <= 0) || (cmp(nodo.clave, *desde) == 0 || cmp(nodo.clave, *hasta) == 0)) && !f(nodo.clave, nodo.dato) {
			return
		}
	} else if desde != nil {
		if (cmp(nodo.clave, *desde) >= 0) || (cmp(nodo.clave, *desde) == 0) || !f(nodo.clave, nodo.dato) {
			return
		}
	} else if hasta != nil {
		if (cmp(nodo.clave, *hasta) <= 0) || (cmp(nodo.clave, *hasta) == 0 && !f(nodo.clave, nodo.dato)) {
			return
		}
	}

	if nodo.der != nil && (hasta == nil || cmp(nodo.der.clave, *hasta) <= 0) {
		nodo.der._iterar_rangoo(f, pila, desde, hasta, cmp)
	}
}

func (nodo *nodoAbb[K, V]) apilarNodos(pila TDApila.Pila[*nodoAbb[K, V]]) {

	pila.Apilar(nodo)

	if nodo != nil && nodo.izq != nil {
		nodo.izq.apilarNodos(pila)
	}
}

func (nodo *nodoAbb[K, V]) apilarNodosRango(pila TDApila.Pila[*nodoAbb[K, V]], cmp func(K, K) int, desde, hasta *K) {

	if nodo.cumpleRango(desde, hasta, cmp) {
		pila.Apilar(nodo)
	}
	if nodo.izq != nil {
		nodo.izq.apilarNodosRango(pila, cmp, desde, hasta)
	}

}

func (nodo *nodoAbb[K, V]) encontrarLugar(clave K, comparar func(K, K) int) **nodoAbb[K, V] {

	actual := &nodo

	for *actual != nil {
		comparacion := comparar((*actual).clave, clave)

		if comparacion == 0 {
			return actual

		} else if comparacion > 0 {
			actual = &(*actual).izq

		} else {
			actual = &(*actual).der
		}
	}
	return actual
}

//func (nodo *nodoAbb[K, V]) encontrarLugarr(clave K, cmp func(K, K) int, cantidad int) **nodoAbb[K, V] {
//
//
//	comparacion := cmp(nodo.clave, clave)
//	if comparacion == 0 {
//		return &nodo
//	}
//
//	if nodo.izq != nil && cmp(nodo.izq.clave, clave) == 0 {
//		return &nodo.izq
//	}
//
//	if nodo.der != nil && cmp(nodo.der.clave, clave) == 0 {
//		return &nodo.der
//	}
//
//	if comparacion > 0 {
//		if nodo.izq != nil {
//			return nodo.izq.encontrarLugarr(clave, cmp, cantidad)
//		}
//		return &nodo.izq
//
//	}
//
//	if nodo.der != nil {
//		return nodo.der.encontrarLugarr(clave, cmp, cantidad)
//	}
//	return &nodo.der
//
//}

func (nodo *nodoAbb[K, V]) masIzquierdoDerecha(padre **nodoAbb[K, V]) (**nodoAbb[K, V], **nodoAbb[K, V]) {

	if nodo.izq == nil {
		return &nodo, padre
	}

	return nodo.izq.masIzquierdoDerecha(&nodo)
}

func (nodo *nodoAbb[K, V]) masDerechoIzquierda(padre **nodoAbb[K, V]) (**nodoAbb[K, V], **nodoAbb[K, V]) {

	if nodo.der == nil {
		return &nodo, padre
	}

	return nodo.der.masDerechoIzquierda(&nodo)
}

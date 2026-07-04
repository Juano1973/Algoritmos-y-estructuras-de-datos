package diccionario

import (
	"fmt"
	TDApila "tdas/pila"
)

const _EL_PRIMERO_ES_MAS_GRANDE = 0
const _EL_SEGUNDO_ES_MAS_GRANDE = 0
const _SON_IGUALES = 0
const _ABB_VACIO = 0
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

type iteradorAbb[K comparable, V any] struct {
	pila   TDApila.Pila[nodoAbb[K, V]]
	actual *nodoAbb[K, V]
	padre  *nodoAbb[K, V]
}

type iteradorAbbrango[K comparable, V any] struct {
	actual   *nodoAbb[K, V]
	pila     TDApila.Pila[nodoAbb[K, V]]
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
		lugar := abb.raiz.encontrarLugar(clave, abb.comparar, abb.cantidad)

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
	if abb.Cantidad() == _ABB_VACIO {
		return false
	}
	nodo := *abb.raiz.encontrarLugar(clave, abb.comparar, abb.cantidad)
	if nodo == nil {
		return false
	}

	return abb.comparar(nodo.clave, clave) == _SON_IGUALES
}

func (abb *abb[K, V]) Obtener(clave K) V {
	if abb.Cantidad() == 0 {
		panic(_PANICO_ABB)
	}
	nodo := *abb.raiz.encontrarLugar(clave, abb.comparar, abb.cantidad)
	if nodo == nil || abb.comparar(nodo.clave, clave) != 0 {
		panic(_PANICO_ABB)
	}
	return nodo.dato
}
func (abb *abb[K, V]) Borrar(clave K) V {
	if abb.Cantidad() == 0 {
		panic(_PANICO_ABB)
	}

	puntero := abb.raiz.encontrarLugar(clave, abb.comparar, abb.cantidad)
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

		lugar_remplazo := abb.raiz.encontrarLugar(contenido_remplazo.clave, abb.comparar, abb.cantidad)

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

func (abb *abb[K, V]) Iterar(f func(clave K, valor V) bool) { //✅

	pila := TDApila.CrearPilaDinamica[*nodoAbb[K, V]]()

	if abb.raiz == nil {
		return
	}

	abb.raiz._iterar(f, pila)
}

func (nodo *nodoAbb[K, V]) _iterar(f func(clave K, valor V) bool, pila TDApila.Pila[*nodoAbb[K, V]]) {
	if nodo.izq != nil {
		nodo.izq._iterar(f, pila)
	}

	if !f(nodo.clave, nodo.dato) {
		return
	}

	if nodo.der != nil {
		nodo.der._iterar(f, pila)
	}
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	iterador := &iteradorAbb[K, V]{}
	pila := TDApila.CrearPilaDinamica[nodoAbb[K, V]]()
	iterador.pila = pila

	if abb.cantidad == 0 {
		return iterador
	}

	nodo := abb.raiz
	for nodo.izq != nil {

		iterador.pila.Apilar(*nodo)
		nodo = nodo.izq
	}
	iterador.pila.Apilar(*nodo)
	return iterador
}

func (iterador *iteradorAbb[K, V]) Avanzar() {
	if !iterador.HayAlgoMas() {
		panic(_PANICO_ITERADOR)
	}

	siguiente := iterador.pila.Desapilar()

	if siguiente.der != nil {
		arreglo := make([]*nodoAbb[K, V], 0)
		arreglo = siguiente.der.apilarNodos(arreglo)
		for _, e := range arreglo {
			iterador.pila.Apilar(*e)
		}
	}
	iterador.actual = &siguiente
}

func (iterador *iteradorAbb[K, V]) VerActual() (K, V) {
	if !iterador.HayAlgoMas() {
		panic(_PANICO_ITERADOR)
	}
	return iterador.pila.VerTope().clave, iterador.pila.VerTope().dato
}

func (iterador *iteradorAbb[K, V]) HayAlgoMas() bool {
	return !(iterador.pila.EstaVacia())

}

func (nodo *nodoAbb[K, V]) hijos_nulos_menores_condiciones(cmp func(K, K) int, desde *K) bool {
	//esta funcion me dice si estoy en la posicion minima donde se cumple desde

	//si los dos hijos son nil, estoy en una posicion irreductible
	if nodo.der == nil && nodo.izq == nil {
		return true
	}

	//si tiene hijos y estos cumplen las condiciones, puedo seguir moviendome
	if (nodo.der != nil && cmp(nodo.der.clave, *desde) < 0) || (nodo.izq != nil && cmp(nodo.izq.clave, *desde) > 0) {
		return false
	}

	return true
}

func (nodo *nodoAbb[K, V]) buscar_primero(desde *K, comparar func(K, K) int) *nodoAbb[K, V] {

	for !nodo.hijos_nulos_menores_condiciones(comparar, desde) {
		if comparar(nodo.clave, *desde) > 0 && nodo.izq != nil {
			nodo = nodo.izq
		} else if nodo.der != nil {
			nodo = nodo.der
		} else {
			break
		}
	}

	//for comparar(nodo.clave, *desde) >= 0 {
	//
	//	if nodo.izq != nil && comparar(nodo.izq.clave, *desde) > 0 {
	//		nodo = nodo.izq
	//	} else if nodo.der != nil {
	//		nodo = nodo.der
	//	} else {
	//		return nodo
	//	}
	//}
	return nodo
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	pila := TDApila.CrearPilaDinamica[nodoAbb[K, V]]()
	iterador := &iteradorAbbrango[K, V]{pila: pila, desde: desde, hasta: hasta, comparar: abb.comparar}

	if abb.cantidad == _ABB_VACIO {
		return iterador
	}

	if desde != nil {
		nodo := abb.raiz.buscar_primero(desde, abb.comparar)

		if abb.comparar(*desde, nodo.clave) >= 0 {
			iterador.pila.Apilar(*nodo)
		}
	} else {
		nodo := abb.raiz
		for nodo.izq != nil {

			iterador.pila.Apilar(*nodo)
			nodo = nodo.izq
		}
		iterador.pila.Apilar(*nodo)
		contenido := *nodo
		desde = &contenido.clave
	}

	return iterador
}

func (iterador *iteradorAbbrango[K, V]) Avanzar() {
	if !iterador.HayAlgoMas() {
		panic(_PANICO_ITERADOR)
	}

	siguiente := iterador.pila.Desapilar()

	if siguiente.der != nil {

		if iterador.comparar(siguiente.der.clave, *iterador.hasta) >= 0 {
			fmt.Println(siguiente.der.clave)
			arreglo := make([]*nodoAbb[K, V], 0)
			arreglo = siguiente.der.apilarNodosRango(iterador.comparar, arreglo, iterador.desde, iterador.hasta)
			for _, e := range arreglo {
				iterador.pila.Apilar(*e)
			}

		}
	}
	fmt.Println(iterador.pila.VerTope().clave, "tope")

	iterador.actual = &siguiente
}

func (iterador *iteradorAbbrango[K, V]) VerActual() (K, V) {
	if !iterador.HayAlgoMas() {
		panic(_PANICO_ITERADOR)
	}
	return iterador.pila.VerTope().clave, iterador.pila.VerTope().dato
}

func (iterador *iteradorAbbrango[K, V]) HayAlgoMas() bool {
	pila_no_vacia := !iterador.pila.EstaVacia()
	//if pila_no_vacia {
	//	return !(iterador.comparar(iterador.pila.VerTope().clave, *iterador.hasta) < 0)
	//}
	return pila_no_vacia
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {

	pila := TDApila.CrearPilaDinamica[*nodoAbb[K, V]]()

	nodo := abb.raiz
	for nodo.izq != nil {
		if abb.comparar(nodo.clave, *desde) < 0 {
			continue
		}

		pila.Apilar(nodo)
		nodo = nodo.izq
	}

	nodo._iterar_rango(visitar, pila, desde, hasta, abb.comparar)
}

func (nodo *nodoAbb[K, V]) _iterar_rango(f func(clave K, dato V) bool, pila TDApila.Pila[*nodoAbb[K, V]], desde, hasta *K, cmp func(K, K) int) {

	if nodo.izq != nil && cmp(nodo.izq.clave, *desde) <= 0 {
		nodo.izq._iterar_rango(f, pila, desde, hasta, cmp)
	}

	if ((cmp(nodo.clave, *desde) <= 0 && cmp(nodo.clave, *hasta) >= 0) || (cmp(nodo.clave, *desde) == 0 || cmp(nodo.clave, *hasta) == 0)) && !f(nodo.clave, nodo.dato) {
		return
	}

	if nodo.der != nil && cmp(nodo.der.clave, *hasta) >= 0 {
		nodo.der._iterar_rango(f, pila, desde, hasta, cmp)
	}
}

func (nodo *nodoAbb[K, V]) apilarNodos(arr []*nodoAbb[K, V]) []*nodoAbb[K, V] {

	arr = append(arr, nodo)
	if nodo.izq != nil {
		arr = nodo.izq.apilarNodos(arr)
	}

	return arr
}

func (nodo *nodoAbb[K, V]) apilarNodosRango(cmp func(K, K) int, arr []*nodoAbb[K, V], desde, hasta *K) []*nodoAbb[K, V] {

	if cmp(nodo.clave, *desde) <= 0 && cmp(nodo.clave, *hasta) >= 0 {
		arr = append(arr, nodo)
	}
	if nodo.izq != nil && cmp(nodo.izq.clave, *desde) <= 0 && cmp(nodo.izq.clave, *hasta) >= 0 {
		arr = nodo.izq.apilarNodos(arr)
	}

	return arr
}

func (nodo *nodoAbb[K, V]) encontrarLugar(clave K, cmp func(K, K) int, cantidad int) **nodoAbb[K, V] {

	comparacion := cmp(nodo.clave, clave)

	if comparacion == _SON_IGUALES {
		return &nodo
	}

	if nodo.izq != nil && cmp(nodo.izq.clave, clave) == _SON_IGUALES {
		return &nodo.izq
	}

	if nodo.der != nil && cmp(nodo.der.clave, clave) == _SON_IGUALES {
		return &nodo.der
	}

	if comparacion < _EL_PRIMERO_ES_MAS_GRANDE {
		if nodo.izq != nil {
			return nodo.izq.encontrarLugar(clave, cmp, cantidad)
		}

		return &nodo.izq
	}

	if nodo.der != nil {
		return nodo.der.encontrarLugar(clave, cmp, cantidad)
	}
	return &nodo.der
}

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

package diccionario

import (
	"fmt"
	TDApila "tdas/pila"
)

const _EL_PRIMERO_ES_MAS_GRANDE = 0
const _EL_SEGUNDO_ES_MAS_GRANDE = 0
const _SON_IGUALES = 0

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

func nuevoAbb[K comparable, V any](funcion_cmp func(K, K) int) *abb[K, V] {
	return &abb[K, V]{comparar: funcion_cmp}
}

func (abb *abb[K, V]) Guardar(clave K, valor V) { //✅

	nuevoNodo := crearNodo(clave, valor)
	sumar := 1

	if abb.cantidad != 0 {
		lugar := abb.raiz.encontrarLugar(clave, abb.comparar, abb.cantidad)

		contenido := *lugar
		if contenido != nil {

			sumar = 0
			contenido.dato = valor

			return
		}

		*lugar = nuevoNodo

	} else {
		abb.raiz = nuevoNodo
	}
	abb.cantidad += sumar

}

func (abb *abb[K, V]) Pertenece(clave K) bool { //✅
	if abb.Cantidad() == 0 {
		return false
	}
	nodo := *abb.raiz.encontrarLugar(clave, abb.comparar, abb.cantidad)
	if nodo == nil {
		return false
	}
	return nodo.clave == clave
}

func (abb *abb[K, V]) Obtener(clave K) V { //✅
	if abb.Cantidad() == 0 {
		panic(_PANICO_HASH)
	}
	nodo := *abb.raiz.encontrarLugar(clave, abb.comparar, abb.cantidad)
	if nodo == nil || nodo.clave != clave {
		panic(_PANICO_HASH)
	}
	return nodo.dato
}

func (abb *abb[K, V]) Borrar(clave K) V { //✅

	if abb.Cantidad() == 0 {
		panic(_PANICO_HASH)
	}

	puntero := abb.raiz.encontrarLugar(clave, abb.comparar, abb.cantidad) //Esto seria la "flecha" entre el padre y el hijo a borrar, un puntero a puntero
	nodo := *puntero
	//Este seria el hijo a borrar
	if nodo == nil || nodo.clave != clave {
		panic(_PANICO_HASH)
	}
	var dato V
	var remplazo *nodoAbb[K, V] = nil
	if nodo.der != nil {
		remplazo = nodo.der.masIzquierdoDerecha()
	} else if nodo.izq != nil {
		remplazo = nodo.izq.masDerechoIzquierda()
	}

	dato = nodo.dato

	//remplazo.der = nodo.der
	//remplazo.izq = nodo.izq
	*puntero = remplazo

	if nodo == abb.raiz {
		abb.raiz = remplazo
	}
	remplazo = nil
	abb.cantidad--
	return dato
}

func (abb *abb[K, V]) Cantidad() int { //✅
	return abb.cantidad
}

func (abb *abb[K, V]) Iterar(f func(clave K, valor V) bool) { //✅

	pila := TDApila.CrearPilaDinamica[*nodoAbb[K, V]]()

	if abb.raiz == nil {
		return
	}

	abb.raiz._iterador(f, pila, abb.raiz)
}

func (nodo *nodoAbb[K, V]) _iterador(f func(clave K, valor V) bool, pila TDApila.Pila[*nodoAbb[K, V]], ultimo *nodoAbb[K, V]) {
	if nodo.izq != nil {
		//pila.Apilar(nodo)
		//if nodo.der != nil {
		//	arreglo := make([]*nodoAbb[K, V], 0)
		//	arreglo = nodo.der.apilarNodos(arreglo)
		//	for _, e := range arreglo {
		//		pila.Apilar(e)
		//	}
		//}
		nodo.izq._iterador(f, pila, nodo)
	}

	if !f(nodo.clave, nodo.dato) {
		return
	}

	if nodo.der != nil {
		//pila.Apilar(nodo)
		//arreglo := make([]*nodoAbb[K, V], 0)
		//arreglo = nodo.der.apilarNodos(arreglo)
		//for _, e := range arreglo {
		//	pila.Apilar(e)
		//}
		nodo.der._iterador(f, pila, nodo)
	}
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] { //✅
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

func (iterador *iteradorAbb[K, V]) Avanzar() { //✅
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

func (iterador *iteradorAbb[K, V]) VerActual() (K, V) { //✅
	if !iterador.HayAlgoMas() {
		panic(_PANICO_ITERADOR)
	}
	return iterador.pila.VerTope().clave, iterador.pila.VerTope().dato
}

func (iterador *iteradorAbb[K, V]) HayAlgoMas() bool { //✅
	return !(iterador.pila.EstaVacia())

}

func (nodo *nodoAbb[K, V]) buscar_primero(desde *K, comparar func(K, K) int) *nodoAbb[K, V] {
	for !(comparar(nodo.clave, *desde) < 0) {

		if nodo.izq != nil {
			nodo = nodo.izq
		} else if nodo.der != nil {
			nodo = nodo.der
		} else {
			return nodo
		}
	}
	return nodo
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] { //✅
	pila := TDApila.CrearPilaDinamica[nodoAbb[K, V]]()
	iterador := &iteradorAbbrango[K, V]{pila: pila, desde: desde, hasta: hasta, comparar: abb.comparar}

	if abb.cantidad == 0 {
		return iterador
	}

	nodo := abb.raiz.buscar_primero(desde, abb.comparar)

	//for !(abb.comparar(nodo.clave, *desde) < 0) {
	//	if nodo
	//}
	//
	//
	//if abb.comparar(nodo.clave, *desde) < 0 {
	//
	//	//repetir()
	//	for nodo != nil {
	//		if iterador.comparar(*desde, nodo.clave) < 0 && iterador.comparar(*hasta, nodo.clave) > 0 {
	//			iterador.pila.Apilar(*nodo)
	//		} else {
	//		}
	//		nodo = nodo.izq
	//	}
	//}
	//
	//if abb.comparar(nodo.clave, *desde) > 0 {
	//
	//}
	//
	//for abb.comparar(nodo.clave, *desde) > 0 {
	//	if nodo.izq != nil && abb.comparar(nodo.izq.clave, *desde) > 0 {
	//		iterador.pila.Apilar(*nodo.izq)
	//		nodo = nodo.izq
	//	}
	//}

	fmt.Println(nodo.clave)
	iterador.pila.Apilar(*nodo)
	return iterador
}

func (iterador *iteradorAbbrango[K, V]) Avanzar() { //✅
	if !iterador.HayAlgoMas() {
		panic(_PANICO_ITERADOR)
	}

	siguiente := iterador.pila.Desapilar()

	if siguiente.der != nil && iterador.comparar(siguiente.der.clave, *iterador.hasta) > 0 {
		arreglo := make([]*nodoAbb[K, V], 0)
		arreglo = siguiente.der.apilarNodos(arreglo)
		for _, e := range arreglo {
			iterador.pila.Apilar(*e)
		}
	}
	iterador.actual = &siguiente

}

func (iterador *iteradorAbbrango[K, V]) VerActual() (K, V) { //✅
	if !iterador.HayAlgoMas() {
		panic(_PANICO_ITERADOR)
	}
	return iterador.pila.VerTope().clave, iterador.pila.VerTope().dato
}

func (iterador *iteradorAbbrango[K, V]) HayAlgoMas() bool { //✅
	return !(iterador.pila.EstaVacia())
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {

	pila := TDApila.CrearPilaDinamica[*nodoAbb[K, V]]()

	nodo := abb.raiz
	for nodo.izq != nil {
		if abb.comparar(nodo.clave, *desde) < 0 {
			continue
		}
		//fmt.Println(nodo.clave)

		pila.Apilar(nodo)
		nodo = nodo.izq
	}

	nodo._iterar_rango(visitar, pila, desde, hasta, abb.comparar)
}

func (nodo *nodoAbb[K, V]) _iterar_rango(f func(clave K, dato V) bool, pila TDApila.Pila[*nodoAbb[K, V]], desde, hasta *K, cmp func(K, K) int) {

	if nodo.izq != nil && cmp(nodo.izq.clave, *desde) > 0 {
		nodo.izq._iterar_rango(f, pila, desde, hasta, cmp)
	}

	if ((cmp(nodo.clave, *desde) < 0 && cmp(nodo.clave, *hasta) > 0) || (cmp(nodo.clave, *desde) == 0 || cmp(nodo.clave, *hasta) == 0)) && !f(nodo.clave, nodo.dato) {
		return
	}

	if nodo.der != nil && (cmp(nodo.der.clave, *hasta) > 0 || cmp(nodo.der.clave, *hasta) == 0) {
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

func (nodo *nodoAbb[K, V]) apilarNodosRangos(desde, hasta *K, comparar func(K, K) int, arr []*nodoAbb[K, V]) []*nodoAbb[K, V] {

	if comparar(nodo.clave, *hasta) < 0 {
		arr = append(arr, nodo)
	}
	if nodo.izq != nil && comparar(nodo.izq.clave, *desde) > 0 {
		arr = nodo.izq.apilarNodos(arr)
	}

	return arr
}

func (nodo *nodoAbb[K, V]) encontrarLugar(clave K, cmp func(K, K) int, cantidad int) **nodoAbb[K, V] {

	//Esta funcion tiene que encontrar el lugar que ocupa o ocuparia cualquier nodo

	var comparacion int = cmp(nodo.clave, clave)

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

func (nodo *nodoAbb[K, V]) masIzquierdoDerecha() *nodoAbb[K, V] {

	if nodo.izq == nil {
		return nodo
	}

	if nodo.der != nil {
		return nodo.der.masIzquierdoDerecha()
	}

	return nodo
}

func (nodo *nodoAbb[K, V]) masDerechoIzquierda() *nodoAbb[K, V] {

	if nodo.der == nil {
		return nodo
	}

	if nodo.izq != nil {
		return nodo.izq.masDerechoIzquierda()
	}

	return nodo
}

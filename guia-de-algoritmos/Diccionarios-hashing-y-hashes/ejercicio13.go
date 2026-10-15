package main

import (
	Tdalista "entrenamiento/tdas/lista"
	"fmt"
)

type hashElem[K comparable] struct {
	clave K
	dato  int
}

type hashAbierto[K comparable] struct {
	tabla    []Tdalista.Lista[hashElem[K]]
	cantidad int
	tam      int
	sumado   int
}

type iterDiccionario[K comparable] struct {
	indice    int
	hash      *hashAbierto[K]
	iterLista Tdalista.IteradorLista[hashElem[K]]
}

const _PANICO_HASH = "La clave no pertenece al diccionario"
const _PANICO_ITERADOR = "El iterador termino de iterar"
const _CAPACIDAD_INICIAL = 10
const _HASH_VACIO = 0
const _FACTOR_DE_CARGA_MAXIMO = 3
const _FACTOR_DE_CARGA_MINIMO = 0.25
const _DOBLE = 2

func CrearHash[K comparable, V any]() *hashAbierto[K] {
	lista := make([]Tdalista.Lista[hashElem[K]], _CAPACIDAD_INICIAL)

	for i := range lista {
		lista[i] = Tdalista.CrearListaEnlazada[hashElem[K]]()
	}
	return &hashAbierto[K]{lista, _HASH_VACIO, _CAPACIDAD_INICIAL, 0}
}

func crearhashElem[K comparable](clave K, dato int) hashElem[K] {
	return hashElem[K]{clave, dato}
}

func hashFNV1aBytes(datos []byte) uint64 {
	var hash uint64 = 14695981039346656037 // FNV offset basis
	const prime uint64 = 1099511628211     // FNV prime

	// Recorremos el arreglo byte por byte
	for _, b := range datos {
		hash ^= uint64(b) // Hacemos un XOR del hash con el byte actual
		hash *= prime     // Lo multiplicamos por el número primo
	}

	return hash
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func ClaveHaseada(clave any, largo int) uint64 {
	clave_en_bytes := convertirABytes(clave)
	return hashFNV1aBytes(clave_en_bytes) % uint64(largo)
}

func (hash *hashAbierto[K]) buscarCelda(clave K, clave_hasheada uint64) Tdalista.IteradorLista[hashElem[K]] {
	lista := Tdalista.CrearListaEnlazada[hashElem[K]]()

	if hash.tabla[clave_hasheada] != nil {
		return hash.tabla[clave_hasheada].Iterador()
	}

	hash.tabla[clave_hasheada] = lista
	return hash.tabla[clave_hasheada].Iterador()

}

func (hash *hashAbierto[K]) redimensionar(nuevoTam int) {

	nuevo_arreglo := make([]Tdalista.Lista[hashElem[K]], nuevoTam)
	arreglo := make([]Tdalista.Lista[hashElem[K]], 0)

	for _, elemento := range hash.tabla {
		if elemento != nil {
			arreglo = append(arreglo, elemento)
		}
	}

	hash.tabla = nuevo_arreglo
	hash.tam = nuevoTam

	for _, elemento := range arreglo {
		for i := elemento.Iterador(); i.HayAlgoMas(); i.Avanzar() {
			claveValor := crearhashElem(i.VerActual().clave, i.VerActual().dato)
			posicion := hash.buscarCelda(i.VerActual().clave, ClaveHaseada(i.VerActual().clave, nuevoTam))
			posicion.Insertar(claveValor)
		}
	}
}

func (hash *hashAbierto[K]) Guardar(clave K, valor int) {
	factorDeCarga := float64(hash.cantidad) / float64(hash.tam)
	if factorDeCarga >= _FACTOR_DE_CARGA_MAXIMO {
		hash.redimensionar(hash.tam * _DOBLE)
	}
	par_clave_valor := crearhashElem(clave, valor-hash.sumado)
	clave_hasheada := ClaveHaseada(clave, len(hash.tabla))

	for iterador := hash.buscarCelda(clave, clave_hasheada); iterador.HayAlgoMas(); iterador.Avanzar() {
		if iterador.VerActual().clave == clave {
			iterador.Borrar()
			iterador.Insertar(par_clave_valor)
			return
		}
	}

	hash.tabla[clave_hasheada].InsertarUltimo(par_clave_valor)
	hash.cantidad++
}

func (hash *hashAbierto[K]) Pertenece(clave K) bool {

	clave_hasheada := ClaveHaseada(clave, len(hash.tabla))

	if hash.tabla[clave_hasheada] == nil || hash.tabla[clave_hasheada].EstaVacia() {
		return false
	}

	iterador := hash.tabla[clave_hasheada].Iterador()
	for iterador.HayAlgoMas() {
		if iterador.VerActual().clave == clave {
			return true
		}
		iterador.Avanzar()
	}
	return false
}

func (hash *hashAbierto[K]) Obtener(clave K) int {

	if hash.cantidad == _HASH_VACIO {
		panic(_PANICO_HASH)
	}

	clave_hasheada := ClaveHaseada(clave, len(hash.tabla))

	for iterador := hash.tabla[clave_hasheada].Iterador(); iterador.HayAlgoMas(); iterador.Avanzar() {
		if iterador.VerActual().clave == clave {
			return iterador.VerActual().dato + hash.sumado
		}
	}
	panic(_PANICO_HASH)
}

func (hash *hashAbierto[K]) Borrar(clave K) int {

	if hash.cantidad == _HASH_VACIO {
		panic(_PANICO_HASH)
	}

	clave_hasheada := ClaveHaseada(clave, len(hash.tabla))

	iterador := hash.tabla[clave_hasheada].Iterador()
	for iterador.HayAlgoMas() {
		if iterador.VerActual().clave == clave {
			dato := iterador.Borrar().dato
			hash.cantidad--
			factorDeCarga := float64(hash.cantidad) / float64(hash.tam)
			if factorDeCarga <= _FACTOR_DE_CARGA_MINIMO && hash.tam > _CAPACIDAD_INICIAL {
				hash.redimensionar(hash.tam / _DOBLE)
			}
			return dato + hash.sumado
		}
		iterador.Avanzar()
	}

	panic(_PANICO_HASH)
}

func (hash *hashAbierto[K]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashAbierto[K]) Iterar(f func(clave K, dato int) bool) {

	for _, e := range hash.tabla {
		if e == nil {
			continue
		}
		iterador := e.Iterador()
		for iterador.HayAlgoMas() {
			actual := iterador.VerActual()

			if !f(actual.clave, actual.dato) {
				return
			}

			iterador.Avanzar()
		}
	}
}

func (hash *hashAbierto[K]) Iterador() *iterDiccionario[K] {
	iterador := &iterDiccionario[K]{indice: 0, hash: hash, iterLista: hash.tabla[0].Iterador()}
	iterador.siguienteCelda()
	return iterador
}

func (iterador *iterDiccionario[K]) siguienteCelda() {
	for iterador.indice < len(iterador.hash.tabla) {
		if iterador.hash.tabla[iterador.indice] != nil {
			if !iterador.hash.tabla[iterador.indice].EstaVacia() {
				iterador.iterLista = iterador.hash.tabla[iterador.indice].Iterador()
				return
			}
		}
		iterador.indice++
	}
}

func (iterador *iterDiccionario[K]) HayAlgoMas() bool {
	return iterador.indice < len(iterador.hash.tabla) && iterador.iterLista.HayAlgoMas()
}

func (iterador *iterDiccionario[K]) VerActual() (K, int) {
	if !iterador.HayAlgoMas() {
		panic(_PANICO_ITERADOR)
	}
	return iterador.iterLista.VerActual().clave, iterador.iterLista.VerActual().dato + iterador.hash.sumado
}

func (iterador *iterDiccionario[K]) Avanzar() {
	if !iterador.HayAlgoMas() {
		panic(_PANICO_ITERADOR)
	}
	iterador.iterLista.Avanzar()
	if !iterador.iterLista.HayAlgoMas() {
		iterador.indice++
		iterador.siguienteCelda()
	}
}

func (h *hashAbierto[K]) Add(x K, n int) {
	valor := h.Obtener(x)
	h.Guardar(x, valor+n)
}

func (h *hashAbierto[K]) Addall(m int) {
	h.sumado += m
}

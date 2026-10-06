package diccionario

import (
	Tdalista "entrenamiento/tdas/lista"
	"fmt"
)

type hashElem[K comparable, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K comparable, V any] struct {
	tabla    []Tdalista.Lista[hashElem[K, V]]
	cantidad int
	tam      int
}

type iterDiccionario[K comparable, V any] struct {
	indice    int
	hash      *hashAbierto[K, V]
	iterLista Tdalista.IteradorLista[hashElem[K, V]]
}

const _PANICO_HASH = "La clave no pertenece al diccionario"
const _PANICO_ITERADOR = "El iterador termino de iterar"
const _CAPACIDAD_INICIAL = 10
const _HASH_VACIO = 0
const _FACTOR_DE_CARGA_MAXIMO = 3
const _FACTOR_DE_CARGA_MINIMO = 0.25
const _DOBLE = 2

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	lista := make([]Tdalista.Lista[hashElem[K, V]], _CAPACIDAD_INICIAL)

	for i := range lista {
		lista[i] = Tdalista.CrearListaEnlazada[hashElem[K, V]]()
	}
	return &hashAbierto[K, V]{lista, _HASH_VACIO, _CAPACIDAD_INICIAL}
}

func crearhashElem[K comparable, V any](clave K, dato V) hashElem[K, V] {
	return hashElem[K, V]{clave, dato}
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

func (hash *hashAbierto[K, V]) buscarCelda(clave K, clave_hasheada uint64) Tdalista.IteradorLista[hashElem[K, V]] {
	lista := Tdalista.CrearListaEnlazada[hashElem[K, V]]()

	if hash.tabla[clave_hasheada] != nil {
		return hash.tabla[clave_hasheada].Iterador()
	}

	hash.tabla[clave_hasheada] = lista
	return hash.tabla[clave_hasheada].Iterador()

}

func (hash *hashAbierto[K, V]) redimensionar(nuevoTam int) {

	nuevo_arreglo := make([]Tdalista.Lista[hashElem[K, V]], nuevoTam)
	arreglo := make([]Tdalista.Lista[hashElem[K, V]], 0)

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

func (hash *hashAbierto[K, V]) Guardar(clave K, valor V) {
	factorDeCarga := float64(hash.cantidad) / float64(hash.tam)
	if factorDeCarga >= _FACTOR_DE_CARGA_MAXIMO {
		hash.redimensionar(hash.tam * _DOBLE)
	}
	par_clave_valor := crearhashElem(clave, valor)
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

func (hash *hashAbierto[K, V]) Pertenece(clave K) bool {

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

func (hash *hashAbierto[K, V]) Obtener(clave K) V {

	if hash.cantidad == _HASH_VACIO {
		panic(_PANICO_HASH)
	}

	clave_hasheada := ClaveHaseada(clave, len(hash.tabla))

	for iterador := hash.tabla[clave_hasheada].Iterador(); iterador.HayAlgoMas(); iterador.Avanzar() {
		if iterador.VerActual().clave == clave {
			return iterador.VerActual().dato
		}
	}
	panic(_PANICO_HASH)
}

func (hash *hashAbierto[K, V]) Borrar(clave K) V {

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
			return dato
		}
		iterador.Avanzar()
	}

	panic(_PANICO_HASH)
}

func (hash *hashAbierto[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashAbierto[K, V]) Iterar(f func(clave K, dato V) bool) {

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

func (hash *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	iterador := &iterDiccionario[K, V]{indice: 0, hash: hash, iterLista: hash.tabla[0].Iterador()}
	iterador.siguienteCelda()
	return iterador
}

func (iterador *iterDiccionario[K, V]) siguienteCelda() {
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

func (iterador *iterDiccionario[K, V]) HayAlgoMas() bool {
	return iterador.indice < len(iterador.hash.tabla) && iterador.iterLista.HayAlgoMas()
}

func (iterador *iterDiccionario[K, V]) VerActual() (K, V) {
	if !iterador.HayAlgoMas() {
		panic(_PANICO_ITERADOR)
	}
	return iterador.iterLista.VerActual().clave, iterador.iterLista.VerActual().dato
}

func (iterador *iterDiccionario[K, V]) Avanzar() {
	if !iterador.HayAlgoMas() {
		panic(_PANICO_ITERADOR)
	}
	iterador.iterLista.Avanzar()
	if !iterador.iterLista.HayAlgoMas() {
		iterador.indice++
		iterador.siguienteCelda()
	}
}

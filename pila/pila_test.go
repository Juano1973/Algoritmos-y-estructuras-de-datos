package pila_test

import (
	"strconv"
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	// mas pruebas para este caso...
}

func Test_Apilar_Desapilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)

	require.Equal(t, 3, pila.Desapilar())
	require.Equal(t, 2, pila.Desapilar())
	require.Equal(t, 1, pila.Desapilar())
	require.True(t, pila.EstaVacia())
}

func TestLifo(t *testing.T) {
	pila_int := TDAPila.CrearPilaDinamica[int]()
	for j := 0; j <= 100000; j++ {
		pila_int.Apilar(j)

	}
	for i := 100000; i >= 0; i-- {
		var recupero1 int = pila_int.Desapilar()
		if i != 0 {
			var tope int = pila_int.VerTope()
			require.Equal(t, i-1, tope)
		} else {
			require.Panics(t, func() { pila_int.Desapilar() })
			require.Panics(t, func() { pila_int.VerTope() })
			require.True(t, pila_int.EstaVacia())
		}
		require.Equal(t, i, recupero1)

	}

	require.True(t, pila_int.EstaVacia())

	pila_string := TDAPila.CrearPilaDinamica[string]()
	for j := 0; j <= 100000; j++ {
		pila_string.Apilar(strconv.Itoa(j))

	}
	for i := 100000; i >= 0; i-- {
		var recupero1 string = pila_string.Desapilar()
		if i != 0 {
			var tope string = pila_string.VerTope()
			require.Equal(t, strconv.Itoa(i-1), tope)
		} else {
			require.Panics(t, func() { pila_int.Desapilar() })
			require.Panics(t, func() { pila_int.VerTope() })
			require.True(t, pila_string.EstaVacia())
		}
		require.Equal(t, strconv.Itoa(i), recupero1)

	}

}

func TestDesapilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.Panics(t, func() { pila.Desapilar() })
}

func TestVertope(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.Panics(t, func() { pila.VerTope() })
}

func TestApilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for j := 0; j <= 1000; j++ {
		pila.Apilar(j)
	}
	for i := 0; i < 400; i++ {
		pila.Desapilar()
	}
	for x := 0; x <= 20; x++ {
		pila.Apilar(x)
	}
	var tope int = pila.VerTope()
	require.Equal(t, 20, tope)
}

func TestDesapilar_hasta_vacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for j := 0; j <= 400; j++ {
		pila.Apilar(j)
	}
	for i := 0; i <= 400; i++ {
		pila.Desapilar()
	}

	require.True(t, pila.EstaVacia())

	require.Panics(t, func() { pila.VerTope() })
	require.Panics(t, func() { pila.Desapilar() })
}

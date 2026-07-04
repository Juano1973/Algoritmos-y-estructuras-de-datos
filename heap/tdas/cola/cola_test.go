package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.Panics(t, func() { cola.Desencolar() })

	cola.Encolar(1)
	require.False(t, cola.EstaVacia())
	cola.Desencolar()
	require.True(t, cola.EstaVacia())
	require.Panics(t, func() { cola.Desencolar() })
}

func TestEncolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i <= 100; i++ {
		cola.Encolar(i)
	}

	for j := 0; j <= 100; j++ {
		dato := cola.Desencolar()
		require.Equal(t, j, dato)
	}
}

func TestVerPrimero(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	require.Equal(t, 1, cola.VerPrimero())
}

func TestVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for j := 0; j <= 100000; j++ {
		cola.Encolar(j)

	}
	for i := 0; i <= 100000; i++ {
		var recupero1 int = cola.Desencolar()
		if i != 100000 {
			require.Equal(t, i, recupero1)
		} else {
			require.Panics(t, func() { cola.Desencolar() })
			require.Panics(t, func() { cola.VerPrimero() })
			require.True(t, cola.EstaVacia())
		}
		require.Equal(t, i, recupero1)

	}

}

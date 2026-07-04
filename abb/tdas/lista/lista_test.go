package lista_test

import (
	Lista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test del TDAlista
func TestCrearLista(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())

	lista.InsertarPrimero(1)
	require.False(t, lista.EstaVacia())

	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia())

}

func TestVerprimero(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()
	require.Panics(t, func() { lista.VerPrimero() })

	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)
	lista.InsertarUltimo(5)

	require.Equal(t, 3, lista.VerPrimero())
	lista.BorrarPrimero()
	require.Equal(t, 2, lista.VerPrimero())
	lista.BorrarPrimero()
	require.Equal(t, 1, lista.VerPrimero())
	lista.BorrarPrimero()
	require.Equal(t, 5, lista.VerPrimero())

	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia())
}

func TestVerultimo(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(5)
	lista.InsertarPrimero(4)
	lista.InsertarUltimo(6)
	lista.InsertarUltimo(7)
	require.Equal(t, 7, lista.VerUltimo())
	require.Equal(t, 4, lista.VerPrimero())
}

func TestLargo(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()
	for j := 1; j <= 1000; j++ {
		lista.InsertarPrimero(j)
		require.Equal(t, j, lista.VerPrimero())
	}
	require.Equal(t, 1000, lista.Largo())
}

// Test del iterador
func TestIterador(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()

	for i := 0; i <= 10; i++ {
		lista.InsertarPrimero(i)
	}

	iterador := lista.Iterador()

	for j := 10; j > 0; j-- {
		require.Equal(t, j, iterador.VerActual())
		iterador.Avanzar()
	}
}

func TestIteradorBorrar(t *testing.T) {

	lista := Lista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)

	iterador := lista.Iterador()

	require.Equal(t, 3, iterador.VerActual())
	iterador.Avanzar()
	require.Equal(t, 2, iterador.VerActual())

	require.True(t, iterador.HayAlgoMas())
	require.NotPanics(t, func() { iterador.Avanzar() })
	require.Panics(t, func() { iterador.Avanzar() })

}

func TestIteradorInsertarBorrar(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(4)

	iterador := lista.Iterador()
	iterador.Insertar(10)
	require.Equal(t, 10, lista.VerPrimero())
}

func TestVolumenIterador(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()
	volumen := 100000
	for i := 0; i <= volumen; i++ {
		lista.InsertarUltimo(i)
	}
	iterador := lista.Iterador()

	for j := 0; j <= volumen; j++ {
		require.Equal(t, j, iterador.VerActual())
		if iterador.HayAlgoMas() {
			iterador.Avanzar()
		}
	}
}

func TestVolumenIterar(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()
	volumen := 100000
	for i := 0; i <= volumen; i++ {
		lista.InsertarPrimero(i)
	}

	lista.Iterar(func(i int) bool {
		require.Equal(t, volumen, i)
		volumen--
		return true
	})
}

func TestIterarSumarTodos(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()
	contador := 0
	res := 0

	for i := 0; i <= 1000; i++ {
		lista.InsertarPrimero(i)
		contador += i
	}

	lista.Iterar(func(i int) bool {
		res += i
		return true
	})

	require.Equal(t, contador, res)

}

func TestPrimeroIterador(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)

	iterador := lista.Iterador()
	iterador.Insertar(5)
	require.Equal(t, 5, lista.VerPrimero())
}

func TestIteradorInsertarFinal(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	lista.InsertarUltimo(6)
	lista.InsertarUltimo(7)

	iterador := lista.Iterador()
	for iterador.HayAlgoMas() {
		iterador.Avanzar()
	}
	require.False(t, iterador.HayAlgoMas())

	iterador.Insertar(8)

	require.Equal(t, 8, lista.VerUltimo())

	require.NotPanics(t, func() { iterador.Avanzar() })
	require.Panics(t, func() { iterador.Borrar() })

}

func TestBorrarAlcrear(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(5)

	iterador := lista.Iterador()
	iterador.Borrar()
	require.Equal(t, 4, lista.VerPrimero())
	iterador.Borrar()
	require.Equal(t, 3, lista.VerPrimero())
	iterador.Borrar()
	require.Equal(t, 2, lista.VerPrimero())
	iterador.Borrar()
	require.Equal(t, 1, lista.VerPrimero())
	iterador.Borrar()
	require.Panics(t, func() { iterador.Borrar() })
	require.False(t, iterador.HayAlgoMas())

	iterador.Insertar(1)

	iterador.Borrar()
	require.True(t, lista.EstaVacia())
	iterador.Insertar(2)

	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())
}

func TestIteradorInterno(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()
	contador := 0

	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	lista.InsertarUltimo(6)
	lista.InsertarUltimo(7)
	lista.InsertarUltimo(8)
	lista.InsertarUltimo(9)
	lista.InsertarUltimo(10)

	lista.Iterar(func(i int) bool {
		if i == 5 {
			return false
		}
		contador += i
		return true
	})

	require.Equal(t, 1+2+3+4, contador)

}

func TestVolumen(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()
	for e := range 100 {
		lista.InsertarUltimo(e)
	}
	iterador := lista.Iterador()

	for i := range 99 {
		iterador.Avanzar()
		i++
	}
	require.Equal(t, iterador.VerActual(), 99)
}

func TestUnicoElemento(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)

	iterador := lista.Iterador()
	numero := iterador.Borrar()
	require.Equal(t, numero, 1)
}

func TestTresvalores(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iterador := lista.Iterador()

	require.Equal(t, iterador.VerActual(), 1)
	require.True(t, iterador.HayAlgoMas())
	require.NotPanics(t, func() { iterador.Avanzar() })

	require.Equal(t, iterador.VerActual(), 2)
	require.True(t, iterador.HayAlgoMas())
	require.NotPanics(t, func() { iterador.Avanzar() })

	require.Equal(t, iterador.VerActual(), 3)
	require.True(t, iterador.HayAlgoMas())
	require.NotPanics(t, func() { iterador.Avanzar() })

}

func TestBorrarunicoelemento(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	iterador := lista.Iterador()
	require.NotPanics(t, func() { iterador.Borrar() })
	iterador.Insertar(2)
	require.Equal(t, 2, lista.VerPrimero())
}

func TestBorrarListaVacia(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	require.Panics(t, func() { iterador.Borrar() })
}

func TestBorrarEnOrden(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(5)
	lista.InsertarPrimero(6)
	lista.InsertarPrimero(7)

	iterador := lista.Iterador()

	iterador.Insertar(8)
	iterador.Avanzar()
	iterador.Avanzar()
	iterador.Avanzar()
	iterador.Avanzar()
	iterador.Insertar(4)

	// 8 7 6 5 4 3 2 1

	i := 3
	for iterador.HayAlgoMas() {
		iterador.Avanzar()
		if iterador.HayAlgoMas() {
			require.Equal(t, i, iterador.VerActual())
		}
		i--
	}

	for e := 8; e >= 1; e-- {
		require.Equal(t, e, lista.BorrarPrimero())
	}

	require.Panics(t, func() { iterador.Avanzar() })
	require.True(t, lista.EstaVacia())

}

func TestIterarAlFinal(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()

	for e := range 10 {
		lista.InsertarUltimo(e)
	}

	iterador := lista.Iterador()

	for iterador.HayAlgoMas() {
		iterador.Avanzar()
	}

	iterador.Insertar(11)
	require.Equal(t, 11, lista.VerUltimo())

	require.NotPanics(t, iterador.Avanzar)
	require.Panics(t, iterador.Avanzar)
}

func TestInsertaralMedio(t *testing.T) {

	lista := Lista.CrearListaEnlazada[int]()

	for e := range 10 {
		lista.InsertarUltimo(e)
	}

	iter := lista.Iterador()

	for range 5 {
		iter.Avanzar()
	}
	iter.Insertar(99)

	iter2 := lista.Iterador()
	contador := 0
	for contador != 5 {
		iter2.Avanzar()
		contador++
	}
	require.Equal(t, 99, iter2.VerActual())
	iter2.Avanzar()
	require.Equal(t, 5, iter2.VerActual())

}

func TestBorrarPrimero(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)

	iter := lista.Iterador()
	dato := iter.Borrar()
	require.Equal(t, 1, dato)
	require.Equal(t, 2, lista.VerPrimero())
}

func TestBorrarUltimo(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)

	iter := lista.Iterador()
	for iter.HayAlgoMas() {
		iter.Avanzar()
	}
	require.Equal(t, 5, lista.VerUltimo())
	require.Panics(t, func() { iter.Avanzar() })
	require.Panics(t, func() { iter.VerActual() })
	require.Panics(t, func() { iter.Borrar() })

}

func TestInsertarEnMedioOrden(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()

	for e := range 1000 {
		if (e + 1) != 340 {
			lista.InsertarUltimo(e + 1)
		}
	}

	iterador := lista.Iterador()
	for range 339 {
		iterador.Avanzar()
	}
	iterador.Insertar(340)

	iter2 := lista.Iterador()
	for z := range 1000 {
		require.Equal(t, z+1, iter2.VerActual())
		iter2.Avanzar()
	}

	for j := range 1000 {
		require.Equal(t, j+1, lista.BorrarPrimero())
	}
}

func TestInsertar(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(6)
	lista.InsertarUltimo(7)
	lista.InsertarUltimo(8)
	lista.InsertarUltimo(9)
	lista.InsertarUltimo(10)
	lista.InsertarUltimo(11)

	iterador := lista.Iterador()
	for range 4 {
		iterador.Avanzar()
	}
	require.Equal(t, 6, iterador.VerActual())
	iterador.Insertar(5)
	require.Equal(t, 5, iterador.VerActual())
	iterador.Avanzar()
	require.Equal(t, 6, iterador.VerActual())
	iterador.Avanzar()
	require.Equal(t, 7, iterador.VerActual())

	arr := []int{}
	iterador2 := lista.Iterador()
	for iterador2.HayAlgoMas() {
		if iterador2.HayAlgoMas() {
			arr = append(arr, iterador2.VerActual())
		}
		iterador2.Avanzar()
	}
	arrEjemplo := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	require.Equal(t, arr, arrEjemplo)
}

func LlenarListaVacia(t *testing.T) {
	Lista := Lista.CrearListaEnlazada[int]()
	iter := Lista.Iterador()

	require.Panics(t, func() { iter.Borrar() })
	require.Panics(t, func() { iter.Avanzar() })
	require.Panics(t, func() { iter.VerActual() })

	for e := range 10000 {
		iter.Insertar(e + 1)
	}

	iter2 := Lista.Iterador()

	j := 0
	for iter2.HayAlgoMas() {
		require.Equal(t, j+1, iter2.VerActual())
		iter.Avanzar()
		j++
	}
	require.NotPanics(t, func() { iter2.Avanzar() })
	require.Panics(t, func() { iter2.Avanzar() })
}

func TestOrden(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(10)
	lista.InsertarUltimo(20)
	lista.InsertarUltimo(30)
	lista.InsertarUltimo(40)
	iter := lista.Iterador()
	iter.Insertar(1)
	require.Equal(t, 1, lista.VerPrimero())
	iter.Borrar()
	require.Equal(t, 10, lista.VerPrimero())
	iter.Avanzar()

	iter.Avanzar()
	dato := iter.VerActual()
	iter.Insertar(15)
	iter.Avanzar()
	require.Equal(t, dato, iter.VerActual())

	arr := []int{}

	for iter.HayAlgoMas() {
		iter.Avanzar()
	}

	for !lista.EstaVacia() {
		arr = append(arr, lista.BorrarPrimero())
	}

}

func TestIterarConOtro(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	for e := range 10 {
		iter.Insertar(e + 1)
	}

	iter2 := lista.Iterador()
	j := 10
	for j != 1 {
		require.Equal(t, j, iter2.VerActual())
		iter2.Avanzar()
		j--
	}
	require.NotPanics(t, func() { iter2.VerActual() })
}

func TestInsertarListaVacia(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()

	iter := lista.Iterador()
	iter.Insertar(1)
	require.False(t, lista.EstaVacia())

	require.NotPanics(t, func() { iter.VerActual() })
}

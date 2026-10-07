package lista

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la lista no tiene elementos, falso en caso contrario
	EstaVacia() bool

	//Insertar primero inserta el parametro que se le de en la primera posicion de la lista
	InsertarPrimero(T)

	//Insertar primero inserta el parametro que se le de en la ultima posicion de la lista
	InsertarUltimo(T)

	//Borrar primero borra el primer elemento de la lista, en caso de que esta este vacia, levanta panic
	BorrarPrimero() T

	//Ver primero devuelve el primer elemento de la lista, tira panic si la lista esta vacia
	VerPrimero() T

	//Ver primero devuelve el ultimo elemento de la lista, tira panic si la lista esta vacia
	VerUltimo() T

	//Largo devuelve el atriubto "largo" de la lista
	Largo() int

	//Iterar itera la lista mientras se cumpla la condicion de la funcion visitar que recibe
	Iterar(visitar func(T) bool)

	//Iterador genera un interador externo para poder iterar y modificar la lista
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	//Devuelve el elemento actual sobre el que esta parado el iterador
	VerActual() T

	//Esta funcion indica si el siguiente elemento de la lista existe
	HayAlgoMas() bool

	//Esta funcion hace que el iterador avance a la siguiente posicion
	Avanzar()

	//Esta funcion hace que el iterador inserte un elemento en la lista
	Insertar(T)

	//ESta funcion hace que el iterador borre el elemento sobre el que esta parado
	Borrar() T
}

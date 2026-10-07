package main

import (
	TDApila "entrenamiento/tdas/pila"
)

func balanceado(texto string) bool {
	pila_aux := TDApila.CrearPilaDinamica[string]()

	for _, c := range texto {
		letra := string(c)
		if letra == "(" || letra == "[" || letra == "{" {
			pila_aux.Apilar(letra)
		} else {
			if pila_aux.EstaVacia() {
				return false
			}

			tope := pila_aux.Desapilar()
			if tope != obtener_cierre(letra) {
				return false
			}
		}
	}

	if !pila_aux.EstaVacia() {
		return false
	}
	return true

}

func obtener_cierre(caracter string) string {

	switch caracter {
	case ")":
		return "("

	case "]":
		return "["

	case "}":
		return "{"

	}

	return ""
}

/*
func main() {
	fmt.Println(balanceado("[{([])}]"))
	fmt.Println(balanceado("[{}"))
	fmt.Println(balanceado("[(])"))
	fmt.Println(balanceado("()[{}]"))
	fmt.Println(balanceado("()()(())"))
}

*/

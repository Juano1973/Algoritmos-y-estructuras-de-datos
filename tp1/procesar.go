package main

import (
	"fmt"
	"strconv"
	"strings"
	TdaPila "tdas/pila"
)

func asignar_valor(simbolo string) *operacion {
	//a esta cadena le llega un caracter y le asigna un valor
	res := new(operacion)
	res.asociatividad = IZQUIERDA

	switch simbolo {
	case PRNTSIS_ABRIR:
		res.valor = 0
	case SUMA, RESTA:
		res.valor = 1
	case DIVISION, MULTIPLICACION:
		res.valor = 2
	case ELEVAR:
		res.valor = 3
		res.asociatividad = DERECHA
	case PRNTSIS_CERRAR:
		return res
	}

	res.simbolo = simbolo
	return res
}

func comparar_simbolos(operacion1, operacion2 *operacion) bool {
	//Esta funcion devuelve si el caracter de la izuiqerda tiene mas prioridad que el de la derrecha
	return operacion1.valor > operacion2.valor || (operacion1.valor == operacion2.valor && operacion2.asociatividad == IZQUIERDA)
}

func esNumero(cadena string) bool {
	//esta funcion dice si una cadena es un numero
	_, err := strconv.Atoi(cadena)
	if err != nil {
		return false
	}
	return true
}

func aplicarOperacion(arreglo []string, operadores TdaPila.Pila[operacion], valor operacion) []string {
	if !operadores.EstaVacia() && valor.simbolo != PRNTSIS_ABRIR {

		tope := operadores.VerTope()
		for !operadores.EstaVacia() && comparar_simbolos(&tope, &valor) {
			arreglo = append(arreglo, operadores.Desapilar().simbolo)
			if !operadores.EstaVacia() {
				tope = operadores.VerTope()
			}
		}
	}
	return arreglo
}

func inyectarParentesis(operadores TdaPila.Pila[operacion], arreglo []string) []string {
	for (operadores.VerTope()).simbolo != PRNTSIS_ABRIR {
		elemento := operadores.Desapilar().simbolo
		if elemento != "" {
			arreglo = append(arreglo, elemento)
		}
		if operadores.EstaVacia() {
			break
		}
	}
	operadores.Desapilar()
	return arreglo
}

func agregarNumero(arreglo, arreglo_auxiliar []string) ([]string, []string) {
	if len(arreglo_auxiliar) > ARREGLO_VACIO {
		arreglo = append(arreglo, strings.Join(arreglo_auxiliar, ""))
	}

	arreglo_auxiliar = make([]string, ARREGLO_VACIO)
	return arreglo, arreglo_auxiliar
}

func procesar_operacion(cadena string) {

	operadores := TdaPila.CrearPilaDinamica[operacion]()
	res := make([]string, ARREGLO_VACIO)
	aux := make([]string, ARREGLO_VACIO)

	for i := 1; i <= len(cadena); i++ {

		caracter := string(cadena[i-1])

		if caracter == " " || caracter == "" {
			continue
		}

		if esNumero(caracter) {
			aux = append(aux, caracter)
		} else if caracter != PRNTSIS_CERRAR {
			res, aux = agregarNumero(res, aux)

			valor := asignar_valor(caracter)

			if !operadores.EstaVacia() && valor.simbolo != PRNTSIS_ABRIR {
				res = aplicarOperacion(res, operadores, *valor)
			}
			operadores.Apilar(*valor)
		}

		if caracter == PRNTSIS_CERRAR {

			res, aux = agregarNumero(res, aux)
			res = inyectarParentesis(operadores, res)
		}

	}
	res, aux = agregarNumero(res, aux)

	for !operadores.EstaVacia() {
		res = append(res, operadores.Desapilar().simbolo)
	}

	fmt.Println(strings.Join(res, " "))
}

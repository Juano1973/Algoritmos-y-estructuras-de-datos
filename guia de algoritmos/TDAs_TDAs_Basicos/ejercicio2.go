package main

import (
	"fmt"
	"math"
)

type numComplejo struct {
	parte_real       float64
	parte_imaginaria float64
}

func (n *numComplejo) Imprimir() {
	if n.parte_imaginaria >= 0 {
		fmt.Print(n.parte_real, "+", n.parte_imaginaria, "i")
	} else {
		fmt.Print(n.parte_real, "-", n.parte_imaginaria*-1, "i")
	}
}

func CrearNumeroComplejo(real, imaginaria float64) *numComplejo {
	return &numComplejo{parte_real: real, parte_imaginaria: imaginaria}
}

func (n *numComplejo) Multiplicar(otro numComplejo) {
	re := n.parte_real*otro.parte_real - n.parte_imaginaria*otro.parte_imaginaria
	im := n.parte_real*otro.parte_imaginaria + n.parte_imaginaria*otro.parte_real
	num := CrearNumeroComplejo(re, im)
	num.Imprimir()
}

func (n *numComplejo) Sumar(otro numComplejo) {
	num := CrearNumeroComplejo(n.parte_real+otro.parte_real, n.parte_imaginaria+otro.parte_imaginaria)
	num.Imprimir()
}

func (n *numComplejo) ParteReal() float64 {
	return n.parte_real
}

func (n *numComplejo) ParteImaginaria() float64 {
	return n.parte_imaginaria
}

func (n *numComplejo) Modulo() float64 {
	return math.Sqrt(n.parte_real*n.parte_real + n.parte_imaginaria*n.parte_imaginaria)
}

func (n *numComplejo) Angulo() float64 {
	if n.parte_real == 0 {
		if n.parte_imaginaria > 0 {
			return float64(90)
		} else {
			return float64(270)
		}
	}
	return math.Atan(n.parte_imaginaria / n.parte_real)
}

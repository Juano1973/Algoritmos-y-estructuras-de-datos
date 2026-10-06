package main

import (
	"fmt"
)

// EL nombre estaba trolleando el package

type fraccion struct {
	denominador int
	numerador   int
}

func simplificar(num *int, den *int) {
	if *num == 0 {
		return
	}
	// Si el denominador es negativo, entonces invierto.
	// Entonces si el numerador era negativo, ambos quedan positivos,
	// y si no me queda el negativo en el numerador.
	if *den < 0 {
		*num *= -1
		*den *= -1
	}

	var maxPosibleDivisor int
	if abs(*num) < abs(*den) {
		maxPosibleDivisor = abs(*num)
	} else {
		maxPosibleDivisor = abs(*den)
	}
	for i := 2; i <= maxPosibleDivisor; i++ {
		for *num%i == 0 && *den%i == 0 {
			*num /= i
			*den /= i
		}
	}
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func CrearFraccion(numerador, denominador int) *fraccion {
	if denominador == 0 {
		panic("Dividiste por 0 pa")
	}
	simplificar(&numerador, &denominador)
	return &fraccion{numerador: numerador, denominador: denominador}
}

func (f *fraccion) Sumar(otra fraccion) *fraccion {
	num := f.numerador*otra.denominador + otra.numerador*f.denominador
	denom := f.denominador * f.denominador

	return CrearFraccion(num, denom)
}

func (f *fraccion) ParteEntera() int {
	return f.numerador / f.denominador
}

func (f *fraccion) Representacion() string {
	if f.denominador == 1 {
		return fmt.Sprintf("%d", f.numerador)
	} else {
		return fmt.Sprintf("%d/%d", f.numerador, f.denominador)
	}
}

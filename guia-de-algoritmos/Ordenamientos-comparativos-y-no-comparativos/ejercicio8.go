package main

import (
	"fmt"
	"strconv"
)

type pelicula struct {
	nombre string
	genero string
}

func nombre(p pelicula) int {
	num, _ := strconv.Atoi(p.genero)
	return num
}

func countinggsortt_peliculas(arreglo []pelicula, k int, f func(pelicula) int) []pelicula {

	frecuencias := make([]int, k)
	for _, e := range arreglo {
		frecuencias[f(e)] += 1
	}

	inicios := make([]int, k)
	for i := 1; i < k; i++ {
		inicios[i] = inicios[i-1] + frecuencias[i-1]
	}

	ordenado := make([]pelicula, len(arreglo))
	for j := 0; j < len(arreglo); j++ {
		elem := arreglo[j]
		funcELem := f(elem)
		pos := inicios[funcELem]
		inicios[funcELem] += 1
		ordenado[pos] = elem
	}

	return ordenado
}

func main() {
	p1 := new(pelicula)
	p1.nombre = "Donnie Darko"
	p1.genero = "thriller"
	p2 := new(pelicula)
	p2.nombre = "Juno"
	p2.genero = "coming of age"
	p3 := new(pelicula)
	p3.nombre = "The shinning"
	p3.genero = "thriller"
	p4 := new(pelicula)
	p4.nombre = "Labyrinth"
	p4.genero = "fantasia"
	p5 := new(pelicula)
	p5.nombre = "Ferries buellers day off"
	p5.genero = "coming of age"

	arreglo := make([]pelicula, 0)
	arreglo = append(arreglo, *p1)
	arreglo = append(arreglo, *p2)
	arreglo = append(arreglo, *p3)
	arreglo = append(arreglo, *p4)
	arreglo = append(arreglo, *p5)
	arr := countinggsortt_peliculas(arreglo, 100, nombre)
	fmt.Println(arr)
}

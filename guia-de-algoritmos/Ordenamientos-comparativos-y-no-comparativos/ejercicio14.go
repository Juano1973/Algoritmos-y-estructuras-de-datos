package main

type alumnos struct {
	nombre string
	notas  []int
}

func cntsort(arreglo []alumnos, k int, f func(alumnos) int) []alumnos {

	frecuencias := make([]int, k)

	for _, e := range arreglo {
		frecuencias[f(e)] += 1
	}

	inicios := make([]int, 0)
	for j := 1; j < k; j++ {
		inicios[j] = inicios[j-1] + frecuencias[j-1]
	}

	ordenado := make([]alumnos, len(arreglo))

	for i := 0; i < len(arreglo); i++ {
		elem := arreglo[i]
		funcELem := f(elem)
		pos := inicios[funcELem]
		inicios[funcELem] += 1
		ordenado[pos] = elem
	}

	return ordenado
}

func obtener_promedio(al alumnos) int {
	parcial_a := al.notas[0]
	parcial_b := al.notas[1]
	parcial_c := al.notas[2]

	return (parcial_a + parcial_b + parcial_c) / 3

}

func nota_a(al alumnos) int {
	return al.notas[0]
}

func nota_b(al alumnos) int {
	return al.notas[1]
}

func nota_c(al alumnos) int {
	return al.notas[2]
}

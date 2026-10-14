package main

func countinggsortt(arreglo []int, k int, f func(int) int) []int {

	frecuencias := make([]int, k)
	for _, e := range arreglo {
		frecuencias[f(e)] += 1
	}

	inicios := make([]int, k)
	for i := 1; i < k; i++ {
		inicios[i] = inicios[i-1] + frecuencias[i-1]
	}

	ordenado := make([]int, len(arreglo))
	for j := 0; j < len(arreglo); j++ {
		elem := arreglo[j]
		funcELem := f(elem)
		pos := inicios[funcELem]
		inicios[funcELem] += 1
		ordenado[pos] = elem
	}

	return ordenado
}

func parte_a(num int) int {
	return num - (num % 100)
}

func parte_b(num int) int {
	return (num % 100) - (num % 10)
}

func parte_c(num int) int {
	return num % 10
}

/*
func main() {
	arreglo := []int{122, 369, 332, 180, 486, 349, 326, 101}
	ordenado := countinggsortt(arreglo, 100, parte_b)
	ordenado_b := countinggsortt(ordenado, 1000, parte_a)
	ordenado_c := countinggsortt(ordenado_b, 10, parte_c)

	fmt.Println(ordenado_c)
}
*/

package main

func countinggsortt(arreglo []int, k int, f func(int) int) []int {

	frecuencias := make([]int, 0)
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

func main() {
	arreglo := []int{543, 654, 345, 234, 654, 765, 876}
	ordenad := countinggsortt(arreglo, 10, parte_b)
	ordenad_b := countinggsortt(ordenad, 10, parte_a)
	ordenad_c := countinggsortt(ordenad_b, 10, parte_c)

	println(ordenad_c)
}

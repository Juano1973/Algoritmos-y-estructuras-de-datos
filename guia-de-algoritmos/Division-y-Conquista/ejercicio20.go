package main

type joyas struct {
	id   int
	peso int
}

// Implementación de prueba para la función balanza solicitada por el enunciado
func balanza(arr1, arr2 []joyas) int {
	peso1, peso2 := 0, 0

	for _, j := range arr1 {
		peso1 += j.peso
	}
	for _, j := range arr2 {
		peso2 += j.peso
	}

	if peso1 > peso2 {
		return 1
	} else if peso1 < peso2 {
		return -1
	}
	return 0
}

/*
func main() {
	// Se arma un caso de prueba con 5 joyas: 4 imitaciones (peso 10) y 1 verdadera (peso 15)
	cofreDePrueba := []joyas{
		{id: 1, peso: 10},
		{id: 2, peso: 10},
		{id: 3, peso: 15}, // Esta es la joya de la corona
		{id: 4, peso: 10},
		{id: 5, peso: 10},
		{id: 6, peso: 10},
		{id: 7, peso: 10},
		{id: 8, peso: 10},
		{id: 9, peso: 10},
		{id: 10, peso: 10},
	}

	fmt.Println("--- Iniciando prueba del algoritmo ---")
	fmt.Printf("Joya verdadera esperada: ID %d\n", 3)

	// Llamada a tu función principal
	resultado := verdadersJoya(cofreDePrueba)

	fmt.Printf("Joya devuelta por el algoritmo: ID %d (Peso: %d)\n", resultado.id, resultado.peso)
}
*/
func verdadersJoya(cofre_joyas []joyas) joyas {
	return _obtener_joya(cofre_joyas, 0, len(cofre_joyas))
}

func _obtener_joya(cofre []joyas, inicio, fin int) joyas {

	if inicio >= fin-1 {
		return cofre[inicio]
	}

	medio := (inicio + fin) / 2

	var parte1, parte2 []joyas

	if (inicio+fin)%2 == 0 {
		parte1 = cofre[inicio:medio]
		parte2 = cofre[medio:fin]

	} else {
		parte1 = cofre[inicio : medio+1]
		parte2 = cofre[medio:fin]

	}

	pesaje := balanza(parte1, parte2)

	//if impar && pesaje == 0 {
	//	return cofre[medio+1]
	//}

	if pesaje > 0 {
		return _obtener_joya(cofre, inicio, medio)
	}

	return _obtener_joya(cofre, medio, fin)
}

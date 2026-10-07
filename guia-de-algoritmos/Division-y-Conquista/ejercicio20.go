package main

type joyas struct{}

func balanza(arr1, arr2 []joyas) int

func verdadersJoya(cofre_joyas []joyas) joyas {
	return _obtener_joya(cofre_joyas, 0, len(cofre_joyas))
}

func _obtener_joya(cofre []joyas, inicio, fin int) joyas {

	if inicio >= fin-1 {
		return cofre[inicio]
	}

	medio := (inicio + fin) / 2
	parte1 := cofre[inicio:medio]
	parte2 := cofre[medio:fin]

	if balanza(parte1, parte2) > 0 {
		return _obtener_joya(cofre, inicio, medio)
	}

	return _obtener_joya(cofre, medio, fin)
}

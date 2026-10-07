package main

func parteEntera(k int) int {

	return _parteEntera(k, 0, k)
}

func _parteEntera(k, inicio, fin int) int {
	if inicio >= fin || inicio+1 == fin {
		return inicio
	}

	medio := (inicio + fin) / 2
	cuadrado := medio * medio

	if cuadrado == k {
		return medio
	}

	if cuadrado > k {
		return _parteEntera(k, inicio, medio)
	}

	return _parteEntera(k, medio, fin)
}

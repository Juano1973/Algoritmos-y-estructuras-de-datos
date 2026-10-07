package main

func raiz(f func(int) int, a, b int) int {

	if a >= b {
		return a
	}

	medio := (a + b) / 2

	if f(medio) == 0 {
		return medio
	}

	if f(medio)*f(a) > 0 { //Esto sirve para saber si hubo cambio de signo entre el comeinzo y el medio, de forma que se para donde moverme
		return raiz(f, medio+1, b)
	}

	return raiz(f, a, medio-1)
}

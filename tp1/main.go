package main

import (
	"bufio"
	"os"
)

const (
	ARREGLO_VACIO    int    = 0
	SUMA             string = "+"
	RESTA            string = "-"
	MULTIPLICACION   string = "*"
	DIVISION         string = "/"
	ELEVAR           string = "^"
	PRNTSIS_ABRIR    string = "("
	PRNTSIS_CERRAR   string = ")"
	OPERACION_ELEVAR int    = 3
	DERECHA          string = "derecha"
	IZQUIERDA        string = "izquierda"
)

type operacion struct {
	asociatividad string
	simbolo       string
	valor         int
}

func main() {
	leer_archivo()
}

func leer_archivo() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		procesar_operacion(scanner.Text())
	}

}

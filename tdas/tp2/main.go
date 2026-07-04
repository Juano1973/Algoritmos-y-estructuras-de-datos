package tp2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	Clinica "tdas/tp2/clinica"
	Mensajes "tdas/tp2/tp2"
	mensajes "tdas/tp2/tp2"
)

func main() {
	var argumentos []string = os.Args
	archivos := strings.Split(argumentos[1], ",")

	if len(archivos) != 2 {
		fmt.Printf(mensajes.ENOENT_CANT_PARAMS)
		return
	}

	doctores, pacientes := archivos[0], archivos[1]

	clinica, err_docs, err_pacs := Clinica.CrearClinica(doctores, pacientes)
	if err_docs != nil || err_pacs != nil {
		return
	}
	interfazClinica(&clinica)
}

func ejercutar_instruccion(clinic *Clinica.Clinica, operacion, operandos string) {

	oprds := strings.Split(operandos, ",")
	cant_operandos := len(oprds)

	valido := false
	operaciones := []string{"PEDIR_TURNO", "ATENDER_SIGUIENTE", "INFORME"}
	for _, e := range operaciones {
		if e == operacion {
			valido = true
			break
		}
	}

	if !valido {
		arr := []string{operacion, operandos}
		fmt.Printf(Mensajes.ENOENT_FORMATO, strings.Join(arr, " "))
	}

	switch operacion {
	case "PEDIR_TURNO":
		if cant_operandos != 3 {
			fmt.Printf(Mensajes.ENOENT_PARAMS, strconv.Itoa(cant_operandos))
			return
		}
		nombre := oprds[0]
		especialidad := oprds[1]
		urgencia := oprds[2]

		(*clinic).PedirTurno(nombre, especialidad, urgencia)

	case "ATENDER_SIGUIENTE":
		if cant_operandos != 1 {
			fmt.Printf(Mensajes.ENOENT_PARAMS, strconv.Itoa(cant_operandos))
			return
		}

		nombre := oprds[0]
		(*clinic).AtenderSiguiente(nombre)

	case "INFORME":
		if cant_operandos != 2 {
			fmt.Printf(Mensajes.ENOENT_PARAMS, strconv.Itoa(cant_operandos))
			return
		}

		cota_inferior := oprds[0]
		cota_superior := oprds[1]

		(*clinic).CrearInforme(cota_inferior, cota_superior)

	default:
		fmt.Printf(Mensajes.ENOENT_FORMATO, operacion)
	}

}

func interfazClinica(clinic *Clinica.Clinica) {

	for true {

		var operacion []string

		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			if s.Text() == "" {
				break
			}
			operacion = strings.Split(s.Text(), ":")
		}
		err := s.Err()
		if err != nil {
			fmt.Println(err)
		}

		if len(operacion) < 2 {
			fmt.Printf(mensajes.ENOENT_FORMATO, strconv.Itoa(len(operacion)))
			continue
		}

		ejercutar_instruccion(clinic, operacion[0], operacion[1])
	}
}

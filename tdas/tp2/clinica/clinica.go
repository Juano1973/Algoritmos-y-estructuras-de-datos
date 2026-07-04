package tp2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDAcola "tdas/cola"
	TDAheap "tdas/cola_prioridad"
	TDAdiccionario "tdas/diccionario"
	Mensajes "tdas/tp2/tp2"
)

type paciente struct {
	nombre string
	//urgencia bool ni lo use
	anio int
}

type doctor struct {
	nombre                 string
	especialidadPracticada string
	cantAtendidos          int
}

type especialidad struct {
	//nombre    string ni lo use
	urgentes   TDAcola.Cola[*paciente]
	regulares  TDAheap.ColaPrioridad[*paciente]
	cantEspera int
}

type clinica struct {
	doctores       TDAdiccionario.DiccionarioOrdenado[string, *doctor] //abb xq tenemos q iterar por rangos
	pacientes      TDAdiccionario.Diccionario[string, *paciente]
	especialidades TDAdiccionario.Diccionario[string, *especialidad]
}

func CrearClinica(archivo_doctores, archivo_pacientes string) (Clinica, error, error) {
	nuevaClinica := &clinica{doctores: TDAdiccionario.CrearABB[string, *doctor](strings.Compare), pacientes: TDAdiccionario.CrearHash[string, *paciente](), especialidades: TDAdiccionario.CrearHash[string, *especialidad]()}

	err_docs := nuevaClinica.cargar_informacion_doctores(archivo_doctores)
	err_pacs := nuevaClinica.cargar_informacion_pacientes(archivo_pacientes)

	return nuevaClinica, err_docs, err_pacs
}
func compararPac(pac1, pac2 *paciente) int {
	if pac1.anio < pac2.anio {
		return 1
	}
	if pac1.anio > pac2.anio {
		return -1
	}

	return strings.Compare(pac1.nombre, pac2.nombre)
}

func (clinica *clinica) cargar_informacion_doctores(arch_docs string) error {
	archivo_doctores, err_docs := os.Open(arch_docs)

	if err_docs != nil {
		fmt.Printf(Mensajes.ENOENT_ARCHIVO, arch_docs)
		return *new(error)
	}

	defer archivo_doctores.Close()

	s := bufio.NewScanner(archivo_doctores)
	for s.Scan() {
		doc_esp := strings.Split(s.Text(), ",")
		doc, esp := doc_esp[0], doc_esp[1]

		medico := &doctor{especialidadPracticada: esp, nombre: doc}
		clinica.doctores.Guardar(doc, medico)

		if !clinica.especialidades.Pertenece(esp) {
			nuevaEsp := &especialidad{urgentes: TDAcola.CrearColaEnlazada[*paciente](), regulares: TDAheap.CrearHeap[*paciente](compararPac)}
			clinica.especialidades.Guardar(esp, nuevaEsp)
		}

	}
	return nil
}

func (clinica *clinica) cargar_informacion_pacientes(arch_pacs string) error {

	archivo_pacientes, err_pacs := os.Open(arch_pacs)

	if err_pacs != nil {
		fmt.Printf(Mensajes.ENOENT_ARCHIVO, arch_pacs)
		return *new(error)
	}

	defer archivo_pacientes.Close()

	s := bufio.NewScanner(archivo_pacientes)
	for s.Scan() {
		paciente_anio := strings.Split(s.Text(), ",")
		pacient, anios := paciente_anio[0], paciente_anio[1]

		longevidad, err := strconv.Atoi(anios)
		if err != nil {
			fmt.Printf(Mensajes.ENOENT_ANIO, anios)
			return *new(error)
		}

		pac := &paciente{nombre: pacient, anio: longevidad}
		clinica.pacientes.Guardar(pacient, pac)
	}

	return nil

}

func (c *clinica) PedirTurno(pac, esp, urg string) {
	//no c si el error va a aca o ya en el manejo de comandos
	if !c.pacientes.Pertenece(pac) {
		fmt.Printf(Mensajes.ENOENT_PACIENTE, pac)
		return
	}

	pacient := c.pacientes.Obtener(pac)

	if !c.especialidades.Pertenece(esp) {
		fmt.Printf(Mensajes.ENOENT_ESPECIALIDAD, esp)
		return
	}

	espBuscada := c.especialidades.Obtener(esp)

	if urg == "Urgente" {
		espBuscada.urgentes.Encolar(pacient)
	} else {
		espBuscada.regulares.Encolar(pacient)
	}
	espBuscada.cantEspera++
	fmt.Println("Paciente: ", pac, "Encolado", "\n", espBuscada.cantEspera, "paciente(s) en espera para ", esp)
}

func (c *clinica) AtenderSiguiente(doctor string) {

	if !c.doctores.Pertenece(doctor) {
		fmt.Printf(Mensajes.ENOENT_DOCTOR, doctor)
		return
	}

	doc := c.doctores.Obtener(doctor)
	esp := doc.especialidadPracticada
	espBuscada := c.especialidades.Obtener(esp)
	urg := espBuscada.urgentes

	var pacienteAtendido *paciente

	if !urg.EstaVacia() {
		pacienteAtendido = urg.Desencolar()
		espBuscada.cantEspera--
	} else {
		regs := espBuscada.regulares
		pacienteAtendido = regs.Desencolar()
		espBuscada.cantEspera--
	}

	doc.cantAtendidos++
	fmt.Println("Se atiende a ", pacienteAtendido.nombre, "\n", espBuscada.cantEspera, "paciente(s) en espera para ", esp)
}

func (c *clinica) CrearInforme(inicio, fin string) {

	doctorInforme := []*doctor{}

	for iter := c.doctores.IteradorRango(&inicio, &fin); iter.HayAlgoMas(); iter.Avanzar() {
		_, doc := iter.VerActual()
		doctorInforme = append(doctorInforme, doc)
	}

	fmt.Print(len(doctorInforme), Mensajes.DOCTORES_SISTEMA) //como me cago esta mierda, lo hacia en una iteracion sino

	for i, doc := range doctorInforme {
		fmt.Println(i+1, ": ", doc.nombre, "especialidad ", doc.especialidadPracticada, ", ", doc.cantAtendidos, "paciente(s) atendido(s)")

	}
}

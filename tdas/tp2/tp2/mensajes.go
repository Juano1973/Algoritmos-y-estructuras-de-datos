package mensajes

const (
	PACIENTE_ENCOLADO        = "Paciente %s encolado\n"
	CANT_PACIENTES_ENCOLADOS = "%d paciente(s) en espera para %s\n"
	ENOENT_PACIENTE          = "ERROR: no existe el paciente '%s'\n"
	ENOENT_ESPECIALIDAD      = "ERROR: no existe la especialidad '%s'\n"
	ENOENT_URGENCIA          = "Error: grado de urgencia no identificado ('%s')\n"

	PACIENTE_ATENDIDO = "Se atiende a %s\n"
	SIN_PACIENTES     = "No hay pacientes en espera\n"
	ENOENT_DOCTOR     = "ERROR: no existe el doctor '%s'\n"

	DOCTORES_SISTEMA = "%d doctor(es) en el sistema\n"
	INFORME_DOCTOR   = "%d: %s, especialidad %s, %d paciente(s) atendido(s)\n"

	ENOENT_CANT_PARAMS = "No se recibieron los 2 (dos) parametros: <archivo doctores> y <archivo pacientes>\n"
	ENOENT_ARCHIVO     = "No se pudo leer archivo %s\n"
	ENOENT_ANIO        = "Valor no numerico en campo de anio: %s\n"
	ENOENT_FORMATO     = "ERROR: formato de comando incorrecto ('%s')\n"
	ENOENT_CMD         = "ERROR: no existe el comando '%s'\n"

	ENOENT_PARAMS = "ERROR: cantidad de parametros invalidos para comando '%s'\n"
)

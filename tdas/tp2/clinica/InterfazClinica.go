package tp2

type Clinica interface {
	PedirTurno(pac, esp, urg string)

	AtenderSiguiente(doctor string)

	CrearInforme(inicio, fin string)
}

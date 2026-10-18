package main

type ArbolAbb struct {
	hijoder *ArbolAbb
	hijoizq *ArbolAbb
	dato    int
}

func (a *ArbolAbb) sumatoria() int {

	if a.hijoder == nil && a.hijoizq == nil {
		return a.dato
	}

	var izq int = 0
	var der int = 0

	if a.hijoizq != nil {
		izq = a.hijoizq.sumatoria()

	}

	if a.hijoder != nil {
		der = a.hijoder.sumatoria()
	}

	return izq + der + a.dato
}

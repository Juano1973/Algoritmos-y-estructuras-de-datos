package main

func (a *abb[K, V]) dos_hijos() int {

	if a.hijoder == nil && a.hijoizq == nil {
		return 0
	}

	var izq int = 0
	var der int = 0

	if a.hijoizq != nil {
		izq = a.hijoizq.dos_hijos()
	}

	if a.hijoder != nil {
		der = a.hijoder.dos_hijos()
	}

	if a.hijoizq != nil && a.hijoder != nil {
		return 1 + izq + der
	}

	return 0 + izq + der
}

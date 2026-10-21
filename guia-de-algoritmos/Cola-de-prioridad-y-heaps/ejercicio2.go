package main

type heap struct {
	arreglo []int
}

func (h *heap) tres_mayores() (int, int, int) {

	largo := len(h.arreglo)
	if largo == 0 {
		return 0, 0, 0
	}

	mayor := 0
	segundo := 0
	tercero := 0

	if largo >= 1 {
		mayor = h.arreglo[0]
	}

	if largo >= 2 {
		segundo = h.arreglo[1]
	}

	if largo >= 3 {
		tercero = h.arreglo[2]
	}

	return mayor, segundo, tercero

}

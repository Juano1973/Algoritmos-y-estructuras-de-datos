package main

type mamushka struct {
	tam      int
	color    string
	guardada *mamushka
}

func crearMamushka(tam int, color string) *mamushka {
	mam := new(mamushka)
	mam.color = color
	mam.tam = tam

	return mam
}

func (m *mamushka) ObtenerColor() string {
	return m.color
}

func (m *mamushka) Guardar(mam *mamushka) bool {

	if mam == nil || m.tam <= mam.tam {
		return false
	}

	for true {
		if m.guardada == nil {
			m.guardada = mam
			break
		}

		if m.guardada.tam <= mam.tam {
			return false
			//panic("No brother, usted no puede guardal esa mamushka mi helmano")
		}

		m = m.guardada
	}
	return true
}

package main

type composicion struct {
	funciones []func(float64) float64
}

func CrearComposicion() *composicion {
	return &composicion{}
}

func (c *composicion) AgregarFuncion(f func(float64) float64) {
	c.funciones = append(c.funciones, f)
}

func (c *composicion) Aplicar(x float64) float64 {
	var actual float64 = x
	for i := len(c.funciones) - 1; i >= 0; i-- {
		actual = c.funciones[i](actual)
	}
	return actual
}

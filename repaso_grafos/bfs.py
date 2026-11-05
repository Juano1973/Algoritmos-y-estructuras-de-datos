
class Cola:
	def __init__(self):
		...


def bfs(grafo, origen):
	
	padre = {}
	orden = {}
	visitados = set()

	cola = Cola()
	cola.encolar(origen)
	
	while not cola.estaVacia():
		v = cola.desencolar()
		for w in grafo.adyacentes(v):
			if w not in visitados:
				padre[w] = v
				orden[w] = orden[v] + 1
				cola.encolar(w)	
	visitados.add(v)

	return padre, orden





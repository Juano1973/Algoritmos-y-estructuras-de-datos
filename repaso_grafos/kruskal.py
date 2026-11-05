


def obtener_aristas(grafo):

	aristas = []
	visitados = set()

	for v in grafo:
		for w in grafo.adyacentes(v):
			if w not in visitados:
				aristas.append((v,w, grafo.peso(v,w)))
		visitados.add(v)
	return aristas



def kruskal(grafo)


	aristas = sorted(obtener_aristas(grafo))
	conjuntos = unionFind(grado.obtener_vertices())
	arbol = Grafo(False, grafo.obtener_vertices())
	
	for v, w, peso in aristas:
		
		if conjuntos.find(v) == conjuntos.find(w):
			continue

		conjuntos.union(v,w)
		arbol.agregar_arista(v,w,peso)
	
	return arbol
	




def obtener_aristas(grafo):

	aristas = []
	for v in grafo:
		for w in grafo.adyacentes(v):
			aristas.apend((v,w, grafo.peso_arista(v,w)))
	return aristas


def bellam_ford(grafo, origen)


	dist = {}
	padre = {}
	

	aristas = obtener_aristas(grafo)

	for v in grafo:
		dist[v] = float("inf")

	dist[orige] = 0
	padre[origen] = None

	for i in range(len(grafo.obtener_vertices)-1):
		cambio = False
		for origen, destino, peso in aristas:
			distancia_por_aca = dist[origen] + peso
			if distancia_por_aca < dist[destino]
				distancia[destino] = distancia_por_aca
				padre[destino] = origen
				cambio = True
		if not cambio:
			return padre, distancia


	for v, w , peso in aristas:
		if dist[v] + peso < dist[w]:
			return None

	return padre, distancia

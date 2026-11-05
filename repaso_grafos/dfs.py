




def dfs(grafo, origen):

	visitados = set()
	padre = {}
	orden = {}
	
	_dfs(grafo, origen, visitados, padre, orden)

	for v in grafo:
		if v not in visitados:
			_dfs(grafo, v, visitados, padre, orden)

	return padre, orden



def _dfs(grafo, v, visitados, padre, orden):

	visitados.add(v)
	for w in grafo.adyacentes(v):
		if w not in visitados:
			padre[w] = v
			orden[w] = orden[v] + 1
			_dfs(grafo, w, visitados, padre, orden)



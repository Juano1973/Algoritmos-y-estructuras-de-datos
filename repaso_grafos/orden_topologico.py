






def orden_topologico(grafo):

	orden = {v:0 for v in grafo.obtener_vertices()}
	res = []

	for v in grafo.obtener_vertices():
		for w in grafo.adyacentes(v):
			orden[w] = orden.get(w,0) + 1
	
	cola = Cola()

	for v in grafo.obtener_vertices():
		if orden[v] == 0:
			cola.encolar(v)

	while not cola.EstaVacia():
		v = cola.desencolar()
		for w in grafo.adyacentes(v):
			orden[w] -= 1
			if orden[w] == 0:
				cola.encolar(w)
		res.append(v)

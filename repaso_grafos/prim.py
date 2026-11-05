






def prim(grafo):


	v = random.choice(grafo.obtener_vertices())
	arbol = Grafo(False, grafo.obtenr_vertices())
	h = Heap()
	visitados = set()

	visitados.add(v)
	
	for w in grafo.adyacenteS(v):
		heap.encolar((v,w),grafo.peso(v,w))

	while not heap.EstaVacio():
		(v,w), peso = heap.Desencolar()
		
		if w in visitados:
			continue

		arbol.agregar_arista(v,w)
		visitados.add(w)
		
		
		for x in grafo.adyacentes(w):
			if x not in visitados:
				heap.encolar((w,x), grafo.peso(w,x))
		return arbol
	
	
	

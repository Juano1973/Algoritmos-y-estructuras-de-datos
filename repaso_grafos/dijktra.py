



def dijkstra(grafo, origen):

	h = Heap()
	padre = {}
	dist = {}

	for v in grafo.obtener_vertices():
		distacia[v] = float("inf")
	
	distancia[origen] = 0
	padre[origen] = None
	
	heap.encolar((dist[origen],origen))

	while !heap.EstaVacio():
	
		_, v = heap.Desencolar()
		for w in grafo.adyacentes(w):
			distancia_por_aca  = dist[v] +  grafo.obtener_peso(v,w)
			if distancia_por_aca  < dist[w]:
				padre[w] = v
				dist[w] = distancia_por_aca
				heap.encolar((distancai_por_aca,w))
	
	return padre, dist

from grafo import Grafo

class Heap:
    ...

def dijkstra(grafo: Grafo, origen, destino):


    heap = Heap()
    
    distancia = {v: float("inf") for v in grafo.obtener_vertices()}
    padre = {origen: None}
    distancia[origen] = 0


    heap.Encolar((0,origen))

    while not heap.EstaVacio():

        _, v = heap.Desencolar()
        for w in grafo.adyacentes(v):
            distancia_por_aca = distancia[v] + grafo.peso_arista(v,w)
            if distancia_por_aca < distancia[w]:
                distancia[w] = distancia_por_aca
                padre[w] = v
                heap.Encolar((distancia_por_aca,w))


    camino = []
    actual = destino
    while actual != None:
        camino.append(actual)
        actual = padre[actual]

    return camino[::-1]
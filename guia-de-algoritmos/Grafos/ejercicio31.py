from grafo import Grafo


class Heap:
    ...

def dijkstra(grafo: Grafo, origen):

    heap = Heap()
    dist = {v: float("inf") for v in grafo.obtener_vertices()}
    padre = {}
    dist[origen] = 0
    padre[origen] = None

    heap.Encolar((0,origen))
    while not heap.EstaVacio():

        _, v = heap.Desencolar()

        for w in grafo.adyacentes(v):
            dist_por_aca = dist[v] + grafo.peso_arista(v,w)
            if dist_por_aca < dist[w]:
                padre[w] = v
                dist[w] = dist_por_aca
                heap.Encolar((dist_por_aca,w))

    return padre, dist


def cuantas_cambian(grafo: Grafo, s, t):

    grafo_soporte = Grafo(True, grafo.obtener_vertices())

    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            if not grafo_soporte.pertenece_arista(v, w) or grafo_soporte.peso_arista(v, w) > 0:
                

                grafo_soporte.agregar_arista(v,w,0)

                if v not in grafo.adyacentes(w):
                    grafo_soporte.agregar_arista(w,v,1)


    padres, dist = dijkstra(grafo_soporte, s)

    return dist[t]





from grafo import Grafo
from collections import deque


def diametro(grafo: Grafo):
    res = []
    for v in grafo.obtener_vertices():
        res.append(bfs(grafo, v))
    return max(res)


def bfs(grafo: Grafo, origen):

    cola = deque()
    orden = {origen:0}

    cola.append(origen)

    while cola:

        v = cola.popleft()
        
        for w in grafo.adyacentes(v):
            if w not in orden:
                orden[w] = orden[v] + 1
                cola.append(w)


    ord = orden.values()
    return max(ord)
        
from grafo import Grafo
from collections import deque



def orden_topologico(grafo: Grafo):

    res = []
    orden = {v:0 for v in grafo.obtener_vertices()}

    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            orden[w] += 1


    cola = deque()
    for clave, valor in orden.items():
        if valor == 0:
            cola.append(clave)

    while cola:

        actual = cola.popleft()
        for w in grafo.adyacentes(actual):
            orden[w] -= 1
            if orden[w] == 0:
                cola.append(w)

        res.append(actual)

    return res


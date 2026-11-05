from grafo import Grafo

class Cola:
    ...

def orden_topologico(grafo: Grafo):
    res = []
    orden = {v:0 for v in grafo.obtener_vertices()}

    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            orden[w] += 1


    cola = Cola()
    for v in grafo.obtener_vertices():
        if orden[v] == 0:
            cola.Encolar(v)


    while not cola.EstaVacia():
        v = cola.Desencolar()
        for w in grafo.adyacentes(v):
            orden[w] -= 1
            if orden[w] == 0:
                cola.Encolar(w)

        res.append(v)


    return res
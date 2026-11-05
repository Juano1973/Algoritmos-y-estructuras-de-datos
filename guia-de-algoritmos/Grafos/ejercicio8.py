from grafo import Grafo

class Cola:
    ...

def grafos_separacion(grafo: Grafo):


    for v in grafo.obtener_vertices():
        _, orden = bfs(grafo, v)

        for v in orden.values():
            if v > 6:
                return False


    return True


def bfs(grafo, origen):

    cola = Cola()
    padres = {}
    orden = {} 
    visitados = set()

    padres[origen] = None
    orden[origen] = 0

    cola.Encolar(origen)
    visitados.add(origen)

    while not cola.EstaVacia():

        v = cola.Desencolar()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                cola.Encolar(w)
                padres[w] = v
                orden[w] = orden[v] + 1
        

    return padres, orden

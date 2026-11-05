from grafo import Grafo


class Cola():
    ...


def n_distancia(grafo: Grafo, vert, n):

    res = []
    visitados = set()
    cola = Cola()
    orden = {}
    orden[vert] = 0

    cola.Encolar(vert)
    visitados.add(vert)

    while not cola.EstaVacia():

        v = cola.Desencolar()

        if orden[v] > n:
            return res
        

        if orden[v] == n:
            res.append(w)
            continue
        
     
        for w in grafo.adyacentes(v):
            if w not in visitados:
                
                orden[w] = orden[v] + 1
                cola.Encolar(w)
                visitados.add(w)

           

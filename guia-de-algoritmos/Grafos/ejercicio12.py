from grafo import Grafo



def es_no_dirigido(grafo: Grafo):  #O(V x E )

    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            if not v in grafo.adyacentes(w):
                return False


    return True
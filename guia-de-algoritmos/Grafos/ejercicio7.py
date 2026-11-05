from grafo import Grafo


class Cola:
    ...

def grafo_traspuesto(grafo: Grafo):


    nuevo_arbol = Grafo(True,grafo.obtener_vertices())

    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            nuevo_arbol.agregar_arista(w,v)

    return nuevo_arbol

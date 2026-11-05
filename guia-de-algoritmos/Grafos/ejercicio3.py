from grafo import Grafo
from collections import deque

def grados_grafo(grafo: Grafo):
    grados = {}

    visitados = set()
    cola = deque()
    primero = grafo.vertice_aleatorio()
    cola.append(primero)
    grados[primero] = 0
    
    while cola:
        v = cola.popleft()
        for w in v.obtener_aristas():
            if w not in visitados:
                grados[w] = grados[v] + 1
                cola.append(w)
        visitados.add(v)

    return grados

    


def grados_salida_grafo(grafo: Grafo):

    grados = {}

    for v in grafo.obtener_vertices():
        grados[v] = len(grafo.adyacentes(v))

    return grados


def grados_entrada_grafo(grafo: Grafo):


    grados = {}


    for v in grafo.obtener_vertices():
        grados[v] = 0

    for vert in grafo.obtener_vertices():
        for ady in grafo.adyacentes(vert):
            grados[ady] += 1

    return grados





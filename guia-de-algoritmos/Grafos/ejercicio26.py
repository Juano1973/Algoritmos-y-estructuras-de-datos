from grafo import Grafo

class Heap:
    ...


def conseguir_aristas(grafo: Grafo):

    res = []
    visitados = set()

    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            if (v,w) not in visitados:
                res.append((grafo.peso_arista(v,w), v, w))
                visitados.add((v,w))
                visitados.add((w,v))

    return res


"""
PRIM:
.Usa un heap
.Arranca de un random
.Encola sus adyacentes
.Si w se visito, continue
.Encola adyacentes w


Heap
V
Encola adyacentes
W si no en visitados
encola adyacentes

"""
def prim(grafo: Grafo):

    heap = Heap()
    v = grafo.vertice_aleatorio()

    arbol = Grafo(False, grafo.obtener_vertices())

    visitados = set()    
    visitados.add(v)


    for w in grafo.adyacentes(v):
        heap.Encolar((grafo.peso_arista(v,w),v,w))

    while not heap.EstaVacio():

        peso, v, w = heap.Desencolar()
        if w in visitados:
            continue


        arbol.agregar_arista(v,w,peso)
        visitados.add(w)

        for x in grafo.adyacentes(w):
            if x not in visitados:
                heap.Encolar((grafo.peso_arista(w,x),w,x))
    

    return arbol

class UnionFind:
    ...

def kruskal(grafo):

    aristas = sorted(conseguir_aristas(grafo))
    conjuntos = UnionFind(grafo.obtener_vertices())
    arbol = Grafo(False, grafo.obtener_vertices())
    for par in aristas:
        peso, v,w = par
        if conjuntos.find(v) == conjuntos.find(w):
            continue

        conjuntos.union(v,w)
        arbol.agregar_arista(v,w,peso)


    return arbol
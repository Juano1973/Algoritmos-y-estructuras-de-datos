from grafo import Grafo
from collections import deque
from heap import Heap 
import random
from union_find import UnionFind


#Reccorido por ancho
def bfs(grafo, origen):
    visitados = set()
    padres = {} 
    orden = {}
    padres[origen] = None 
    orden[origen] = 0 # orden.Guardar(origen, 0)
    visitados.add(origen)
    q = deque()  # usar deque de collections, NO QUEUE!!
    q.encolar(origen)
    while not q.esta_vacia():
        v = q.desencolar()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                padres[w] = v
                orden[w] = orden[v] + 1
                visitados.add(w)
                q.encolar(w)
    return padres, orden



#Recorrido por profundidad
def dfs(grafo, origen):
    padres = {}
    orden = {}
    visitados = set()
    padres[origen] = None
    orden[origen] = 0
    visitados.add(origen)
    _dfs(grafo, origen, visitados, padres, orden)
    return padres, orden


def _dfs(grafo, v, visitados, padres, orden):
    for w in grafo.adyacentes(v):
        if w not in visitados:
            visitados.add(w)
            padres[w] = v
            orden[w] = orden[v] + 1
            _dfs(grafo, w, visitados, padres, orden)



def camino_minimo_dijkstra(grafo, origen, destino):
    dist = {}
    padre = {}
    for v in grafo:
        dist[v] = float("inf")
    dist[origen] = 0
    padre[origen] = None
    q = Heap()
    q.encolar((0, origen))
    while not q.esta_vacio():
        _, v = q.desencolar()
        if v == destino:
            return padre, dist
        for w in grafo.adyacentes(v):
            distancia_por_aca = dist[v] + grafo.peso(v, w)
            if distancia_por_aca < dist[w]:
                dist[w] = distancia_por_aca
                padre[w] = v
                q.encolar((dist[w], w))
                # o: q.actualizar(w, dist[w])
    return padre, dist


def obtener_aristas(grafo):
    aristas = []
    for v in grafo:
        for w in grafo.adyacentes(v):
            aristas.append((v, w, grafo.peso(v, w)))
    return aristas


# Bellman-Ford
def camino_minimo_bf(grafo, origen):
    distancia = {}
    padre = {}
    for v in grafo:                     # O(V)
        distancia[v] = float("inf")
    distancia[origen] = 0
    padre[origen] = None
    aristas = obtener_aristas(grafo) # O(V + E)
    for i in range(len(grafo)): # for i:= 0; i < len(grafo); i++
        cambio = False
        for origen, destino, peso in aristas:
            if distancia[origen] + peso < distancia[destino]:
                cambio = True
                padre[destino] = origen
                distancia[destino] = distancia[origen] + peso
        if not cambio:
            return padre, distancia

    for v, w, peso in aristas:
        if distancia[v] + peso < distancia[w]:
            return None  # Hay un ciclo negativo (lanzar excep) panic(AAAAAAAAA)
    return padre, distancia



def mst_prim(grafo):
    v = random.choice(grafo.keys())
    visitados = set()
    visitados.add(v)
    q = Heap()
    for w in grafo.adyacentes(v):
        q.encolar((v, w), grafo.peso(v, w))
    arbol = Grafo(es_dirigido=False, lista_vertices=grafo.obtener_vertices())
    while not q.esta_vacio():
        (v, w), peso = q.desencolar()
        if w in visitados:
            continue
        arbol.agregar_arista(v, w, peso)
        visitados.add(w)
        for x in grafo.adyacentes(w):
            if x not in visitados:
                q.encolar((w, x), grafo.peso(w, x))
    return arbol


def obtener_aristas(grafo):
    aristas = []
    visitados = set()
    for v in grafo:
        for w in grafo.adyacentes(v):
            if w not in visitados:
                aristas.append((v, w, grafo.peso(v, w)))
        visitados.add(v)
    return aristas


def mst_kruskal(grafo):
    conjuntos = UnionFind(grafo.obtener_vertices())
    aristas = sorted(obtener_aristas(grafo)) # ordeno de menor a mayor por peso de las aristas
    arbol = Grafo(False, grafo.obtener_vertices())
    for a in aristas: # > O(E log V)
        v, w, peso = a
        if conjuntos.find(v) == conjuntos.find(w):
            continue
        arbol.arista(v, w, peso)
        conjuntos.union(v, w)
    return arbol


def grados_entrada(grafo):
    g_ent = {}
    for v in grafo:
        g_ent[v] = 0
    for v in grafo:
        for w in grafo.adyacentes(v):
            g_ent[w] += 1
    return g_ent

def topologico_grados(grafo):
    g_ent = grados_entrada(grafo)
    q = deque()
    resultado = []
    for v in grafo: # O(V)
        if g_ent[v] == 0:
            q.encolar(v)

    while not q.esta_vacia():
        v = q.desencolar()
        resultado.append(v)
        for w in grafo.adyacentes(v):
            g_ent[w] -= 1
            if g_ent[w] == 0:
                q.encolar(w)
    return resultado

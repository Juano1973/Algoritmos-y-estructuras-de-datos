from grafo import Grafo, grafo_prueba
from collections import deque


def encontrar_ciclo(grafo: Grafo):

    padres = {}

    for v in grafo.obtener_vertices():
        visitados = set()
        padres[v] = None
        valor = dfs(grafo, v, padres, visitados)
        if valor:
            return valor





def dfs(grafo: Grafo, v, padres, visitados):
    visitados.add(v)
    for w in grafo.adyacentes(v):
        if w == padres[v]:
            continue

        if w in visitados:

            ciclo = [w,v]
            actual = v

            while actual != w:
                actual = padres[actual]
                ciclo.append(actual)

            return ciclo

        padres[w] = v
        res = dfs(grafo,w,padres,visitados)
        if res:
            return res

    return None



print(encontrar_ciclo(grafo_prueba))

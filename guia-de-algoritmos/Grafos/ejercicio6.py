from grafo import Grafo


def obtener_aristas(grafo: Grafo):

    res = []
    visitados = set()
    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            if (v,w) not in visitados and (w,v) not in visitados:
                visitados.add((v,w))
                visitados.add((w,v))
                res.append((v,w))

    return res

def es_Arbol(grafo: Grafo):

    cant_vertices = len(grafo.obtener_vertices())
    cant_aristas = len(obtener_aristas(grafo))

    visitados = set()
    ver_rand = grafo.vertice_aleatorio()
    dfs(grafo, ver_rand, visitados)


    return len(visitados) == cant_vertices and cant_aristas == cant_vertices - 1



def dfs(grafo, v, visitados):
    visitados.add(v)
    for w in grafo.adyacentes(v):
        if w not in visitados:
            dfs(grafo,w,visitados)

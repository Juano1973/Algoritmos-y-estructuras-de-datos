from grafo import Grafo, grafo_bipartito, grafo_no_bipartito



def es_bipartito(grafo):

    colores = {}
    padres = {}
    for v in grafo.obtener_vertices():
        if v not in colores:
            colores[v] = 1 
        valor = dfs(grafo, v, colores, padres)
        if valor == False:
            return False


    return True



def dfs(grafo, v, color, padres):

    for w in grafo.adyacentes(v):
        if w in color:
            if color[w] == color[v]:
                return False
        else:
            if color[v] == 1:
                color[w] = 0
            else:
                color[w] = 1
            res = dfs(grafo, w, color, padres)
            if res == False:
                return False

    return True



print(es_bipartito(grafo_bipartito))
print(es_bipartito(grafo_no_bipartito))
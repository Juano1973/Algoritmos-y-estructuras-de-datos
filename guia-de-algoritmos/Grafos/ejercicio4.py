from grafo import Grafo, grafo_prueba




def es_conexo(grafo: Grafo):

    visitados = set()

    v = grafo.vertice_aleatorio()
    dfs(grafo, v, visitados)

    return len(visitados) == len(grafo.obtener_vertices())




def dfs(grafo: Grafo, v, visitados):

    visitados.add(v)
    for w in grafo.adyacentes(v):
        if w not in visitados:
            dfs(grafo,w,visitados)




print(es_conexo(grafo_prueba))
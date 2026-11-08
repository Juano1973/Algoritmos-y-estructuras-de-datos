from grafo import Grafo



def tiene_ciclos_impares(grafo: Grafo):

    
    visitados = set()
    niveles = {}

    for v in grafo.obtener_vertices():
        
        if v in visitados:
            continue

        visitados.add(v)

        tiene = dfs(grafo,v, visitados, 0, niveles)
        if tiene:
            return True

    return False



def dfs(grafo, v, visitados, nivel, niveles):

    
    for w in grafo.adyacentes(v):
        
        if w not in visitados:
            visitados.add(w)
            niveles[w] = nivel
            tiene = dfs(grafo,w, visitados, nivel + 1, niveles)
            if tiene:
                return True
            
        else:
            return (nivel - niveles[w]) % 2 == 0

    return False 
        

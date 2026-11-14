from grafo import Grafo




def impar_par(grafo: Grafo):

    grado = {v:0 for v in grafo.obtener_vertices()}

    for v in grafo.obtener_vertices():
        grado[v] = len(grafo.adyacentes(v))


    impar = 0
    for _, valor in grado.items():
        if valor % 2 != 0:
            impar += 1


    return impar % 2 == 0




"""
Esta implementacion es O(v)

Si estuviera implementado con matriz de adyacencia seria
simplemente contar cuantas adyacencias tiene cada vertice

Eso seria, por cada vertice, ver todos los otros vertices,
lo que daria una complejidad de O(V x V) = O(V^2)


"""
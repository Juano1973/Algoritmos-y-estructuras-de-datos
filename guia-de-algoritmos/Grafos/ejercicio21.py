from grafo import Grafo





def fauna_en_peligro(grafo: Grafo):

    grados_entrada = {v:0 for v in grafo.obtener_vertices()}
    dependencias = {v:[] for v in grafo.obtener_vertices()}
    res = []

    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            grados_entrada[w] += 1
            dependencias[w].append(v)


    #condicion 1: todos los depredadores de x tienen a x como unico grado de salida

    for clave, valor in grados_entrada.items():

        rompe_ecosistema = True
        for v in dependencias[clave]:
            if len(grafo.adyacentes(v)) != 1:
                rompe_ecosistema = False
                break
        if rompe_ecosistema:
            res.append(clave)


    #condicion 2: todas las presas de x tienen a x como unico grado de entrada

    for v in grafo.obtener_vertices():
        rompe_ecosistema = True
        for w in grafo.adyacentes(v):
            if len(dependencias[w]) != 1:
                rompe_ecosistema = False
                break

        if rompe_ecosistema:
            res.append(v)


    
        
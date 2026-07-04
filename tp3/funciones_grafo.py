from collections import deque
from tdas.grafo import Grafo
from tdas.heap import MinHeap
from tdas.pila import Pila
import random



# CAMINOS MINIMOS

def dijkstra(grafo, origen, destino):

    dist = {}
    padre = {}

    for v in grafo.obtener_vertices():
        dist[v] = float("inf")
    
    dist[origen] = 0
    padre[origen] = None

    q = MinHeap()

    q.Encolar((0, origen))

    while not q.EstaVacio():
        _,v = q.Desencolar()
        if v == destino:
            return armar_camino(dist, padre, origen, destino)
        for w in grafo.adyacentes(v):
            distancia_por_aca = dist[v] + grafo.peso_arista(v,w)
            if distancia_por_aca < dist[w]:
                dist[w] = distancia_por_aca
                padre[w] = v
                q.Encolar((dist[w], w))

   
    return armar_camino(dist, padre, origen, destino)

def pesoGrafo(grafo: Grafo):
    inicio = random.choice(grafo.obtener_vertices())
    cola = deque()
    cola.append(inicio)
    visitados = set()
    sumatoria = 0
    while cola:
        v = cola.popleft()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                sumatoria += grafo.peso_arista(v,w)
                cola.append(w)
        visitados.add(v)
    
    return sumatoria


def obtener_aristas(grafo):
    res = []
    visitados = set()

    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            if w not in visitados:
                res.append((v,w, grafo.peso_arista(v,w)))
        visitados.add(v)

    return res
    
# ORDEN TOPOLOGICO

def grados_entrada(grafo):
    g_ent = {}
    for v in grafo.obtener_vertices():
        g_ent[v] = 0
    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            g_ent[w] += 1
    return g_ent


def topologico_grados(grafo):
    g_ent = grados_entrada(grafo)
    q = deque()
    resultado = []

    for v in grafo.obtener_vertices(): # O(V)
        if g_ent[v] == 0:
            q.append(v)

    while q:
        v = q.popleft()
        resultado.append(v)
        for w in grafo.adyacentes(v):
            g_ent[w] -= 1
            if g_ent[w] == 0:
                q.append(w)

    return resultado



def mst_prim(grafo, origen):

    sumatoria = 0
    visitados = set()
    visitados.add(origen)
    arbol = Grafo(es_dirigido=False, vertices_iniciales=grafo.obtener_vertices())

    heap = MinHeap()

    for w in grafo.adyacentes(origen):
        heap.Encolar((grafo.peso_arista(origen,w), origen,w))

    while not heap.EstaVacio():
        peso,v,w = heap.Desencolar()
        if w in visitados:
            continue
        arbol.agregar_arista(v,w,peso)
        sumatoria += peso
        visitados.add(w)
        for x in grafo.adyacentes(w): 
            if not x in visitados:
                heap.Encolar((grafo.peso_arista(w, x), w,x))
    return arbol, sumatoria


#Ciclo Euleriano
def tiene_ciclo_euleriano(grafo, origen): 
    if not es_conexo(grafo, origen):
        return False
    
    for v in grafo.obtener_vertices():
        if len(grafo.adyacentes(v)) % 2 != 0:
            return False
    return True 


def dfs_ciclo(grafo, v, origen, camino, tiempo):

    for w in grafo.adyacentes(v):

        if not grafo.estan_unidos(v, w):
            continue

        peso = grafo.peso_arista(v, w)
        grafo.borrar_arista(v, w)
        tiempo[0] += peso

        dfs_ciclo(grafo, w, origen, camino, tiempo)

    camino.append(v)

def insertar_ciclo(ciclo, nuevo, vertice):

    res = []

    for v in ciclo:

        if v == vertice:

            for x in nuevo[:-1]:
                res.append(x)

        else:
            res.append(v)

    return res

def buscar_vertice(grafo, ciclo):

    for v in ciclo:
        if len(grafo.adyacentes(v)) > 0:
            return v

    return None

def hierholzer(grafo, origen):

    if not tiene_ciclo_euleriano(grafo, origen):
        return None, None

    copia = copiar_grafo(grafo)

    pila = Pila()
    pila.Apilar(origen)

    camino = []
    tiempo = 0

    while not pila.EstaVacia():

        v = pila.VerTope()

        ady = copia.adyacentes(v)

        if len(ady) == 0:
            camino.append(pila.Desapilar())
            continue

        w = ady[0]

        tiempo += copia.peso_arista(v, w)

        copia.borrar_arista(v, w)

        pila.Apilar(w)

    return camino[::-1], tiempo
    
#def hierholzer(grafo, origen):
#
#    if not tiene_ciclo_euleriano(grafo, origen):
#        return None, None
#
#    copia_grafo = copiar_grafo(grafo)
#
#    tiempo = [0]
#    ciclo = []
#
#    dfs_ciclo(copia_grafo, origen, origen, ciclo, tiempo)
#
#    ciclo.reverse()
#
#    while True:
#
#        vert = buscar_vertice(copia_grafo, ciclo)
#        if vert == None:
#            break
#
#        nuevo = []
#
#        dfs_ciclo(copia_grafo, vert, vert, nuevo, tiempo)
#
#        nuevo.reverse()
#        ciclo = insertar_ciclo(ciclo, nuevo, vert)
#
#    return ciclo, tiempo[0]

#def hierholzer(grafo, origen):
#
#    if not tiene_ciclo_euleriano(grafo, origen):
#        return None, None
#
#    copia = copiar_grafo(grafo)
#
#    camino = []
#    tiempo = [0]
#
#    dfs_ciclo(copia, origen, origen, camino, tiempo)
#
#    return camino[::-1], tiempo[0]
#
#    
def copiar_grafo(grafo):

    copia = Grafo(es_dirigido=False,vertices_iniciales=grafo.obtener_vertices())

    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            if not copia.estan_unidos(v, w):
                copia.agregar_arista(v, w, grafo.peso_arista(v, w))

    return copia


#def holhiezerr(grafo, origen):
#    if not tiene_ciclo_euleriano(grafo,origen):
#        return None, None
#    camino = []
#    ciclo = dfs_ciclo(origen, camino) #devuelve cuando toca devuelta a v(origen)
#
#    while True:
#        vertice = vertice_con_aristas(ciclo)
#
#        if vertice is None:
#            break
#        
#        nuevo_ciclo = dfs_ciclo(vertice) #no considera las aristas ya utilizadas
#
#        ciclo = agregar_ciclo(ciclo, nuevo_ciclo, )
#    return ciclo
#
#
#def dfs_ciclo(v, origen, camino, grafo):
#    for w in grafo.adyacentes(origen):
#        if w == origen:
#            return v  
#        if w not in visitados:
#            
#            visitados.add(w)
#            padres[w] = v
#            orden[w] = orden[v] + 1
#            
#            siguiente = _dfs(grafo, w, visitados, padres, orden, origen, anterior, res)
#            if siguiente != None:
#                return siguiente
#
#
#def _dfs(grafo, v, visitados, padres, orden,origen, anterior, res):
#    
#        
#    for w in grafo.adyacentes(v):
#        if w == origen:
#            return v  
#        if w not in visitados:
#            
#            visitados.add(w)
#            padres[w] = v
#            orden[w] = orden[v] + 1
#            
#            siguiente = _dfs(grafo, w, visitados, padres, orden, origen, anterior, res)
#            if siguiente != None:
#                return siguiente
#

#def holhiezer(grafo, origen):
#    """
#    1) Elegimos cualquier vértice v (o uno en particular, si la aplicación así lo requiere).
#    
#    
#    2) Aplicamos un recorrido DFS desde v, considerando las aristas que estamos utilizando. 
#    Cuando nos volvemos a topar con v (y es imposible que no suceda), nos quedamos con ese camino cerrado.
#    Llamemos a este camino C.
#    
#    
#    3) Si existe algún vértice u dentro del camino C que tenga aristas que no se utilizaron aún,
#    realizar el mismo recorrido mencionado antes desde u (sin considerar las aristas ya utilizadas),
#    y al terminar “agregar” el circuito hecho desde y hasta u en nuestro camino C obtenido antes,
#    en el lugar donde estaba u. Volver a realizar esto hasta que ya no queden aristas sin revisar.
#    
#    """
#
#    if not tiene_ciclo_euleriano(grafo,origen):
#        return None, None
#    
#    res = []
#    orden = {}
#    sumatoria = {"suma":0}
#
#    for v in grafo.obtener_vertices():
#        for w in grafo.adyacentes(v):
#            orden[(v,w)] = orden.get((v,w),0) + 1
#    dfs_euler(grafo,origen,orden, res,sumatoria)
#
#    return res[::-1],sumatoria["suma"]



#def dfs_euler(grafo: Grafo,origen,orden, res,sumatoria):
#
#    for w in grafo.adyacentes(origen):
#        if orden[(origen, w)] > 0 and orden[(w,origen)] > 0:
#            sumatoria["suma"] += grafo.peso_arista(origen,w)
#            orden[(origen, w)] -= 1
#            orden[(w, origen)] -= 1
#            dfs_euler(grafo,w,orden,res,sumatoria)
#    res.append(origen)



def armar_camino(distancia, padre,origen,destino):
    camino = []
    sumatoria = 0

    actual = destino
    sumatoria = distancia[destino]

    while actual != None:
        camino.append(actual.strip())
        
        actual = padre[actual]

    return camino[::-1], sumatoria



def es_conexo(grafo: Grafo, origen):
    visitados = set()
    cola = deque()
    visitados.add(origen)
    cola.append(origen)
    while cola:
        v = cola.popleft()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                cola.append(w)
        visitados.add(v)
    return len(visitados) == len(grafo.obtener_vertices())




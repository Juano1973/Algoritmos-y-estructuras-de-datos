
import random

class Grafo:
    def __init__(self, es_dirigido = False, vertices_iniciales = []):
        self.es_dirigido = es_dirigido
        self.vertices = {}

        for v in vertices_iniciales:
            self.agregar_vertice(v)
        
    def agregar_vertice(self,v):
        self.vertices[v] = {}

    def borrar_vertice(self,v):
        if not self.existe_vertice(v):
            raise ValueError("El vertice no existe")
        
        del self.vertices[v]

        adyacentes = self.adyacentes(v)

        for w in adyacentes:
            del self.vertices[w][v]


    def agregar_arista(self,v,w,peso = 1):
        if not self.es_dirigido:
            self.vertices[w][v] = peso
        
        self.vertices[v][w] = peso

    def borrar_arista(self,v,w):
        if not self.es_dirigido:
            del self.vertices[w][v]
        del self.vertices[v][w]

    def estan_unidos(self,v,w):
        return w in self.vertices[v]

    def peso_arista(self,v,w):
        return int(self.vertices[v][w])

    def obtener_vertices(self):
        res = []
        for v in self.vertices:
            res.append(v)
        return res

    def adyacentes(self,v):
        res = []
        for w in self.vertices[v]:
            res.append(w)

        return res
    
    def existe_vertice(self,id):
        return id in self.vertices

    def vertice_aleatorio(self):
        return random.choice(Grafo.obtener_vertices(self))





grafo_prueba = Grafo(es_dirigido=False, vertices_iniciales=['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'])

# Agregamos aristas con distintos pesos para probar algoritmos de caminos (Dijkstra/Prim/Kruskal)
grafo_prueba.agregar_arista('A', 'B', peso=4)
grafo_prueba.agregar_arista('A', 'C', peso=2)
grafo_prueba.agregar_arista('B', 'C', peso=1)
grafo_prueba.agregar_arista('B', 'D', peso=5)
grafo_prueba.agregar_arista('C', 'D', peso=8)
grafo_prueba.agregar_arista('C', 'E', peso=10)
grafo_prueba.agregar_arista('D', 'E', peso=2)
grafo_prueba.agregar_arista('D', 'F', peso=6)
grafo_prueba.agregar_arista('G', 'H', peso=3)
grafo_prueba.agregar_arista('G', 'F', peso=3)



# ==========================================
# 1. GRAFO BIPARTITO (Ciclo par de 4 vértices)
# ==========================================
# Conjunto 1: {'A', 'C'} | Conjunto 2: {'B', 'D'}
# Ningún vértice se conecta con uno de su mismo grupo.
grafo_bipartito = Grafo(es_dirigido=False, vertices_iniciales=['A', 'B', 'C', 'D'])
grafo_bipartito.agregar_arista('A', 'B')
grafo_bipartito.agregar_arista('B', 'C')
grafo_bipartito.agregar_arista('C', 'D')
grafo_bipartito.agregar_arista('D', 'A')


# ==========================================
# 2. GRAFO NO BIPARTITO (Contiene un triángulo - ciclo impar)
# ==========================================
# Al tener un ciclo de 3 vértices, es imposible pintarlo de 
# dos colores sin que choquen dos adyacentes.
grafo_no_bipartito = Grafo(es_dirigido=False, vertices_iniciales=['X', 'Y', 'Z'])
grafo_no_bipartito.agregar_arista('X', 'Y')
grafo_no_bipartito.agregar_arista('Y', 'Z')
grafo_no_bipartito.agregar_arista('Z', 'X')
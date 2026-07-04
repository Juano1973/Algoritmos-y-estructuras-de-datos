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
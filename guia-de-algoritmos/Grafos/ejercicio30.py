from grafo import Grafo
from collections import deque

#Para este ejercicio simplemente pense en un dijsktra y pense en que me beneficiaba tener informacion extra 

class HeapBin:
    def __init__(self):
        self.peso_uno = deque()
        self.peso_dos = deque()

    def Encolar(self, peso, distancia, vertice):
        if peso == 1:
            self.peso_uno.append((peso,distancia,vertice))
            return
        
        self.peso_dos.append((peso,distancia,vertice))

    def Desencolar(self):
        if len(self.peso_uno) != 0:
            return self.peso_uno.popleft()

        elif len(self.peso_dos) != 0:
            return self.peso_dos.popleft()

        return None

    def EstaVacio(self):

        return len(self.peso_uno) == 0 and len(self.peso_dos) == 0
        
    
#Este es un tda donde hay dos pilas, una para 1 y otra para 2, es como un heap, pero en 0(1)

def dijkstra_peso_binario(grafo: Grafo, origen):

    distancia = {v: float("inf") for v in grafo.obtener_vertices()}
    padre = {}
    heap = HeapBin()
    heap.Encolar((0, 0,origen))
    padre[origen] = None
    distancia[origen] = 0

    while not heap.EstaVacio():

        _,_, v = heap.Desencolar()
   
        for w in grafo.adyacentes(v):
            distancia_por_acac = distancia[v] + grafo.peso_arista(v,w)
            if distancia_por_acac < distancia[w]:
                distancia[w] = distancia_por_acac
                padre[w] = v
                heap.Encolar((grafo.peso_arista(v,w),distancia_por_acac,w))
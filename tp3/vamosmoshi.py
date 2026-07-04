#!/usr/bin/python3

import sys
from tdas.grafo import Grafo
import funciones_grafo as funcs
import csv

ERROR = "No se encontro recorrido"
COMANDO_IR = "ir"
COMANDO_ITINERARIO = "itinerario"
COMANDO_VIAJE = "viaje"
COMANDO_RC = "reducir_caminos"
TIEMPO = "Tiempo total:"
PESO = "Peso total:"
SEPARADOR = " -> "


archivo = sys.argv[1]

        
def obtener_aristas(grafo):
    aristas = []
    visitados = set()
    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            if w not in visitados:
                aristas.append((v, w, grafo.peso_arista(v, w)))
        visitados.add(v)
    return aristas
    
def escribir_en_archivo_pajek(archivo, arbol_MST: Grafo, diccionario):
    with open(archivo, "w") as arch:
        arch.write(f"{len(arbol_MST.obtener_vertices())}\n")
        for v in arbol_MST.obtener_vertices():
            latitud, altitud = ciudades[v]
            arch.write(f"{v},{latitud},{altitud}")
        arch.write(f"{len(obtener_aristas(arbol_MST))}\n")
        for origen,destino,peso in obtener_aristas(arbol_MST):
            arch.write(f"{origen},{destino},{peso}\n")

grafo = Grafo(es_dirigido=False)
ciudades = {}
def escribir_en_archivo(archivo, camino, diccionario):
    with open(archivo, "w") as arch:
        arch.write('<?xml version="1.0" encoding="UTF-8"?>\n')
        arch.write('<kml xmlns="http://earth.google.com/kml/2.1">\n')
        arch.write("\t<Document>\n")
        arch.write("\t\t<name>Resultado</name>\n")
        arch.write("\t\t<description>Resultado de la instruccion pedida.</description>\n")
        lat_ant, lon_ant, ant = None, None, None

        escritos = set()
        lista = []
        for lugar in camino:
            if lugar in escritos:
                continue
            escritos.add(lugar)
            arch.write("\t\t<Placemark>\n")
            arch.write(f"\t\t\t<name>{lugar.rstrip()}</name>\n")
            arch.write("\t\t\t<Point>\n")
            latitud, altitud = diccionario[lugar]
            lugar = str
            arch.writelines(f"\t\t\t\t<coordinates>{altitud.rstrip()}, {latitud.rstrip()}</coordinates>\n")
            lista.append((altitud,latitud))
            arch.write("\t\t\t</Point>\n")
            arch.write("\t\t</Placemark>\n")

        for lugar in camino:
            latitud, altitud = diccionario[lugar]
            if lat_ant == None and lon_ant == None:
                lat_ant = latitud
                lon_ant = altitud
                ant = lugar
                continue
            arch.write("\t\t<Placemark>\n")
            #arch.write(f"\t\t\t<name>Recorrido</name>\n")
            arch.write("\t\t\t<LineString>\n")
            arch.writelines(f"\t\t\t\t<coordinates>{lon_ant.rstrip()},{lat_ant.rstrip()} {altitud.rstrip()},{latitud.rstrip()}</coordinates>\n")
            arch.write("\t\t\t</LineString>\n")
            arch.write("\t\t</Placemark>\n")
            lat_ant = latitud
            lon_ant = altitud
            ant = lugar

        arch.write("\t</Document>\n")
        arch.write("</kml>\n")

def cargar_ciudades(archivo):
    with open(archivo, "r") as arch:

        aristas = []

        indice = 1
        lineas = arch.readlines()#[1:]
        primera_mitad = False
        for linea in lineas:

            if len(linea.split(",")) == 1 and primera_mitad:
                break

            if len(linea.split(",")) == 1 and not primera_mitad:
                primera_mitad = True
                continue

            
            nombre, lat, lon = linea.split(",")
            grafo.agregar_vertice(nombre) 
            ciudades[nombre] = (lat, lon)
            indice += 1
 

        for aristas in lineas[indice+1:]:
            ciudad_1, ciudad_2, tiempo = aristas.split(",")
            grafo.agregar_arista(ciudad_1,ciudad_2, tiempo)

        arch.close()

def normalizar_cadena(cadena):
    
    cadena_vacia = ""
    partes = cadena.split(" ")
    for palabras in partes:
        if palabras == "":
            continue
        cadena_vacia += palabras
        if palabras != partes[-1]:
            cadena_vacia += " "
    return cadena_vacia

def ejecutar_comandos():
    try:
        entrada = input()
    except EOFError:
        return ""

    if entrada.startswith(COMANDO_IR): 
        operacion = entrada.split(",")
        origen = normalizar_cadena(operacion[0].replace(COMANDO_IR," "))
        destino = normalizar_cadena(operacion[1])
        archivo = operacion[2].replace(" ","")
        try:
            camino, distancia = funcs.dijkstra(grafo, origen, destino)
        except:
            print(ERROR)
            return

        if len(camino) == 0:
            print(ERROR)
            return

        print(SEPARADOR.join(camino))
        print(TIEMPO,distancia)
        escribir_en_archivo(archivo, camino, ciudades)

    if entrada.startswith(COMANDO_VIAJE): 
        operacion = entrada.split(",")
        origen = normalizar_cadena(operacion[0].replace(COMANDO_VIAJE,""))
        archivo = operacion[1].replace(" ","")

        camino, tiempo = funcs.hierholzer(grafo,origen)
        if camino == None:
            print(ERROR)
            return
        

        print(SEPARADOR.join(camino))
        print(TIEMPO, tiempo)
        escribir_en_archivo(archivo, camino, ciudades)
        
    if entrada.startswith(COMANDO_ITINERARIO): 

        grafo_dep = Grafo(es_dirigido=True, vertices_iniciales=grafo.obtener_vertices())

        archivo_recomendaciones = entrada.split()[1]
        with open(archivo_recomendaciones, mode="r") as arch:
            lector = csv.reader(arch, delimiter=",")
            for ciudad_1, ciudad_2 in lector:
                grafo_dep.agregar_arista(ciudad_1,ciudad_2)
        
        orden_dep = funcs.topologico_grados(grafo_dep)
        if len(orden_dep) < len(grafo_dep.obtener_vertices()):
            print(ERROR)
            return
        print(SEPARADOR.join(orden_dep))
        

    if entrada.startswith(COMANDO_RC): 
        archivo_res = entrada.split()[1]
        nuevo_grafo, peso = funcs.mst_prim(grafo, grafo.vertice_aleatorio())
        print(PESO, peso)
        escribir_en_archivo_pajek(archivo_res,nuevo_grafo, ciudades)
    
def main():
    cargar_ciudades(archivo)
    while True:
        entrada = ejecutar_comandos()
        if entrada == "":
            break

if __name__ == "__main__":
    main()




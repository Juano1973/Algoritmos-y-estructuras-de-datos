class Pila :
	'''Python implementation the stack'''
	
	def __init__(self):
		self.items = []
	
	def EstaVacia(self):
		return self.items == []
		
	def Apilar(self,item):
		self.items.append(item)			
		
	def Desapilar(self):
		return self.items.pop()
		
	def VerTope(self):
		return self.items[len(self.items)-1]
	
	def Largo(self):
		return len(self.items)
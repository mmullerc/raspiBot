import time

class perpetualTimer():
	def __init__(self,t,callback):
		self.t=t
		self.callback = callback
		self.running = False

	def start(self):
		self.running = True

	def run(self):
		while self.running:
			time.sleep(self.t)
			self.callback()

	def stop(self):
		self.running = False

	def isRunning(self):
		return self.running
from color_sensor import sense_colors
import json
from threading import Timer, Thread, Event

class perpetualTimer():
	def __init__(self,t,hFunction):
		self.t=t
		self.hFunction = hFunction
		self.thread = Timer(self.t,self.handle_function)

	def handle_function(self):
		self.hFunction()
		self.thread = Timer(self.t,self.handle_function)
		self.thread.start()

	def start(self):
		self.thread.start()

	def cancel(self):
		self.thread.cancel()

def getColor():
    r, g, b = sense_colors()
    data = {}
    data['red'] = r
    data['green'] = g
    data['blue'] = b
    json_data = json.dumps(data)
    print('Color: red={0} green={1} blue={2}'.format(r, g, b))
    return json_data

t = perpetualTimer(1,getColor)


def startReading():
	global t
	t.start()
	return getColor()

def stopReading():
	global t
	t.cancel()
	return "Not reading colors"
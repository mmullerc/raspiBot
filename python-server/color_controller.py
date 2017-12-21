from color_sensor import sense_colors
import json
from threading import Timer, Thread, Event
import requests

urlSetColor = "http://localhost:8080/setcolor"

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

	def isRunning(self):
		return self.thread.isAlive()

#read sensor
def getColor():
    r, g, b = sense_colors()
    data = {}
    data['red'] = r
    data['green'] = g
    data['blue'] = b
    json_data = json.dumps(data)
    print('Color: red={0} green={1} blue={2}'.format(r, g, b))

    if(r == 100 & g == 100 & b == 100):
    	setColor(r,g,b)

    return json_data

#thread global variable
t = perpetualTimer(1,getColor)

#starts thread and returns first color read
def startReading():
	global t
	if t.isRunning() is not True:
		t.start()
	return getColor()

#stops the thread
def stopReading():
	global t
	t.cancel()
	return "Not reading colors"

#send a post request to set color
def setColor(r,g,b):
	data = {}
	data['red'] = r
	data['green'] = g
	data['blue'] = b
	json_data = json.dumps(data)
	r = requests.post(urlSetColor, data)
	print(r.status_code, r.reason)
import color_sensor
import json
import requests
from thread_class import perpetualTimer
import colorsys
import threading

urlSetColor = "http://localhost:8080/setcolor"

#sets a name for color in a range
def colorNameFromHsv(h,s,v):
	colorName = 'unknown'
	if (s > 30):
		if (h < 15 or h > 315):
			colorName = 'red'
		elif (45 > h > 25):
			colorName = 'yellow'
	elif (s < 30):
		if (250 > h > 170):
			colorName = 'blue'
		elif(125 > h > 85):
			colorName = 'green'
		elif(60 > h > 45):
			colorName = 'white'

	return colorName

def currentColor():
    r, g, b = color_sensor.sense_colors()
    # r = 1024
    # g = 0
    # b = 0

    h, s, v = colorsys.rgb_to_hsv(r/1024., g/1024., b/1024.)
    h = h * 360
    s = s * 100
    v = v * 100
    print('HSV')
    print(h,s,v)

    color = colorNameFromHsv(h,s,v)
    data = {}
    data['color'] = color
    json_data = json.dumps(data)

    return json_data

#read sensor
def getColor():
    json_color = currentColor()
    jdata = json.loads(json_color)

    if jdata['color'] != 'unknown':
    	print(jdata['color'])
    	setColor(jdata['color'])

    return json_color

#thread global variable
t = perpetualTimer(1,getColor)

#starts thread and returns first color read
def startReading():
	global t

	h = threading.Thread(target=t.run)
	t.start()
	h.start()

	return "Reading colors"

#stops the thread
def stopReading():
	global t
	t.stop()
	return "Not reading colors"

#send a post request to set color
def setColor(color):
	r = requests.post(urlSetColor, json={"color": color})
	print(r.status_code, r.reason)

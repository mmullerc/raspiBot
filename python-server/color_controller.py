from color_sensor import sense_colors
import json
import requests
from thread_class import perpetualTimer
import colorsys

urlSetColor = "http://localhost:8080/setcolor"

#read sensor
def getColor():
    r, g, b = sense_colors()
    print('RGB')
    print(r,g,b)
    h, s, v = colorsys.rgb_to_hsv(r/255., g/255., b/255.)

    h = h * 360
    s = s * 100
    v = v * 100
    print('HSV')
    print(h,s,v)

    color = colorNameFromHsv(h,s,v)
    data = {}
    data['color'] = color
    json_data = json.dumps(data)

    print(color)
    if color != 'unknown':
    	setColor(color)

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
	t = perpetualTimer(1,getColor)
	return "Not reading colors"

#send a post request to set color
def setColor(color):
	r = requests.post(urlSetColor, json={"color": color})
	print(r.status_code, r.reason)

#sets a name for color in a range
def colorNameFromHsv(h,s,v):
	colorName = ''
	if (h < 15 or h > 315):
		colorName = 'red'
	elif (250 > h > 170):
		colorName = 'blue'
	elif (125 > h > 80):
		colorName = 'green'
	elif (70 > h > 45):
		colorName = 'yellow'
	else:
		colorName = 'unknown'

	return colorName
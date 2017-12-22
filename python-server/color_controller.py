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
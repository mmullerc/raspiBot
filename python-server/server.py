import color_controller
from flask import Flask
import json

app = Flask(__name__)


@app.route("/")
def hello():
    return "Hello raspibot!"

@app.route("/startReading")
def startReading():
	res = color_controller.startReading()
	return res

@app.route("/stopReading")
def stopReading():
	res = color_controller.stopReading()
	return res

@app.route("/currentColor")
def currentColor():
	res = color_controller.currentColor()
	return res


if __name__ == "__main__":
    app.run()

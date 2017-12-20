from flask import Flask
import os
import color_sensor
app = Flask(__name__)


@app.route("/")
def hello():
    return "Hello raspibot!"

@app.route("/getColor")
def getColor():
    #path to python program
    return "green"


if __name__ == "__main__":
    app.run()

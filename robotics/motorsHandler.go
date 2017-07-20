package robotics

import (
	"net/http"
	"time"

	"gobot.io/x/gobot/platforms/raspi"
)

//PIN numbers
var STBY = "7" // GPIO-4

//Speed for motor A & B
var PWMA = "35" // GPIO19
var PWMB = "37" // GPIO26

//Motor A inputs
var AIN1 = "16" // GPIO-23
var AIN2 = "12" // GPIO-18

//Motor B inputs
var BIN1 = "32" // GPIO-12
var BIN2 = "22" // GPIO-25

var r = raspi.NewAdaptor()

func SetUpMotors(w http.ResponseWriter, req *http.Request) {
	//Entrypoint

	speed := byte(254)
	var direction = "forward"

	if direction == "forward" {
		moveForward("a", speed)
		moveForward("b", speed)
		time.Sleep(time.Second * 5)
		stop()
	}

	if direction == "backward" {
		moveBackward("a", speed)
		moveBackward("b", speed)
	}

	if direction == "right" {
		moveForward("a", speed)
		moveBackward("b", speed)
	}

	if direction == "left" {
		moveBackward("a", speed)
		moveForward("b", speed)
	}
}

func moveForward(motor string, speed byte) {
	move(speed)
	if motor == "a" {
		r.DigitalWrite(AIN1, 1)
		r.DigitalWrite(AIN2, 0)
	}

	if motor == "b" {
		r.DigitalWrite(BIN1, 1)
		r.DigitalWrite(BIN2, 0)
	}
}

func moveBackward(motor string, speed byte) {
	move(speed)
	if motor == "a" {
		r.DigitalWrite(AIN1, 0)
		r.DigitalWrite(AIN2, 1)
	}

	if motor == "b" {
		r.DigitalWrite(BIN1, 0)
		r.DigitalWrite(BIN2, 1)
	}
}

func move(speed byte) {
	r.DigitalWrite(STBY, 1)
	r.PwmWrite(PWMA, speed)
	r.PwmWrite(PWMB, speed)
}

func stop() {
	r.DigitalWrite(STBY, 0)
}

package robotics

import (
	"net/http"
	"time"

	"gobot.io/x/gobot/platforms/firmata"
)

//PIN numbers
var STBY = "7"

//Speed for motor A & B
var PWMA = "19"
var PWMB = "21"

//Motor A inputs
var AIN1 = "16"
var AIN2 = "12"

//Motor B inputs
var BIN1 = "32"
var BIN2 = "22"

var adaptor = firmata.NewAdaptor()

func SetUpMotors(w http.ResponseWriter, req *http.Request) {
	//Entrypoint
	//adaptor := firmata.NewAdaptor()
	speed := byte(128)
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
		adaptor.DigitalWrite(AIN1, 1)
		adaptor.DigitalWrite(AIN2, 0)
	}

	if motor == "b" {
		adaptor.DigitalWrite(BIN1, 1)
		adaptor.DigitalWrite(BIN2, 0)
	}
}

func moveBackward(motor string, speed byte) {
	move(speed)
	if motor == "a" {
		adaptor.DigitalWrite(AIN1, 0)
		adaptor.DigitalWrite(AIN2, 1)
	}

	if motor == "b" {
		adaptor.DigitalWrite(BIN1, 0)
		adaptor.DigitalWrite(BIN2, 1)
	}
}

func move(speed byte) {
	adaptor.DigitalWrite(STBY, 1)
	adaptor.PwmWrite(PWMA, speed)
	adaptor.PwmWrite(PWMB, speed)
}

func stop() {
	adaptor.DigitalWrite(STBY, 0)
}

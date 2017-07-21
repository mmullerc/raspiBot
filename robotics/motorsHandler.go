package robotics

import (
	"net/http"
	"time"
	"fmt"
	"raspibot/db"

	"gobot.io/x/gobot/platforms/raspi"
)

const (
	STBY = "7" //GPIO-4
	PWMA = "33" //GPIO-12
	PWMB = "12" //GPIO-18
	AIN1 = "38" //GPIO-19
	AIN2 = "40" //GPIO-12
	BIN1 = "35" //GPIO-16
	BIN2 = "37" //GPIO-20
)

var r = raspi.NewAdaptor()

func SetUpMotors(w http.ResponseWriter, req *http.Request) {
	//Entrypoint
	motor := db.GetStateByComponent(w, "motor")
	fmt.Fprint(w, motor.Name, ": ", motor.State, ",", motor.Direction, ",", motor.Speed, "\n")

	speed := byte(255)
	if motor.Speed == "slow" {
		speed = byte(128)
	}

	if motor.Direction == "forward" {
		moveForward("a", speed)
		moveForward("b", speed)
		time.Sleep(time.Second * 5)
		stop()
	}

	if motor.Direction == "backward" {
		moveBackward("a", speed)
		moveBackward("b", speed)
		time.Sleep(time.Second * 5)
		stop()
	}

	if motor.Direction == "right" {
		moveForward("a", speed)
		moveBackward("b", speed)
		time.Sleep(time.Second * 1)
		stop()
	}

	if motor.Direction == "left" {
		moveBackward("a", speed)
		moveForward("b", speed)
		time.Sleep(time.Second * 1)
		stop()
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

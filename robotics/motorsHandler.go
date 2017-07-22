package robotics

import (
	"fmt"
	"net/http"
	"raspibot/db"
	"time"

	"gobot.io/x/gobot/platforms/raspi"
)

const (
	//PINS GROUP
	STBY = "7"  //GPIO-4
	PWMA = "33" //GPIO-12
	PWMB = "12" //GPIO-18
	AIN1 = "38" //GPIO-19
	AIN2 = "40" //GPIO-12
	BIN1 = "35" //GPIO-16
	BIN2 = "37" //GPIO-20

	//MOTORS GROUP
	MOTOR_LEFT  = "a"
	MOTOR_RIGHT = "b"

	//TURN ANGLES
	angle45  = 500
	angle90  = 700
	angle180 = 1300
)

var (
	r     = raspi.NewAdaptor()
	speed = byte(255)
	motor = db.GetStateByComponent("motor")
)

func SetUpMotors(w http.ResponseWriter, req *http.Request) {
	//Entrypoint
	fmt.Fprint(w, motor.Name, ": ", motor.State, ",", motor.Direction, ",", motor.Speed, "\n")

	if motor.Speed == "slow" {
		speed = byte(128)
	}

	if motor.Direction == "forward" {
		moveForward(MOTOR_LEFT, speed)
		moveForward(MOTOR_RIGHT, speed)
		time.Sleep(time.Second * 5)
		stop()
	}

	if motor.Direction == "backward" {
		moveBackward(MOTOR_LEFT, speed)
		moveBackward(MOTOR_RIGHT, speed)
		time.Sleep(time.Second * 5)
		stop()
	}

	if motor.Direction == "right" {
		moveForward(MOTOR_LEFT, speed)
		moveBackward(MOTOR_RIGHT, speed)
		time.Sleep(time.Second * 1)
		stop()
	}

	if motor.Direction == "left" {
		moveBackward(MOTOR_LEFT, speed)
		moveForward(MOTOR_RIGHT, speed)
		time.Sleep(time.Second * 1)
		stop()
	}
}

func moveForward(motor string, speed byte) {
	move(speed)
	if motor == MOTOR_LEFT {
		r.DigitalWrite(AIN1, 1)
		r.DigitalWrite(AIN2, 0)
	}

	if motor == MOTOR_RIGHT {
		r.DigitalWrite(BIN1, 1)
		r.DigitalWrite(BIN2, 0)
	}
}

func moveBackward(motor string, speed byte) {
	move(speed)
	if motor == MOTOR_LEFT {
		r.DigitalWrite(AIN1, 0)
		r.DigitalWrite(AIN2, 1)
	}

	if motor == MOTOR_RIGHT {
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
	time.Sleep(time.Second * 1)
}

func turn(turns int) {
	if turns == 1 {
		turnRight(angle90)
		fmt.Printf("Angle 1: %v\n", angle90)
	} else if turns == 2 {
		turnLeft(angle180)
		fmt.Printf("Angle 2: %v\n", angle180)
	} else {
		turnLeft(angle90)
		fmt.Printf("Angle 180: %v\n", angle180)
	}
}

func turnRight(timeToRun time.Duration) {
	moveForward(MOTOR_LEFT, speed)
	moveBackward(MOTOR_RIGHT, speed)
	time.Sleep(time.Millisecond * timeToRun)
}

func turnLeft(timeToRun time.Duration) {
	moveForward(MOTOR_RIGHT, speed)
	moveBackward(MOTOR_LEFT, speed)
	time.Sleep(time.Millisecond * timeToRun)
}

func moveAllMotorsForward() {
	moveForward(MOTOR_LEFT, speed)
	moveForward(MOTOR_RIGHT, speed)
}

func KillMotors(w http.ResponseWriter, req *http.Request) {
	r.DigitalWrite(AIN1, 0)
	r.DigitalWrite(AIN2, 0)
	r.DigitalWrite(BIN1, 0)
	r.DigitalWrite(BIN2, 0)
}

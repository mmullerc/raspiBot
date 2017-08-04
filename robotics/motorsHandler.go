package robotics

import (
	"gobot.io/x/gobot/platforms/raspi"
	"time"
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

	//TURN ANGLES
	// angle45  = 500
	// angle90  = 700
	// angle180 = 1300
)

var r raspi.Adaptor

func StartMotors(speed byte, adaptor *raspi.Adaptor) {

	//board := raspi.NewAdaptor()
	//board.DigitalWrite(STBY, 1)
	//r = adaptor

	adaptor.DigitalWrite(STBY, 1)
	adaptor.PwmWrite(PWMA, speed)
	adaptor.PwmWrite(PWMB, speed)

	MoveForward()
}

func MoveForward() {
	moveMotors(1, 0, 1, 0)
}

func MoveBackward() {
	moveMotors(0, 1, 0, 1)
}

func MoveRight() {
	moveMotors(1, 0, 0, 1)
}

func MoveLeft() {
	moveMotors(0, 1, 1, 0)
}

func turn(direction byte, angle time.Duration) {
	if direction == 0 {
		MoveRight()
	} else {
		MoveLeft();
	}
	time.Sleep(time.Millisecond * angle)
	MoveForward()
}

func KillMotors() {
	moveMotors(0, 0, 0, 0)
	r.DigitalWrite(STBY, 0)
}

func moveMotors(ain1 byte, ain2 byte, bin1 byte, bin2 byte) {

	// LEFT MOTOR
	r.DigitalWrite(AIN1, ain1)
	r.DigitalWrite(AIN2, ain2)

	// RIGHT MOTOR
	r.DigitalWrite(BIN1, bin1)
	r.DigitalWrite(BIN2, bin2)
}
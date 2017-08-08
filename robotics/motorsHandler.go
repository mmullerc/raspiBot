package robotics

import (
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

	//TURN ANGLES
	angle45  = 500
	angle90  = 700
	angle180 = 1300
)

var motorsAdaptor *raspi.Adaptor = new(raspi.Adaptor)

func StartMotors(speed byte, adaptor *raspi.Adaptor) {

	motorsAdaptor = adaptor
	//board := raspi.NewAdaptor()
	//board.DigitalWrite(STBY, 1)
	//r = adaptor

	motorsAdaptor.DigitalWrite(STBY, 1)
	motorsAdaptor.PwmWrite(PWMA, speed)
	motorsAdaptor.PwmWrite(PWMB, speed)

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
		MoveLeft()
	}

	time.Sleep(time.Millisecond * angle)
	distance := GetDistance(raspiAdaptor)

	if distance < 50 {
		turn(direction, angle)
	}

	MoveForward()
}

func KillMotors() {
	moveMotors(0, 0, 0, 0)
	motorsAdaptor.DigitalWrite(STBY, 0)
}

func moveMotors(ain1 byte, ain2 byte, bin1 byte, bin2 byte) {

	// LEFT MOTOR
	motorsAdaptor.DigitalWrite(AIN1, ain1)
	motorsAdaptor.DigitalWrite(AIN2, ain2)

	// RIGHT MOTOR
	motorsAdaptor.DigitalWrite(BIN1, bin1)
	motorsAdaptor.DigitalWrite(BIN2, bin2)
}

func MoveCar() {
	distance := GetDistance(raspiAdaptor)
	println(distance)
	if distance < 50 {
		turn(0, angle45) //por ahora, solo gira a la derecha
	}
}

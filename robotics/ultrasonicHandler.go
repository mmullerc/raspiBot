package robotics

import (
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
	"time"
)

const (
	//PINS GROUP
	trig = "29" //GPIO5
	echo = "31" //GPIO6
)

func GetDistance() float64 {
	r := raspi.NewAdaptor()

	trigPin := gpio.NewDirectPinDriver(r, trig)
	echoPin := gpio.NewDirectPinDriver(r, echo)

	trigPin.DigitalWrite(byte(0))
	time.Sleep(2 * time.Microsecond)

	trigPin.DigitalWrite(byte(1))
	time.Sleep(10 * time.Microsecond)

	trigPin.DigitalWrite(byte(0))
	start := time.Now()
	end := time.Now()

	for {
		val, err := echoPin.DigitalRead()
		start = time.Now()
		if err != nil {
			println(err)
			break
		}
		if val == 0 {
			continue
		}
		break
	}

	for {
		val, err := echoPin.DigitalRead()
		end = time.Now()
		if err != nil {
			println(err)
			break
		}
		if val == 1 {
			continue
		}
		break
	}

	duration := end.Sub(start)
	distance := duration.Seconds() * 34300
	distance = distance / 2 //one way travel time

	return distance
}
package robotics

import (
	"fmt"
	"net/http"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func Blinking(w http.ResponseWriter, req *http.Request) {

	fmt.Fprint(w, "Starting LED!\n")
	r := raspi.NewAdaptor()
	led := gpio.NewLedDriver(r, "7")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{led},
		work,
	)
	robot.Start()
}

func Start(w http.ResponseWriter, req *http.Request) {

	fmt.Fprint(w, "Starting Motor!\n")
	r := raspi.NewAdaptor()
	motor := gpio.NewLedDriver(r, "7")

	startMotor := func() {
		motor.On()
	}

	robot := gobot.NewRobot("ServoMotors",
		[]gobot.Connection{r},
		[]gobot.Device{motor},
		startMotor,
	)
	robot.Start()
}

func Stop(w http.ResponseWriter, req *http.Request) {

	fmt.Fprint(w, "Stoping Motor!\n")
	r := raspi.NewAdaptor()
	motor := gpio.NewLedDriver(r, "7")

	stopMotor := func() {
		motor.Off()
	}

	robot := gobot.NewRobot("ServoMotors",
		[]gobot.Connection{r},
		[]gobot.Device{motor},
		stopMotor,
	)
	robot.Start()
}

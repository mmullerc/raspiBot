package robotics

import (
	"net/http"
	"raspibot/db"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

var (
	turns = 0
)

func UltrasonicSensor(w http.ResponseWriter, req *http.Request) {
	r := raspi.NewAdaptor()

	trigPin := gpio.NewDirectPinDriver(r, "29") // GPIO5
	echoPin := gpio.NewDirectPinDriver(r, "31") // GPIO6
	led := gpio.NewLedDriver(r, "36")           // GPIO16

	moveAllMotorsForward()

	work := func() {

		gobot.Every(100*time.Millisecond, func() {

			motor = db.GetStateByComponent("motor")

			if motor.State == "on" {
				println("Starting probing ")

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

				if distance < 50 {
					turns = turns + 1
					led.On()
					turn(turns)
				} else {
					led.Off()
					turns = 0
					moveAllMotorsForward()
				}
			} else {
				stop()
			}
		})
	}

	robot := gobot.NewRobot("makeyBot",
		[]gobot.Connection{r},
		[]gobot.Device{trigPin, echoPin, led},
		work,
	)

	robot.Start()
}

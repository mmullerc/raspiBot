package ultrasonic

import (
"net/http"
"gobot.io/x/gobot/platforms/raspi"
"gobot.io/x/gobot/drivers/gpio"
"gobot.io/x/gobot"
"time"
"fmt"
)

func UltrasonicSensor(w http.ResponseWriter, req *http.Request) {
	r := raspi.NewAdaptor()

	trigPin := gpio.NewDirectPinDriver(r, "29") // GPIO5
	echoPin := gpio.NewDirectPinDriver(r, "31") // GPIO6
	led := gpio.NewLedDriver(r, "36") // GPIO21

	work := func() {

	gobot.Every(250 * time.Millisecond, func() {

		println("Starting probing ")
		
		trigPin.DigitalWrite(byte(0))
		time.Sleep(2 * time.Microsecond)

		trigPin.DigitalWrite(byte(1))
		time.Sleep(10 * time.Microsecond)

		trigPin.DigitalWrite(byte(0))
		start := time.Now()
		end := time.Now()

		for {
			val, err := echoPin.DigitalRead();
			start = time.Now()
			if (err != nil) {
				println(err)
				break
			}
			if val == 0 { continue }
			break
		}

		for {
			val, err := echoPin.DigitalRead();
			end = time.Now()
			if (err != nil) {
				println(err)
				break
			}
			if val == 1 {
				continue
			}
			break
		}

		duration := end.Sub(start)
		durationAsInt64 := int64(duration)
		distance := duration.Seconds() * 34300
		distance = distance / 2 //one way travel time
		
		if distance < 20 {
			led.On()
		}else {
			led.Off()
		}
		fmt.Printf("Duration : %v %v %v \n", distance, duration.Seconds(), durationAsInt64)
	})
}

robot := gobot.NewRobot("makeyBot",
	[]gobot.Connection{r},
	[]gobot.Device{trigPin, echoPin, led},
	work,
)

robot.Start()
}

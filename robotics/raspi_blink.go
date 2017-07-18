package robotics

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"raspibot/logger"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

var STOP = 0
var START = 1

func TurnLedsOn(w http.ResponseWriter, req *http.Request) {

	fmt.Fprint(w, "Starting Leds on Multiple Pins!\n")
	r := raspi.NewAdaptor()
	ledPin7 := gpio.NewLedDriver(r, "7")   // GPIO-4
	ledPin20 := gpio.NewLedDriver(r, "19") // SPIMOSI
	ledPin21 := gpio.NewLedDriver(r, "21") // SPIMISO
	ledPin16 := gpio.NewLedDriver(r, "16") // GPIO-23

	ledPin12 := gpio.NewLedDriver(r, "12") // GPIO-18
	ledPin32 := gpio.NewLedDriver(r, "32") // GPIO-12
	ledPin22 := gpio.NewLedDriver(r, "22") // GPIO-25

	blinkingTest := func() {
		ledPin7.On()
		ledPin20.On()
		ledPin21.On()
		ledPin16.On()
		ledPin12.On()
		ledPin32.On()
		ledPin22.On()
	}
	blinkBot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{ledPin7},
		[]gobot.Device{ledPin20},
		[]gobot.Device{ledPin21},
		[]gobot.Device{ledPin16},
		[]gobot.Device{ledPin12},
		[]gobot.Device{ledPin32},
		[]gobot.Device{ledPin22},
		blinkingTest,
	)
	blinkBot.Start()
}

func TurnLedsOff(w http.ResponseWriter, req *http.Request) {

	fmt.Fprint(w, "Starting Leds on Multiple Pins!\n")
	r := raspi.NewAdaptor()
	ledPin7 := gpio.NewLedDriver(r, "7")
	ledPin20 := gpio.NewLedDriver(r, "19")
	ledPin21 := gpio.NewLedDriver(r, "21")
	ledPin16 := gpio.NewLedDriver(r, "16")

	ledPin12 := gpio.NewLedDriver(r, "12")
	ledPin32 := gpio.NewLedDriver(r, "32")
	ledPin22 := gpio.NewLedDriver(r, "22")

	blinkingTest := func() {
		ledPin7.Off()
		ledPin20.Off()
		ledPin21.Off()
		ledPin16.Off()
		ledPin12.Off()
		ledPin32.Off()
		ledPin22.Off()
	}
	blinkBot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{ledPin7},
		[]gobot.Device{ledPin20},
		[]gobot.Device{ledPin21},
		[]gobot.Device{ledPin16},
		[]gobot.Device{ledPin12},
		[]gobot.Device{ledPin32},
		[]gobot.Device{ledPin22},
		blinkingTest,
	)
	blinkBot.Start()
}

func Start(w http.ResponseWriter, req *http.Request) {
	logger.Print("Starting Motor!", w)
	manageMotor(START)
}

func Stop(w http.ResponseWriter, req *http.Request) {
	logger.Print("Stoping Motor!", w)
	manageMotor(STOP)
}

func DriveForward(w http.ResponseWriter, req *http.Request) {
	logger.Print("Drive Forward!", w)
	t, err := strconv.ParseInt(req.URL.Query().Get("time"), 10, 64)

	if err != nil {
		t = int64(rand.Intn(5))
	}

	logger.Print(fmt.Sprintf("%s%d", "Time: ", t), w)

	for a := 0; a < 10; a++ {
		time.Sleep(2 * time.Second)
		logger.Print(fmt.Sprintf("%s%d", "Lap #", a), w)
		go manageMotor(START)
		time.Sleep(time.Duration(rand.Int31n(int32(t))) * time.Second)
		go manageMotor(STOP)
	}
}

func manageMotor(action int) {
	logger.Print("Manage motor function", nil)
	r := raspi.NewAdaptor()
	motor := gpio.NewLedDriver(r, "7")

	actionMotor := func() {
		if action == STOP {
			motor.Off()
		} else if action == START {
			motor.On()
		}
	}

	robot := gobot.NewRobot("ServoMotors",
		[]gobot.Connection{r},
		[]gobot.Device{motor},
		actionMotor,
	)
	robot.Start()
}

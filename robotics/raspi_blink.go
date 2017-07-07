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

var STOP 			= 0
var START 			= 1

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
	logger.Print("Starting Motor!", w)
	manageMotor(START);
}

func Stop(w http.ResponseWriter, req *http.Request) {
	logger.Print("Stoping Motor!", w)
	manageMotor(STOP);
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

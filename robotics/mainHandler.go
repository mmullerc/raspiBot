package robotics

import (
	"fmt"
	"net/http"
	"raspibot/db"
	"raspibot/utilities"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/raspi"
)

var adaptor = raspi.NewAdaptor()
var robot *gobot.Robot
var ticker *time.Ticker

func TurnOnCar(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Starting car!\n")
	//StartMotors(byte(255), adaptor)
	db.StartCar()
	fmt.Printf("%+v\n", robot)
	if robot != nil {
		if !robot.Running() {
			robot = gobot.NewRobot("raspiBot",
				[]gobot.Connection{adaptor},
				StartListening,
			)
			robot.Start()
		}
	}
}

func StopCar(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Stoping car!\n")
	db.StopCar()
}

func Move(w http.ResponseWriter, req *http.Request) {
	StartMotors(byte(255), adaptor)
}

func StartListening() {
	ticker = gobot.Every(100*time.Millisecond, func() {

		motor, err := db.GetStateByComponent("motor")
		utilities.CheckForStringErr(err)

		//Check motor
		if motor.State == "on" {
			println("Starting motors!")

			distance := GetDistance()

			if distance < 50 {
				MoveRight()
			} else {
				MoveForward()
			}
		}

		//Check general car status
		car, err := db.GetStateByComponent("car")
		utilities.CheckForStringErr(err)
		if car.State == "off" {
			stopListening()
		}
	})
}

func stopListening() {
	ticker.Stop()
}

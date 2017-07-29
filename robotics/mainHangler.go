package robotics

import (
	"gobot.io/x/gobot"
	"net/http"
	"raspibot/db"
    "time" 
    "gobot.io/x/gobot/platforms/raspi"
)

var adaptor = raspi.NewAdaptor()
var robot gobot.Robot
var ticker *Ticker

func TurnOnCar(w http.ResponseWriter, req *http.Request) {
	//StartMotors(byte(255), adaptor)
    db.StartCar()
    if (robot == nil) {
        robot = gobot.NewRobot("makeyBot",
            []gobot.Connection{adaptor},
            StartListening,
        )
	    robot.Start()
    }


}

func StopCar(w http.ResponseWriter, req *http.Request) {
    db.StopCar()
}

func Move(w http.ResponseWriter, req *http.Request) {
    StartMotors(byte(255), adaptor)
}

func StartListening() {
    ticker = gobot.Every(100*time.Millisecond, func() {

        motor := db.GetStateByComponent("motor")

        //Check motor
        if motor.State == "on" {
            println("Starting probing ")

            distance := GetDistance()

            if distance < 50 {
                MoveRight()
            } else {
                MoveForward()
            }
        }

        //Check general car status
        car := db.GetStateByComponent("car")
        if car.State == "off" {
            stopListening()
        }
    })
}

func stopListening() {
    ticker.Stop()
}
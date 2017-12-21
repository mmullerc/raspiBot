package robotics

import (
	"fmt"
	"net/http"
	"raspibot/db"
	"raspibot/utilities"
	"time"
	"encoding/json"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/raspi"
)

var raspiAdaptor = raspi.NewAdaptor()
var robot *gobot.Robot
var ticker *time.Ticker

type User struct {
	User string
}

func initRobot() {
	robot = gobot.NewRobot("raspiBot",
		[]gobot.Connection{raspiAdaptor},
		StartListening,
	)
}

func TurnOnCar(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Starting car!\n")
	fmt.Printf("%+v\n", robot)

	db.StartCar()
	if robot != nil {
		if !robot.Running() {
			robot.Start()
		}
	} else {
		initRobot()
		TurnOnCar(w, req)
	}
}

func StopCar(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Stoping car!\n")
	db.StopCar()
}

func Move(w http.ResponseWriter, req *http.Request) {
	StartMotors(byte(255), raspiAdaptor)
}

func Navegate(w http.ResponseWriter, req *http.Request) {

	var u User;
	var initColor Color;

    if req.Body == nil {
        http.Error(w, "Please send a request body", 400)
        return
    }
    err := json.NewDecoder(req.Body).Decode(&u)
    if err != nil {
        http.Error(w, err.Error(), 400)
        return
    }
    
    color := db.FindLocation(u.User)
    fmt.Printf(color)

	response, err2 := http.Get("http://localhost:5000/startReading")
    if err2 != nil {
        fmt.Printf("%s", err2)
    } else {
        defer response.Body.Close()
        errInitColor := json.NewDecoder(response.Body).Decode(&initColor)
        if errInitColor != nil {
            fmt.Printf("%s", errInitColor)
        }
    }

	fmt.Printf("%v\n", initColor)

	if initColor.Color != color {
		output := fmt.Sprintf("%v%s%v%s%v", u.User," is in the ",color, " table. The car location is in", initColor)

		StartMotors(byte(255), raspiAdaptor)
		fmt.Printf(output)
    	fmt.Fprint(w, output)
	} else {
		output := fmt.Sprintf("%s%v%s","The car is in the ",u.User,"'s table already")
		fmt.Printf(output)
    	fmt.Fprint(w, output)
	}
    

}

func StartListening() {
	ticker = gobot.Every(100*time.Millisecond, func() {

		motor, err := db.GetStateByComponent("motor")
		utilities.CheckForStringErr(err)

		//Check motor
		if motor.State == "on" {
			println("Starting motors!")
			MoveCar()
		}

		//Check general car status
		car, err := db.GetStateByComponent("car")
		utilities.CheckForStringErr(err)
		if car.State == "off" {
			robot.Stop()
		}
	})
}

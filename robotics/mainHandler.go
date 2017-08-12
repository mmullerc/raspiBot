package robotics

import (
	"fmt"
	"net/http"
	"raspibot/db"
	"raspibot/system"
	"raspibot/utilities"
	"regexp"
	"strconv"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/raspi"
)

type Signal struct {
	macAddress string
	intensity  int64
}

var raspiAdaptor = raspi.NewAdaptor()
var robot *gobot.Robot
var ticker *time.Ticker

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

func CheckConnectionStrength(w http.ResponseWriter, req *http.Request) []Signal {

	networkObjs := []Signal{}
	c := make(chan string)
	go system.CheckNetworkStrngth(c)
	result := <-c
	println(result)

	//var essidExists *regexp
	// re := regexp.MustCompile("KG-Network")
	// networkName := re.FindString(result)
	// println(re)
	// println(networkName)
	// if len(networkName) > 0 {

	re := regexp.MustCompile("/([0-9A-F]{2}[:-]){5}([0-9A-F]{2})/g")
	accessPoints := re.FindAllString(result, -1)
	re = regexp.MustCompile(`([-][0-9])\w+`)
	intensityNumbers := re.FindAllString(result, -1)

	for index, accessPoint := range accessPoints {
		print(accessPoint)
		print(intensityNumbers)
		print(index)

		if db.RouterExists(accessPoint) {
			intensityNumber, err := strconv.ParseInt(intensityNumbers[index], 10, 64)
			if err != nil {
				panic(err)
			}
			addrObject := Signal{macAddress: accessPoint, intensity: intensityNumber}
			networkObjs = append(networkObjs, addrObject)
		}
	}
	return networkObjs
}

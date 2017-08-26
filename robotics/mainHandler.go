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

		result := CheckConnectionStrength()

		for index := 0; index < len(result); index++ {
			println("here")
			fmt.Printf("%+v\n", result[index])
		}

		// println(result)
		// fmt.Printf("%+v\n", result)

		//Check general car status
		car, err := db.GetStateByComponent("car")
		utilities.CheckForStringErr(err)
		if car.State == "off" {
			robot.Stop()
		}
	})
}

func CheckConnection(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "checking!\n")
	result := CheckConnectionStrength()
	println(result)
	fmt.Printf("%+v\n", result)
	for index := 0; index < len(result); index++ {
		fmt.Printf("%+v\n", result[index])
	}
}

func CheckConnectionStrength() []Signal {
	networkObjs := []Signal{}
	c := make(chan string)
	go system.CheckNetworkStrngth(c)
	result := <-c

	re := regexp.MustCompile(`([0-9A-F]{2}[:-]){5}([0-9A-F]{2})`)
	accessPoints := re.FindAllString(result, -1)

	re = regexp.MustCompile(`([-][0-9])\w+`)
	intensityNumbers := re.FindAllString(result, -1)

	for index, accessPoint := range accessPoints {
		accesPointExists, err := db.RouterExists(accessPoint)
		println(err)

		if accesPointExists {
			intensityNumber, err := strconv.ParseInt(intensityNumbers[index], 10, 64)
			if err != nil {
				panic(err)
			}
			addrObject := Signal{macAddress: accessPoint, intensity: intensityNumber}
			println("creating obj")
			networkObjs = append(networkObjs, addrObject)
		}
	}
	return networkObjs
}

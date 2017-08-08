package routes

import (
	"net/http"
	"raspibot/controllers"
	"raspibot/robotics"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type AllRoutes []Route

var Routes = AllRoutes{
	Route{
		"Index",
		"GET",
		"/",
		controllers.Index,
	},
	Route{
		"startLeds",
		"GET",
		"/turnLedsOn",
		robotics.TurnLedsOn,
	},
	Route{
		"stopLeds",
		"GET",
		"/turnLedsOff",
		robotics.TurnLedsOff,
	},
	Route{
		"startMotor",
		"GET",
		"/startMotor",
		robotics.Start,
	},
	Route{
		"stopMotor",
		"GET",
		"/stopMotor",
		robotics.Stop,
	},
	Route{
		"driveForward",
		"GET",
		"/driveForward",
		robotics.DriveForward,
	},
	Route{
		"setUpMotors",
		"GET",
		"/setUpMotors",
		robotics.SetUpMotors,
	},
	Route{
		"setUpUltrasonic",
		"GET",
		"/startUltrasonic",
		robotics.UltrasonicSensor,
	},
	Route{
		"killMotors",
		"GET",
		"/killMotors",
		robotics.KillMotors,
	},
}

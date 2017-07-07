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
		"startLed",
		"GET",
		"/startLed",
		robotics.Blinking,
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
}

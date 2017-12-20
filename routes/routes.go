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
		"TurnOnCar",
		"GET",
		"/startcar",
		robotics.TurnOnCar,
	},
	Route{
		"StopCar",
		"GET",
		"/stopcar",
		robotics.StopCar,
	},
	Route{
		"Move",
		"GET",
		"/move",
		robotics.Move,
	},
	Route{
		"SetColor",
		"POST",
		"/setcolor",
		robotics.SetColor,
	},
}

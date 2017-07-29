package robotics

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/raspi"
)

var instance *gobot.Robot

func GetRobot(adaptor *raspi.Adaptor) *gobot.Robot {
	if instance == nil {
		botSingleton(adaptor)
	}
	return instance
}

func botSingleton(adaptor *raspi.Adaptor) {
	instance = gobot.NewRobot("makeyBot",
				[]gobot.Connection{adaptor},
				Work,
			)
}
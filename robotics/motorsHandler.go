package robotics

import (
	"net/http"
	"strconv"

	"gobot.io/x/gobot/platforms/firmata"
	"gobot.io/x/gobot/platforms/firmata/client"
)

//type Adapdtor raspi.Adaptor

type firmataAdaptor firmata.Adaptor

func SetUpMotors(w http.ResponseWriter, req *http.Request) {
	//Entrypoint
}

// PwmWrite writes the 0-254 value to the specified pin
func (f *firmataAdaptor) PwmWrite(pin string, level byte) (err error) {
	p, err := strconv.Atoi(pin)
	if err != nil {
		return err
	}

	if f.Board.Pins()[p].Mode != client.Pwm {
		err = f.Board.SetPinMode(p, client.Pwm)
		if err != nil {
			return err
		}
	}
	err = f.Board.AnalogWrite(p, int(level))
	return
}

// DigitalWrite writes a value to the pin. Acceptable values are 1 or 0.
func (f *firmataAdaptor) DigitalWrite(pin string, level byte) (err error) {
	p, err := strconv.Atoi(pin)
	if err != nil {
		return
	}

	if f.Board.Pins()[p].Mode != client.Output {
		err = f.Board.SetPinMode(p, client.Output)
		if err != nil {
			return
		}
	}

	err = f.Board.DigitalWrite(p, int(level))
	return
}

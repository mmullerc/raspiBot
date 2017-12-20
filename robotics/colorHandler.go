package robotics

import (
	"fmt"
	"net/http"
	// "encoding/json"
	 //"io"
	 "io/ioutil"
	// "raspibot/db"
	// "raspibot/utilities"
	// "time"

	// "gobot.io/x/gobot"
	// "gobot.io/x/gobot/platforms/raspi"
)

type Color struct {
	color string
}

func SetColor(w http.ResponseWriter, req *http.Request) {
	
	body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		results := string(body)

		fmt.Fprint(w, results)

    fmt.Printf("%+v\n", results)
}
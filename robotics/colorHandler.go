package robotics

import (
	"fmt"
	"net/http"
	// "encoding/json"
	// "io"
	// "io/ioutil"
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
		results = append(results, string(body))

		fmt.Fprint(w, "POST done")

    fmt.Printf("%+v\n", body)
    //t := RepoCreateTodo(color)
    // w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    // w.WriteHeader(http.StatusCreated)
    // if err := json.NewEncoder(w).Encode(t); err != nil {
    //     panic(err)
    // }
}
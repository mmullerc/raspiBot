package robotics

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Color struct {
	Color string
	Id int
}

func SetColor(w http.ResponseWriter, r *http.Request) {
	
	var u Color

    if r.Body == nil {
        http.Error(w, "Please send a request body", 400)
        return
    }
    err := json.NewDecoder(r.Body).Decode(&u)
    if err != nil {
        http.Error(w, err.Error(), 400)
        return
    }
    output := fmt.Sprintf("%s%v", "Recieved color: ", u.Color)
    fmt.Printf(output)
    fmt.Fprint(w, output)

}
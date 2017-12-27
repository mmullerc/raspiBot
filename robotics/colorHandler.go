package robotics

import (
    "encoding/json"
    "fmt"
    "net/http"
    "raspibot/db"
)

type Color struct {
    Color string
}

func SetColor(w http.ResponseWriter, r *http.Request) {

    var u Color

    if r.Body == nil {
        fmt.Printf("Please send a request body")
        http.Error(w, "Please send a request body", 400)
        return
    }
    err := json.NewDecoder(r.Body).Decode(&u)
    if err != nil {
        fmt.Printf("%v", err)
        http.Error(w, err.Error(), 400)
        return
    }
    output := fmt.Sprintf("%s%v", "Recieved color: ", u.Color)
    fmt.Printf(output)
    fmt.Fprint(w, output)

    if u.Color == Direction {
        response, err2 := http.Get("http://localhost:5000/stopReading")
        if err2 != nil {
            fmt.Printf("%s", err2)
        } else {
            defer response.Body.Close()
        }

        db.StopCar()
    }

}

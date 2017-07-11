package controllers

import (
	"fmt"
	"net/http"
	"raspibot/db"
)

func Index(w http.ResponseWriter, r *http.Request) {
	//Testing DB
	//db.InsertState("motor", "on", "left", "fast")
	db.GetStateByComponent(w, "motor")

	fmt.Fprint(w, "Welcome!\n")
}

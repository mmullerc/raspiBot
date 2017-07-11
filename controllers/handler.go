package controllers

import (
	"fmt"
	"net/http"
	"raspibot/db"
)

func Index(w http.ResponseWriter, r *http.Request) {
	//Testing DB
	db.InsertState("motors", "on", "left", "fast")
	db.GetStateByComponent(w, "motors")

	fmt.Fprint(w, "Welcome!\n")
}

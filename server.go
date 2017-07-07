package main

import (
	"log"
	"net/http"
	"raspberrypi/router"
)

func main() {
	router := router.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}

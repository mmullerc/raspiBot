package main

import (
	"log"
	"net/http"
	"raspibot/router"
)

func main() {
	router := router.NewRouter()
	println("Starting Server on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

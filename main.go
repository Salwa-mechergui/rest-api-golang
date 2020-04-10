package main

import (
	"BookingAPI/Book/Helpers"
	"BookingAPI/Book/Router"
	"log"
	"net/http"
)

func main() {
	// Init router
	Helpers.Migration()
	// config.Connect()
	// Start server
	r := Router.SetupRouter()
	log.Fatal(http.ListenAndServe(":8000", r))
}

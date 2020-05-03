package main

import (
	"fmt"
	"github/reccrides/Helpers"
	"github/reccrides/Router"
	_ "github/reccrides/docs"
	"log"
	"net/http"
	"os"

	"github.com/subosito/gotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

//init env
func Init() {
	gotenv.Load()
}

// @title rideCreation
// @version 1.0
// @description This is a sample service for managing booking requests
// @termsOfService http://swagger.io/terms/
// @contact.name Yuso
// @contact.email salwa@craftfoundry.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
func main() {
	fmt.Println("Starting the application...")
	// Init router
	Init()
	Helpers.Migration()
	// config.Connect()
	// Start server
	r := Router.SetupRouter()
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
		log.Println("[-] No PORT environment variable detected. Setting to ", port)
	}
	// Listen and Serve in 0.0.0.0:8000
	log.Fatal(http.ListenAndServe(":"+port, r))
	fmt.Println("Terminating the application...")
}

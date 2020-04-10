package Router

import (
	"BookingAPI/Book/controllers"

	"github.com/gorilla/mux"
)

//router
func SetupRouter() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/agent/{Id}", controllers.GetEndpoint).Methods("GET")
	router.HandleFunc("/agent", controllers.GetAgentEndpoint).Methods("GET")
	return router
}

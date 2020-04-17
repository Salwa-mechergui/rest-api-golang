package Router

import (
	"github/reccrides/controllers"

	"github.com/gorilla/mux"
)

//router
func SetupRouter() *mux.Router {

	router := mux.NewRouter()
	// routes
	router.HandleFunc("/login", controllers.CreateEndpoint).Methods("POST")
	router.HandleFunc("/agent/{id}", controllers.GetByIdEndpoint).Methods("GET")
	router.HandleFunc("/agent", controllers.GetAgentEndpoint).Methods("GET")
	return router
}

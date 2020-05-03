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
	router.HandleFunc("/frontUser", controllers.GetAllFrontUsersEndpoint).Methods("GET")
	router.HandleFunc("/frontUser/{id}", controllers.GetFrontUsersByIdEndpoint).Methods("GET")
	router.HandleFunc("/passenger", controllers.GetAllPassengersEndpoint).Methods("GET")
	router.HandleFunc("/companies", controllers.GetAllCompaniesEndpoint).Methods("GET")
	return router
}

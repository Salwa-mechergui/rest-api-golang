package controllers

import (
	"encoding/json"
	"fmt"
	"github/reccrides/Helpers"
	"github/reccrides/Models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Init Agent var as a slice of Agent struct

var customers []Models.Customer
var front_end_users []Models.Front_end_user
var passengers []Models.Passenger
var companies []Models.Company
var rides []Models.Ride
var results []Models.Result

// // GetAllCustomerEndpoint godoc
// // @Summary Get details of all front_end_users
// // @Description Get description of all front_end_users
// // @Tags front_end_users
// // @Accept  json
// // @Produce  json
// // @Success 200 {object} Models.Front_end_user
// // @Router /frontUser [get]
// func GetAllCustomerEndpoint(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	db := Helpers.InitMigration()
// 	query := `
// 	select
// 	front_end_users.id,
// 	front_end_users.first_name,
// 	front_end_users.last_name,
// 	front_end_users.email,
// 	front_end_users.phone_number,
// 	front_end_users.saas_company_id,
// 	passengers.id,
// 	passengers.first_name,
// 	passengers.last_name,
// 	passengers.phone_number,
// 	passengers.email,
// 	passengers.default_comment,
// 	passengers.company_id,
// 	customers.id,
// 	customers.front_end_user_id,
// 	customers.default_passenger_id,
// 	customers.default_comment,
// 	customers.referral_code,
// 	customers.saas_company_id,
// 	CASE WHEN passengers.id IS NULL
// 		THEN front_end_users.first_name || ' ' || front_end_users.last_name ELSE
// 		  passengers.first_name || ' ' || passengers.last_name END
// 	from
// 	customers, front_end_users, passengers
// 	where
// 	front_end_users.id=customers.front_end_user_id
// 	limit 10
// 	`
// 	var err error
// 	if err != nil {
// 		fmt.Printf("The request failed with error %s\n", err)
// 	}
// 	err = db.Raw(query).Scan(&results).Error
// 	for _, element := range results {
// 		json.NewEncoder(w).Encode(element)
// 	}
// }

// GetAllFrontUsersEndpoint godoc
// @Summary Get details of all front_end_users
// @Description Get description of all front_end_users
// @Tags front_end_users
// @Accept  json
// @Produce  json
// @Success 200 {object} Models.Front_end_user
// @Router /frontUser [get]
func GetAllFrontUsersEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := Helpers.InitMigration()
	query := ` 
	select 
	front_end_users.id,
	front_end_users.first_name,
	front_end_users.last_name,
	front_end_users.email,
	front_end_users.phone_number,
	front_end_users.saas_company_id,
	customers.front_end_user_id,
	customers.default_passenger_id,
	customers.default_comment,
	customers.referral_code
	from 
	customers, front_end_users
	where 
	front_end_users.id=customers.front_end_user_id 
	limit 10
	`
	var err error
	if err != nil {
		fmt.Printf("The request failed with error %s\n", err)
	}
	err = db.Raw(query).Scan(&front_end_users).Error
	for _, element := range front_end_users {
		json.NewEncoder(w).Encode(element)
	}
}

// GetAllPassengersEndpoint godoc
// @Summary Get details of all passengers
// @Description Get description of all passengers
// @Tags passengers
// @Accept  json
// @Produce  json
// @Success 200 {object} Models.Passenger
// @Router /passenger [get]
func GetAllPassengersEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := Helpers.InitMigration()
	query := ` 
	select 
	passengers.id,
	passengers.first_name,
	passengers.last_name,
	passengers.phone_number,
	passengers.email,
	passengers.default_comment,
	passengers.company_id,
	customers.id,
	customers.front_end_user_id,
	customers.default_passenger_id,
	customers.default_comment,
	customers.referral_code,
	customers.saas_company_id,
	CASE WHEN passengers.id IS NULL
		THEN front_end_users.first_name || ' ' || front_end_users.last_name ELSE
		  passengers.first_name || ' ' || passengers.last_name END 
	from 
	customers,front_end_users, passengers 
	where 
	front_end_users.id=customers.front_end_user_id 
	limit 10
	`
	var err error
	if err != nil {
		fmt.Printf("The request failed with error %s\n", err)
	}
	err = db.Raw(query).Scan(&passengers).Error
	for _, element := range passengers {
		json.NewEncoder(w).Encode(element)
	}
}

// GetAllCompaniesEndpoint godoc
// @Summary Get details of all companies
// @Description Get description of all companies
// @Tags companies
// @Accept  json
// @Produce  json
// @Success 200 {object} Models.Company
// @Router /companies [get]
func GetAllCompaniesEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := Helpers.InitMigration()
	query := `
	select
	front_end_users.id,		"id"
	front_end_users.saas_company_id,	"saas_company_id"
	customers.referral_code,	"referral_code"
	customers.saas_company_id,	"saas_company_id"
	companies.id,	"id"
	companies.name,	"name"
	companies.saas_company_id,	"saas_company_id"	
	companies.default_comment	"default_comment"
	from
	customers, front_end_users, companies
	LEFT OUTER JOIN customers customer on front_end_users.id=customers.front_end_user_id
	LEFT OUTER JOIN customers customer on front_end_users.saas_company_id=customers.saas_company_id
	LEFT OUTER JOIN customers customer on companies.saas_company_id=customers.saas_company_id
	limit 10
	`
	var err error
	if err != nil {
		fmt.Printf("The request failed with error %s\n", err)
	}
	err = db.Raw(query).Scan(&companies).Error
	for _, element := range companies {
		json.NewEncoder(w).Encode(element)
	}
}

// GetFrontUsersByIdEndpoint godoc
// @Summary Get details for a given front_end_users
// @Description Get details of agent corresponding to the input front_end_users
// @Tags front_end_users
// @Accept  json
// @Produce  json
// @Param  id path int true "ID of the Front_end_users"
// @Success 200 {object} Models.Front_end_user
// @Router /frontUser/{id} [get]
func GetFrontUsersByIdEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := Helpers.InitMigration()
	query := `
	select
	front_end_users.id,
	front_end_users.first_name,
	front_end_users.last_name,
	front_end_users.email,
	front_end_users.phone_number,
	front_end_users.saas_company_id,
	customers.front_end_user_id,
	customers.default_passenger_id,
	customers.default_comment,
	customers.referral_code
	from
	customers, front_end_users
	where
	front_end_users.id=customers.front_end_user_id
	ORDER BY front_end_users.id ASC
	limit 10
	`
	var err error
	err = db.Raw(query).Scan(&front_end_users).Error
	if err != nil {
		fmt.Println(err)
	}
	params := mux.Vars(r)
	for _, item := range front_end_users {
		id, err := strconv.ParseUint(params["id"], 10, 32)
		if err != nil {
			fmt.Println(err)
		}
		if uint(id) == item.ID {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

package controllers

import (
	"BookingAPI/Book/Helpers"
	"BookingAPI/Book/Models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Init Agent var as a slice book struct
var Agentt []Models.Agent

// Get all Agent
func GetAgentEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := Helpers.InitMigration()
	json.NewEncoder(w).Encode(db.Find(&Agentt))
}

// Get single Agentt
func GetEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Gets params
	// Loop through Agent and find one with the phonenumber from the params
	for _, item := range Agentt {
		id, err1 := strconv.ParseUint(r.FormValue("id"), 10, 32)
		if err1 != nil {
			fmt.Println(err1)
		}
		if uint(id) == item.ID {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Models.Agent{})
}

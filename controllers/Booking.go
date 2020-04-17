package controllers

import (
	"encoding/json"
	"github/reccrides/Helpers"
	"github/reccrides/Models"
	"net/http"

	"github.com/gorilla/mux"
)

// Init Agent var as a slice of Agent struct
var agents []Models.Agent

//GetAgentEndpoint
// @Summary Get details of all agents
// @Description Get description of all agents
// @Tags agents
// @Accept  json
// @Produce  json
// @Success 200 {object} Models.Agent
// @Router /agent [get]
func GetAgentEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := Helpers.InitMigration()
	json.NewEncoder(w).Encode(db.Find(&agents))
}

// GetByIdEndpoint godoc
// @Summary Get details for a given agent
// @Description Get details of agent corresponding to the input agents
// @Tags agents
// @Accept  json
// @Produce  json
// @Param  id path int true "ID of the agent"
// @Success 200 {object} Models.Agent
// @Router /agent/{id} [get]
func GetByIdEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := Helpers.InitMigration()
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(db.Find(&agents, params["id"]))
}

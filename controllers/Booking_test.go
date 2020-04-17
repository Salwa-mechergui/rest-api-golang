package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAgent(t *testing.T) {
	var jsonStr = []byte(`{
		ID:2,
        Firstname: "Salwa",
		Lastname: "MECHERGUI",
		Email:"salwa@craftfoundry.com",
		Phonenumber:54122587,
		Upcomingride:5000,
		Currentride:6000,
		Comment:"test20"
    }`)
	mux := http.NewServeMux()
	mux.HandleFunc("/agent", GetAgentEndpoint)
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/agent", bytes.NewBuffer(jsonStr))
	mux.ServeHTTP(response, request)
	if response.Code != 200 {
		t.Errorf("Response code is %v", response.Code)
	}
	var result map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &result)
	// x := "success"
	fmt.Println(result["success"])
}

func TestGetByIdAgent(t *testing.T) {
	var jsonStr = []byte(`{
		ID:5,
        Firstname: "Salwa",
		Lastname: "MECHERGUI",
		Email:"salwa@craftfoundry.com",
		Phonenumber:54122587,
		Upcomingride:5000,
		Currentride:6000,
		Comment:"test20"
    }`)
	mux := http.NewServeMux()
	mux.HandleFunc("/agent{id}", GetAgentEndpoint)
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/agent{id}", bytes.NewBuffer(jsonStr))
	mux.ServeHTTP(response, request)
	if response.Code != 200 {
		t.Errorf("Response code is %v", response.Code)
	}
	var result map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &result)
	// x := "success"
	fmt.Println(result["success"])
}

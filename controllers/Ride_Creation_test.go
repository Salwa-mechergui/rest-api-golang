package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCustomer(t *testing.T) {
	var jsonStr = []byte(`{
	"ID": 0,
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "0001-01-01T00:00:00Z",
    "DeletedAt": null,
    "first_name": "salwa",
    "last_name": "mech",
    "email": "alart.be+testpdf@gmail.com",
    "phone_number": "+33625091164",
    "ride_request_id": 108921587421587412
    }`)
	mux := http.NewServeMux()
	mux.HandleFunc("/frontUser", GetAllPassengersEndpoint)
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/frontUser", bytes.NewBuffer(jsonStr))
	mux.ServeHTTP(response, request)
	if response.Code != 200 {
		t.Errorf("Response code is %v", response.Code)
	}
	var result map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &result)
	// x := "success"
	fmt.Println(result["success"])
}

func TestGetByIdPassenger(t *testing.T) {
	var jsonStr = []byte(`{
	"ID": 0,
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "0001-01-01T00:00:00Z",
    "DeletedAt": null,
    "first_name": "Bechir",
    "last_name": "Tourkii",
    "email": "alart.be+testpdf@gmail.com",
    "phone_number": "+33625091164",
    "ride_request_id": 5455
    }`)
	mux := http.NewServeMux()
	mux.HandleFunc("/frontUser/{id}", GetFrontUsersByIdEndpoint)
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/frontUser/{id}", bytes.NewBuffer(jsonStr))
	mux.ServeHTTP(response, request)
	if response.Code != 200 {
		t.Errorf("Response code is %v", response.Code)
	}
	var result map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &result)
	// x := "success"
	fmt.Println(result["success"])
}

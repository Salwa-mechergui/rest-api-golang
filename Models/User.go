package Models

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Passenger stuct
type Booking struct {
	gorm.Model
	ID int `json:"id,omitempty"`
}
type Passenger struct {
	gorm.Model
	ID                  int    `json:"id,omitempty"`
	Customer_id         string `json:"customer_id,omitempty"`
	First_name          string `json:"first_name,omitempty"`
	Last_name           string `json:"last_name,omitempty"`
	Email               string `json:"email,omitempty"`
	Phone_number        string `json:"phone_number,omitempty"`
	Company_id          int    `json:"company_id,omitempty"`
	language            string `json:"language,omitempty"`
	Deleted             bool   `json:"deleted,omitempty"`
	Locale              string `json:"locale,omitempty"`
	Send_invoice        bool   `json:"send_invoice,omitempty"`
	Show_price          string `json:"show_price,omitempty"`
	Type                string `json:"type,omitempty"`
	Default_favorite_id int    `json:"default_favorite_id,omitempty"`
	Default_comment     string `json:"default_comment,omitempty"`
}

//Customer stuct
type Customer struct {
	gorm.Model
	Id                         int    `json:"id,omitempty"`
	System_user_id             int    `json:"system_user_id,omitempty"`
	Default_passenger_id       int    `json:"default_passenger_id,omitempty"`
	Front_end_user_id          int    `json:"front_end_user_id,omitempty"`
	Saas_company_id            int    `json:"saas_company_id	,omitempty"`
	Customer_company_name      string `json:"customer_company_name,omitempty"`
	Referral_code              string `json:"referral_code,omitempty"`
	Sponsor_id                 int    `json:"sponsor_id,omitempty"`
	Cost_center                string `json:"cost_center,omitempty"`
	Default_comment            string `json:"default_comment,omitempty"`
	Send_pdf_email             bool   `json:"send_pdf_email,omitempty"`
	Default_delivery_sender_id int    `json:"default_delivery_sender_id,omitempty"`
	Sponsor_type               string `json:"sponsor_type,omitempty"`
	Company_preferences_id     int    `json:"company_preferences_id,omitempty"`
	// Uuid                       {}uuid   `json:"uuid,omitempty"`
}

//ride stuct
type Ride struct {
	gorm.Model
	Id                        int                 `json:"id,omitempty"`
	Ride_request_id           int                 `json:"ride_request_id,omitempty"`
	Ride_start_date_requested timestamp.Timestamp `json:"ride_start_date_requested,omitempty"`
	Customer_id               int                 `json:"customer_id,omitempty"`
}

//Result struct
type Result struct {
	gorm.Model
	First_name      string `json:"first_name,omitempty"`
	Last_name       string `json:"last_name,omitempty"`
	Email           string `json:"email,omitempty"`
	Phone_number    string `json:"phone_number,omitempty"`
	Ride_request_id int    `json:"ride_request_id,omitempty"`
	// Estimate_drop_off_date string `json:"estimate_drop_off_date,omitempty"`
	// Estimate_pick_up_date  string `json:"estimate_pick_up_date,omitempty"`
}
type Front_end_user struct {
	gorm.Model
	Id                   int    `json:"id_front_user,omitempty"`
	First_name           string `json:"first_name,omitempty"`
	Last_name            string `json:"last_name,omitempty"`
	Email                string `json:"email,omitempty"`
	Phone_number         string `json:"phone_number,omitempty"`
	Saas_company_id      int    `json:"saas_company_id,omitempty"`
	front_end_user_id    int    `json:"front_end_user_id,omitempty"`
	Default_passenger_id string `json:"default_passenger_id,omitempty"`
	Default_comment      string `json:"default_comment,omitempty"`
	Referral_code        string `json:"referral_code,omitempty"`
	Name                 string `json:"Name,omitempty"`
	// estimate_drop_off,
	// estimate_pick_up
}
type Company struct {
	gorm.Model      `json:"model"`
	ID              int    `json:"id,omitempty"`
	Saas_company_id string `json:"saas_company_id,omitempty"`
	Refferal_code   string `json:"referral_code,omitempty"`
	ID_company      int    `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Default_comment string `json:"default_comment,omitempty"`
}

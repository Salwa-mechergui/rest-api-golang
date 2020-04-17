package Models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//agent stuct
type Agent struct {
	gorm.Model
	ID           int    `json:"id,omitempty"`
	Firstname    string `json:"firstname,omitempty"`
	Lastname     string `json:"lastname,omitempty"`
	Email        string `json:"email,omitempty"`
	Phonenumber  int    `json:"phonenumber,omitempty"`
	Upcomingride int    `json:"upcomingride,omitempty"`
	Currentride  int    `json:"currentride,omitempty"`
	Comment      string `json:"comment,omitempty"`
}

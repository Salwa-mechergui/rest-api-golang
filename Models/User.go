package Models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//agent stuct
type Agent struct {
	gorm.Model
	ID           uint   `json:"id"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Email        string `json:"email"`
	Phonenumber  int    `json:"phonenumber"`
	Upcomingride int    `json:"upcomingride"`
	Currentride  int    `json:"currentride"`
	Comment      string `json:"comment"`
	// Agentt       []Agent `gorm:"foreignkey:BookRefer"`
}


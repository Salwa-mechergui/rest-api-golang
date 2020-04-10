package Helpers

import (
	"BookingAPI/Book/Models"
	"fmt"

	"github.com/jinzhu/gorm"
)

//
func InitMigration() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=123456 sslmode=disable")
	if err != nil {
		// defer db.Close()
		fmt.Println(err)
	}
	return db

}

//
func Migration() {
	db := InitMigration()
	db.AutoMigrate(&Models.Agent{})
}

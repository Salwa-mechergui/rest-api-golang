package Helpers

import (
	"fmt"
	"github/reccrides/Models"
	"os"

	"github.com/jinzhu/gorm"
)

//init migration function of the db parameters
func InitMigration() *gorm.DB {
	dbString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DBUSER"),
		os.Getenv("DBPASS"),
		os.Getenv("DBNAME"))

	fmt.Printf(dbString)
	db, err := gorm.Open("postgres", dbString)
	if err != nil {
		// defer db.Close()
		fmt.Println(err)
	}
	return db
}

//Migration to db
func Migration() {
	db := InitMigration()
	db.AutoMigrate(&Models.Agent{})
}

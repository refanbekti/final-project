package database

import (
	"final-project/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host = "localhost"

	user = "postgres"

	password = "admin"

	dbPort = "5432"

	dbname = "final_project"

	db *gorm.DB

	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}

	fmt.Println("sukses koneksi ke database")
	db.AutoMigrate(models.User{})
}

func GetDB() *gorm.DB {
	return db
}

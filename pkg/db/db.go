package db

import (
	"fmt"
	"log"
	"os"

	"github.com/Almazatun/golephant/infrastucture/model"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func Init() *gorm.DB {
	pg := os.Getenv("DB_PG")
	host := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_DATABASE")
	password := os.Getenv("DB_PASSWORD")

	// Database connection strings
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)

	db, err = gorm.Open(pg, dbURI)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to databases")
	}

	db.AutoMigrate(&model.Book{})
	db.AutoMigrate(&model.Person{})

	return db
}

package db

import (
	"fmt"
	"log"
	"os"

	entity "github.com/Almazatun/golephant/infrastucture/entity"
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

	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Company{})
	db.AutoMigrate(&entity.CompanyAddress{})
	db.AutoMigrate(&entity.Position{})
	db.AutoMigrate(&entity.PositionType{})
	db.AutoMigrate(&entity.UserEducation{})
	db.AutoMigrate(&entity.UserExperience{})

	db.Model(&entity.UserEducation{}).AddForeignKey("user_id", "users(user_id)", "CASCADE", "CASCADE")
	db.Model(&entity.UserExperience{}).AddForeignKey("user_id", "users(user_id)", "CASCADE", "CASCADE")
	db.Model(&entity.CompanyAddress{}).AddForeignKey("company_id", "companies(company_id)", "CASCADE", "CASCADE")
	db.Model(&entity.Position{}).AddForeignKey("position_type_id", "position_types(position_type_id)", "CASCADE", "CASCADE")

	return db
}

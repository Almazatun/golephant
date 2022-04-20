package postgresql

import (
	"fmt"
	"log"
	"os"

	entity "github.com/Almazatun/golephant/internal/infrastucture/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() *gorm.DB {
	host := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_DATABASE")
	password := os.Getenv("DB_PASSWORD")
	db_extensions := os.Getenv("POSTGRES_EXTENSIONS")

	// Database connection strings
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)

	db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to databases")
	}

	//PostgresSQL extension
	db.Exec(db_extensions)

	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Resume{})
	db.AutoMigrate(&entity.Company{})
	db.AutoMigrate(&entity.CompanyAddress{})
	db.AutoMigrate(&entity.Position{})
	db.AutoMigrate(&entity.UserEducation{})
	db.AutoMigrate(&entity.UserExperience{})

	return db
}

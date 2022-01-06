package main

import (
	"fmt"

	"github.com/Almazatun/golephant/pkg/db"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func main() {
	loadENVs()
	DB := db.Init()

	defer DB.Close()

}

func loadENVs() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error  loading .env variables")
	}
}

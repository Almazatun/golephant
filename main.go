package main

import (
	"fmt"
	"log"
	"net/http"

	usecase "github.com/Almazatun/golephant/domain"
	handler "github.com/Almazatun/golephant/handler"
	repository "github.com/Almazatun/golephant/infrastucture"
	router "github.com/Almazatun/golephant/router"

	"github.com/Almazatun/golephant/pkg/db"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func main() {
	loadENVs()
	DB := db.Init()

	userRepository := repository.InitUserRepo(DB)
	userUseCase := usecase.InitUserUseCase(userRepository)
	userHandler := handler.InitUserHandler(userUseCase)

	handlers := router.Handler{
		UserHandler: userHandler,
	}

	router := router.NewRouter(handlers)

	log.Fatal(http.ListenAndServe(":3000", router))

	defer DB.Close()

}

func loadENVs() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error  loading .env variables")
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Almazatun/golephant/common"
	usecase "github.com/Almazatun/golephant/domain"
	handler "github.com/Almazatun/golephant/handler"
	repository "github.com/Almazatun/golephant/infrastucture"
	router "github.com/Almazatun/golephant/router"
	mux_handlers "github.com/gorilla/handlers"

	"github.com/Almazatun/golephant/pkg/db"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func main() {
	loadENVs()
	DB := db.Init()

	// User
	userRepository := repository.NewUserRepo(DB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	// UserEducation
	userEducationRepository := repository.NewUserEducationRepo(DB)

	// UserExperience
	userExperienceRepository := repository.NewUserExperienceRepo(DB)

	// Position type
	positionTypeRepository := repository.NewPositionTypeRepo()
	positionTypeHandler := handler.NewPositionTypeHandler(positionTypeRepository)

	// Resume
	resumeRepository := repository.NewResumeRepo(DB)
	resumeUseCase := usecase.NewResumeUseCase(
		resumeRepository,
		userRepository,
		userEducationRepository,
		userExperienceRepository,
	)
	resumeHandler := handler.NewResumeHandler(
		resumeUseCase,
		resumeRepository,
	)

	// Company
	companyRepository := repository.NewCompanyRepo(DB)
	companyUseCase := usecase.NewCompanyUseCase(companyRepository)
	companyHandler := handler.NewCompanyHandler(companyUseCase)

	handlers := router.Handler{
		UserHandler:         userHandler,
		ResumeHandler:       resumeHandler,
		CompanyHandler:      companyHandler,
		PositionTypeHandler: positionTypeHandler,
	}

	router := router.NewRouter(handlers)

	log.Fatal(http.ListenAndServe(":3000",
		mux_handlers.CORS(
			mux_handlers.AllowedHeaders(common.CORS_ALLOWED_HEADERS),
			mux_handlers.AllowedMethods(common.CORS_ALLOWED_METHODS),
			mux_handlers.AllowedOrigins(common.CORS_ALLOWED_ORIGINS))(router)))

	defer DB.Close()

}

func loadENVs() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error  loading .env variables")
	}
}

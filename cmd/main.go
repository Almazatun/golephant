package main

import (
	"fmt"
	"log"
	"net/http"

	usecase "github.com/Almazatun/golephant/internal/domain"
	repository "github.com/Almazatun/golephant/internal/infrastucture"
	router "github.com/Almazatun/golephant/pkg/http"
	"github.com/Almazatun/golephant/pkg/http/cors"
	handler "github.com/Almazatun/golephant/pkg/http/handler"
	mux_handlers "github.com/gorilla/handlers"

	db "github.com/Almazatun/golephant/pkg/postgresql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func main() {
	loadENVs()
	DB := db.Init()

	// Reset password token
	resetPasswordToken := repository.NewResetPasswordTokenRepo(DB)
	// User
	userRepository := repository.NewUserRepo(DB)
	userUseCase := usecase.NewUserUseCase(
		userRepository,
		resetPasswordToken,
	)
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
			mux_handlers.AllowedHeaders(cors.CORS_ALLOWED_HEADERS),
			mux_handlers.AllowedMethods(cors.CORS_ALLOWED_METHODS),
			mux_handlers.AllowedOrigins(cors.CORS_ALLOWED_ORIGINS))(router)))

}

func loadENVs() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error  loading .env variables")
	}
}

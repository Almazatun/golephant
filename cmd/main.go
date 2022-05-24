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

	// Resume education
	resumeEducationRepository := repository.NewResumeEducationRepo(DB)
	resumeEducationUseCase := usecase.NewResumeEducationUseCase()

	// Resume experience
	resumeExperienceRepository := repository.NewResumeExperienceRepo(DB)
	resumeExperienceUseCase := usecase.NewResumeExperienceUseCase()

	// Position type
	positionTypeRepository := repository.NewPositionTypeRepo()
	positionTypeHandler := handler.NewPositionTypeHandler(positionTypeRepository)

	// Position
	positionRepository := repository.NewPositionRepo(DB)
	positionUseCase := usecase.NewPositionUseCase(positionRepository)

	// Specialization
	specializationRepository := repository.NewSpecializationRepo()
	specializationHandler := handler.NewSpecializationHandler(specializationRepository)

	// Resume status
	resumeStatusRepository := repository.NewResumeStatusRepo()
	resumeStatusHandler := handler.NewResumeStatusHandler(resumeStatusRepository)

	// Resume
	resumeRepository := repository.NewResumeRepo(DB)
	resumeUseCase := usecase.NewResumeUseCase(
		resumeRepository,
		userRepository,
		resumeEducationRepository,
		resumeExperienceRepository,
		// use cases
		resumeEducationUseCase,
		resumeExperienceUseCase,
	)
	resumeHandler := handler.NewResumeHandler(
		resumeUseCase,
		resumeRepository,
	)
	// Company address
	companyAddressRepository := repository.NewCompanyAddressRepo(DB)

	// Company
	companyRepository := repository.NewCompanyRepo(DB)
	companyUseCase := usecase.NewCompanyUseCase(
		companyRepository,
		companyAddressRepository,
		positionRepository,
		positionUseCase,
	)
	companyHandler := handler.NewCompanyHandler(companyUseCase)

	// Auth
	authHandler := handler.NewAuthHandler(
		companyUseCase,
		userUseCase,
	)

	handlers := router.Handler{
		UserHandler:           userHandler,
		ResumeHandler:         resumeHandler,
		CompanyHandler:        companyHandler,
		PositionTypeHandler:   positionTypeHandler,
		SpecializationHandler: specializationHandler,
		ResumeStatusHandler:   resumeStatusHandler,
		AuthHandler:           authHandler,
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

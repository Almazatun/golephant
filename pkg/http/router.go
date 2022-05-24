package router

import (
	"encoding/json"
	"net/http"

	error_message "github.com/Almazatun/golephant/pkg/common/error-message"
	handler "github.com/Almazatun/golephant/pkg/http/handler"
	"github.com/Almazatun/golephant/pkg/http/middleware"
	"github.com/gorilla/mux"
)

type Handler struct {
	UserHandler           handler.UserHandler
	ResumeHandler         handler.ResumeHandler
	CompanyHandler        handler.CompanyHandler
	PositionTypeHandler   handler.PositionTypeHandler
	SpecializationHandler handler.SpecializationHandler
	ResumeStatusHandler   handler.ResumeStatusHandler
	AuthHandler           handler.AuthHandler
}

func NewRouter(h Handler) *mux.Router {
	router := mux.NewRouter()

	// Auth
	auth := router.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/register/company", h.AuthHandler.RegisterCompany).Methods("POST")
	auth.HandleFunc("/login/company", h.AuthHandler.LogInCompany).Methods("PUT")
	auth.HandleFunc("/register/user", h.AuthHandler.RegisterUser).Methods("POST")
	auth.HandleFunc("/login/user", h.AuthHandler.LogInUser).Methods("PUT")
	auth.HandleFunc("/me", h.AuthHandler.Me).Methods("PUT")

	// Company
	company := router.PathPrefix("/companies").Subrouter()

	// Company Position
	company.Handle("/{companyId}/positions",
		middleware.Authentication(http.HandlerFunc(h.CompanyHandler.AddPosition))).Methods("POST")

	company.Handle("/{companyId}/positions/{positionId}/reponsobilities",
		middleware.Authentication(http.HandlerFunc(h.CompanyHandler.UpdatePositionResponsibilities))).Methods("PUT")

	company.Handle("/{companyId}/positions/{positionId}/requirements",
		middleware.Authentication(http.HandlerFunc(h.CompanyHandler.UpdatePositionRequirements))).Methods("PUT")

	company.Handle("/{companyId}/positions/{positionId}",
		middleware.Authentication(http.HandlerFunc(h.CompanyHandler.UpdatePosition))).Methods("PATCH")

	company.Handle("/{companyId}/positions/{positionId}/status",
		middleware.Authentication(http.HandlerFunc(h.CompanyHandler.UpdatePositionStatus))).Methods("PATCH")

	company.Handle("/{companyId}/positions/{positionId}",
		middleware.Authentication(http.HandlerFunc(h.CompanyHandler.DeletePosition))).Methods("DELETE")

	// Company Address
	company.Handle("/{companyId}/addressess",
		middleware.Authentication(http.HandlerFunc(h.CompanyHandler.AddAddress))).Methods("POST")
	company.Handle("/{companyId}/addressess/{companyAddressId}",
		middleware.Authentication(http.HandlerFunc(h.CompanyHandler.DeleteAddress))).Methods("DELETE")

	// Resumes
	router.Handle("/resumes/{userId}", middleware.Authentication(http.HandlerFunc(h.ResumeHandler.List))).Methods("GET")

	//Position types
	router.Handle("/positionTypes",
		middleware.Authentication(http.HandlerFunc(h.PositionTypeHandler.List))).Methods("GET")

	//Resume statuses
	router.Handle("/resumeStatuses",
		middleware.Authentication(http.HandlerFunc(h.ResumeStatusHandler.List))).Methods("GET")

	//Specializations
	router.Handle("/specializations",
		middleware.Authentication(http.HandlerFunc(h.SpecializationHandler.List))).Methods("GET")

	//User
	user := router.PathPrefix("/users").Subrouter()
	user.Handle("/{userId}", middleware.Authentication(http.HandlerFunc(h.UserHandler.UpdateUserData))).Methods("PATCH")

	// User reset password
	user.Handle("/{userId}/resetPassword",
		middleware.Authentication(http.HandlerFunc(h.UserHandler.GetLinkResetPassword))).Methods("POST")
	user.HandleFunc("/{userId}/resetPassword/{token}",
		h.UserHandler.ResetPassword).Methods("PUT")

	// Resume
	resume := router.PathPrefix("/resumes").Subrouter()
	resume.Handle("/users/{userId}",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.Create))).Methods("POST")
	resume.Handle("/{resumeId}",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.Delete))).Methods("DELETE")

	resume.Handle("/{resumeId}/users/{userId}/basicInfo",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.UpdateBasicInfo))).Methods("PATCH")

	resume.Handle("/{resumeId}/users/{userId}/aboutMe",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.UpdateAboutMe))).Methods("PATCH")

	resume.Handle("/{resumeId}/users/{userId}/citizenship",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.UpdateCitizenship))).Methods("PATCH")

	resume.Handle("/{resumeId}/users/{userId}/desiredPositions",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.UpdateDesiredPosition))).Methods("PATCH")

	resume.Handle("/{resumeId}/users/{userId}/tags",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.UpdateTags))).Methods("PUT")

	//  Education in resume
	resume.Handle("/{resumeId}/users/{userId}/userEducation",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.UpdateEducation))).Methods("PUT")
	resume.Handle("/{resumeId}/userEducation/{userEducationId}",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.DeleteEducation))).Methods("DELETE")

	// Experience in resume
	resume.Handle("/{resumeId}/users/{userId}/userExperience",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.UpdateExperience))).Methods("PUT")
	resume.Handle("/{resumeId}/userExperience/{userExperienceId}",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.DeleteExperience))).Methods("DELETE")

	// Test
	router.HandleFunc("/helloWorld", handler.HelloWord).Methods("GET")

	// Not found
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(error_message.ERROR_404)
	})

	return router
}

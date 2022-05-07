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
}

func NewRouter(h Handler) *mux.Router {
	router := mux.NewRouter()

	// Company
	company := router.PathPrefix("/company").Subrouter()
	company.HandleFunc("/register", h.CompanyHandler.Register).Methods("POST")
	company.HandleFunc("/login", h.CompanyHandler.LogIn).Methods("PUT")

	// Company Address
	company.Handle("/{companyId}/address",
		middleware.Authentication(http.HandlerFunc(h.CompanyHandler.AddAddress))).Methods("POST")
	company.Handle("/{companyId}/address/{companyAddressId}",
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
	user := router.PathPrefix("/user").Subrouter()
	user.HandleFunc("/register", h.UserHandler.RegisterUser).Methods("POST")
	user.HandleFunc("/login", h.UserHandler.LogIn).Methods("PUT")
	user.HandleFunc("/authMe", h.UserHandler.AuthMe).Methods("POST")
	user.Handle("/{userId}", middleware.Authentication(http.HandlerFunc(h.UserHandler.UpdateUserData))).Methods("PATCH")

	// User reset password
	user.Handle("/{userId}/resetPassword",
		middleware.Authentication(http.HandlerFunc(h.UserHandler.GetLinkResetPassword))).Methods("POST")
	user.HandleFunc("/{userId}/resetPassword/{token}",
		h.UserHandler.ResetPassword).Methods("PUT")

	// Resume
	resume := router.PathPrefix("/resume").Subrouter()
	resume.Handle("/user/{userId}",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.Create))).Methods("POST")
	resume.Handle("/{resumeId}",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.Delete))).Methods("DELETE")

	resume.Handle("/{resumeId}/user/{userId}/basicInfo",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.UpdateBasicInfo))).Methods("PUT")

	resume.Handle("/{resumeId}/user/{userId}/aboutMe",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.UpdateAboutMe))).Methods("PUT")

	resume.Handle("/{resumeId}/user/{userId}/citizenship",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.UpdateCitizenship))).Methods("PUT")

	resume.Handle("/{resumeId}/user/{userId}/desiredPosition",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.UpdateDesiredPosition))).Methods("PUT")

	resume.Handle("/{resumeId}/user/{userId}/tags",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.UpdateTags))).Methods("PUT")

	// User education in resume
	resume.Handle("/{resumeId}/user/{userId}/userEducation",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.UpdateUserEducation))).Methods("PUT")
	resume.Handle("/{resumeId}/userEducation/{userEducationId}",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.DeleteUserEducation))).Methods("DELETE")

	// User experience in resume
	resume.Handle("/{resumeId}/user/{userId}/userExperiences",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.UpdateUserExperiences))).Methods("PUT")
	resume.Handle("/{resumeId}/userExperience/{userExperienceId}",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.DeleteUserExperience))).Methods("DELETE")

	// Test
	router.HandleFunc("/helloWorld", handler.HelloWord).Methods("GET")

	// Not found
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(error_message.ERROR_404)
	})

	return router
}

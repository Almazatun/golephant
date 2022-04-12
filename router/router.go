package router

import (
	"encoding/json"
	"net/http"

	error_message "github.com/Almazatun/golephant/common/error-message"
	handler "github.com/Almazatun/golephant/handler"
	"github.com/Almazatun/golephant/middleware"
	"github.com/gorilla/mux"
)

type Handler struct {
	UserHandler    handler.UserHandler
	ResumeHandler  handler.ResumeHandler
	CompanyHandler handler.CompanyHandler
}

func NewRouter(h Handler) *mux.Router {
	router := mux.NewRouter()

	// Company
	company := router.PathPrefix("/company").Subrouter()
	company.HandleFunc("/register", h.CompanyHandler.RegisterCompany).Methods("POST")
	company.HandleFunc("/login", h.CompanyHandler.LogIn).Methods("PUT")

	// Resumes
	router.Handle("/resumes/{userId}", middleware.Authentication(http.HandlerFunc(h.ResumeHandler.UserResumes))).Methods("GET")

	//User
	user := router.PathPrefix("/user").Subrouter()
	user.HandleFunc("/register", h.UserHandler.RegisterUser).Methods("POST")
	user.HandleFunc("/login", h.UserHandler.LogIn).Methods("PUT")
	user.HandleFunc("/authMe", h.UserHandler.AuthMe).Methods("POST")
	user.Handle("/{id}", middleware.Authentication(http.HandlerFunc(h.UserHandler.UpdateUserData))).Methods("PATCH")

	// Resume
	resume := router.PathPrefix("/resume").Subrouter()
	resume.Handle("/user/{userId}", middleware.Authentication(http.HandlerFunc(h.ResumeHandler.CreateResume))).Methods("POST")
	resume.Handle("/{resumeId}", middleware.Authentication(http.HandlerFunc(h.ResumeHandler.DeleteResume))).Methods("DELETE")

	resume.Handle("/{resumeId}/user/{userId}/basicInfo",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.UpdateBasicInfoResume))).Methods("PUT")

	resume.Handle("/{resumeId}/user/{userId}/aboutMe",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.UpdateAboutMeResume))).Methods("PUT")

	resume.Handle("/{resumeId}/user/{userId}/citizenship",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.UpdateCitizenshipResume))).Methods("PUT")

	resume.Handle("/{resumeId}/user/{userId}/desiredPosition",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.UpdateDesiredPositionResume))).Methods("PUT")

	resume.Handle("/{resumeId}/user/{userId}/tags",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.UpdateTagsResumeInput))).Methods("PUT")

	// User education in resume
	resume.Handle("/{resumeId}/user/{userId}/userEducation",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.UpdateUserEducationResume))).Methods("PUT")
	resume.Handle("/{resumeId}/userEducation/{userEducationId}",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.DeleteUserEducationInResume))).Methods("DELETE")

	// User experience in resume
	resume.Handle("/{resumeId}/user/{userId}/userExperiences",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.UpdateUserExperiencesResume))).Methods("PUT")
	resume.Handle("/{resumeId}/userExperience/{userExperienceId}",
		middleware.Authentication(http.HandlerFunc(h.ResumeHandler.DeleteUserExperienceInResume))).Methods("DELETE")

	// Test
	router.HandleFunc("/helloWorld", handler.HelloWord).Methods("GET")

	// Not found
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(error_message.ERROR_404)
	})

	return router
}
